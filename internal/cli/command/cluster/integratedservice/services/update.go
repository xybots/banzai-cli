// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package services

import (
	"context"
	"encoding/json"
	"fmt"

	"emperror.dev/errors"
	"github.com/AlecAivazis/survey/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/banzaicloud/banzai-cli/.gen/pipeline"
	"github.com/banzaicloud/banzai-cli/internal/cli"
	clustercontext "github.com/banzaicloud/banzai-cli/internal/cli/command/cluster/context"
	"github.com/banzaicloud/banzai-cli/internal/cli/utils"
)

type updateOptions struct {
	clustercontext.Context
	filePath string
}

type updateManager interface {
	ReadableName() string
	ServiceName() string
	BuildUpdateRequestInteractively(clusterCtx clustercontext.Context, request *pipeline.UpdateIntegratedServiceRequest) error
	specValidator
}

func newUpdateCommand(banzaiCLI cli.Cli, use string, mngr updateManager) *cobra.Command {
	options := updateOptions{}

	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"change", "modify", "set"},
		Short:   fmt.Sprintf("Update the %s service of a cluster", mngr.ReadableName()),
		Args:    cobra.NoArgs,
		RunE: func(_ *cobra.Command, args []string) error {
			return runUpdate(banzaiCLI, mngr, options, args, use)
		},
	}

	options.Context = clustercontext.NewClusterContext(cmd, banzaiCLI, fmt.Sprintf("update %s cluster service for", mngr.ReadableName()))

	flags := cmd.Flags()
	flags.StringVarP(&options.filePath, "file", "f", "", "Service specification file")

	return cmd
}

func runUpdate(
	banzaiCLI cli.Cli,
	m updateManager,
	options updateOptions,
	args []string,
	use string,
) error {
	if err := isServiceEnabled(context.Background(), banzaiCLI, use); err != nil {
		return errors.WrapIf(err, "failed to check service")
	}

	if err := options.Init(args...); err != nil {
		return errors.Wrap(err, "failed to initialize options")
	}

	orgID := banzaiCLI.Context().OrganizationID()
	clusterID := options.ClusterID()

	var (
		err     error
		request pipeline.UpdateIntegratedServiceRequest
	)
	if options.filePath == "" && banzaiCLI.Interactive() {
		// get integratedservice details
		details, _, err := banzaiCLI.Client().IntegratedServicesApi.IntegratedServiceDetails(context.Background(), orgID, clusterID, m.ServiceName())
		if err != nil {
			return errors.WrapIf(err, "failed to get service details")
		}

		request.Spec = details.Spec

		if err := m.BuildUpdateRequestInteractively(options.Context, &request); err != nil {
			return errors.WrapIf(err, "failed to build update request interactively")
		}

		// show editor
		if err := showUpdateEditor(m, &request); err != nil {
			return errors.WrapIf(err, "failed during showing editor")
		}
	} else {
		if err := readUpdateReqFromFileOrStdin(options.filePath, &request); err != nil {
			return errors.WrapIf(err, fmt.Sprintf("failed to read %s cluster service specification", m.ReadableName()))
		}
	}

	resp, err := banzaiCLI.Client().IntegratedServicesApi.UpdateIntegratedService(context.Background(), orgID, clusterID, m.ServiceName(), request)
	if err != nil {
		cli.LogAPIError(fmt.Sprintf("update %s cluster service", m.ReadableName()), err, resp.Request)
		log.Fatalf("could not update %s cluster service: %v", m.ReadableName(), err)
	}

	log.Infof("service %q started to update", m.ReadableName())

	return nil
}

func readUpdateReqFromFileOrStdin(filePath string, req *pipeline.UpdateIntegratedServiceRequest) error {
	filename, raw, err := utils.ReadFileOrStdin(filePath)
	if err != nil {
		return errors.WrapIfWithDetails(err, "failed to read", "filename", filename)
	}

	if err := json.Unmarshal(raw, &req); err != nil {
		return errors.WrapIf(err, "failed to unmarshal input")
	}

	return nil
}

func showUpdateEditor(m updateManager, request *pipeline.UpdateIntegratedServiceRequest) error {
	var edit bool
	if err := survey.AskOne(
		&survey.Confirm{
			Message: "Do you want to edit the cluster service update request in your text editor?",
		},
		&edit,
	); err != nil {
		return errors.WrapIf(err, "failure during survey")
	}
	if !edit {
		return nil
	}

	content, err := json.MarshalIndent(*request, "", "  ")
	if err != nil {
		return errors.WrapIf(err, "failed to marshal request to JSON")
	}
	var result string
	if err := survey.AskOne(
		&survey.Editor{
			Default:       string(content),
			HideDefault:   true,
			AppendDefault: true,
		},
		&result,
		survey.WithValidator(makeUpdateRequestValidator(m)),
	); err != nil {
		return errors.WrapIf(err, "failure during survey")
	}
	if err := json.Unmarshal([]byte(result), &request); err != nil {
		return errors.WrapIf(err, "failed to unmarshal JSON as request")
	}

	return nil
}

func makeUpdateRequestValidator(specValidator specValidator) survey.Validator {
	return func(v interface{}) error {
		var req pipeline.UpdateIntegratedServiceRequest
		if err := json.Unmarshal([]byte(v.(string)), &req); err != nil {
			return errors.WrapIf(err, "request is not valid JSON")
		}

		return specValidator.ValidateSpec(req.Spec)
	}
}
