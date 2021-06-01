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

type activateOptions struct {
	clustercontext.Context
	filePath string
}

type activateManager interface {
	ReadableName() string
	ServiceName() string
	BuildActivateRequestInteractively(clusterCtx clustercontext.Context) (pipeline.ActivateIntegratedServiceRequest, error)
	specValidator
}

func newActivateCommand(banzaiCLI cli.Cli, use string, mngr activateManager) *cobra.Command {
	options := activateOptions{}

	cmd := &cobra.Command{
		Use:           "activate",
		Aliases:       []string{"add", "enable", "install", "on"},
		Short:         fmt.Sprintf("Activate the %s service of a cluster", mngr.ReadableName()),
		Args:          cobra.NoArgs,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			return runActivate(banzaiCLI, mngr, options, args, use)
		},
	}

	options.Context = clustercontext.NewClusterContext(cmd, banzaiCLI, fmt.Sprintf("activate %s cluster service for", mngr.ReadableName()))

	flags := cmd.Flags()
	flags.StringVarP(&options.filePath, "file", "f", "", "Service specification file")

	return cmd
}

func runActivate(banzaiCLI cli.Cli, m activateManager, options activateOptions, args []string, use string) error {
	if err := isServiceEnabled(context.Background(), banzaiCLI, use); err != nil {
		return errors.WrapIf(err, "failed to check service")
	}

	if err := options.Init(args...); err != nil {
		return errors.Wrap(err, "failed to initialize options")
	}

	var (
		request pipeline.ActivateIntegratedServiceRequest
		err     error
	)

	if options.filePath == "" && banzaiCLI.Interactive() {
		if request, err = m.BuildActivateRequestInteractively(options.Context); err != nil {
			return errors.WrapIf(err, "failed to build activate request interactively")
		}

		if err := showActivateEditor(m, &request); err != nil {
			return errors.WrapIf(err, "failed during showing editor")
		}
	} else {
		if err = readActivateReqFromFileOrStdin(options.filePath, &request); err != nil {
			return errors.WrapIf(err, fmt.Sprintf("failed to read %s cluster service specification", m.ReadableName()))
		}
	}

	orgId := banzaiCLI.Context().OrganizationID()
	clusterId := options.ClusterID()
	_, err = banzaiCLI.Client().IntegratedServicesApi.ActivateIntegratedService(context.Background(), orgId, clusterId, m.ServiceName(), request)
	if err != nil {
		cli.LogAPIError(fmt.Sprintf("activate %s cluster service", m.ReadableName()), err, request)
		log.Fatalf("could not activate %s cluster service: %v", m.ReadableName(), err)
	}

	log.Infof("service %q started to activate", m.ReadableName())

	return nil
}

func readActivateReqFromFileOrStdin(filePath string, req *pipeline.ActivateIntegratedServiceRequest) error {
	filename, raw, err := utils.ReadFileOrStdin(filePath)
	if err != nil {
		return errors.WrapIfWithDetails(err, "failed to read", "filename", filename)
	}

	if err := json.Unmarshal(raw, &req); err != nil {
		return errors.WrapIf(err, "failed to unmarshal input")
	}

	return nil
}

func showActivateEditor(m activateManager, req *pipeline.ActivateIntegratedServiceRequest) error {
	var edit bool
	if err := survey.AskOne(
		&survey.Confirm{
			Message: "Do you want to edit the cluster service activation request in your text editor?",
		},
		&edit,
	); err != nil {
		return errors.WrapIf(err, "failure during survey")
	}
	if !edit {
		return nil
	}

	content, err := json.MarshalIndent(*req, "", "  ")
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
		survey.WithValidator(makeActivationRequestValidator(m)),
	); err != nil {
		return errors.WrapIf(err, "failure during survey")
	}
	if err := json.Unmarshal([]byte(result), &req); err != nil {
		return errors.WrapIf(err, "failed to unmarshal JSON as request")
	}

	return nil
}

func makeActivationRequestValidator(specValidator specValidator) survey.Validator {
	return func(v interface{}) error {
		var req pipeline.ActivateIntegratedServiceRequest
		if err := json.Unmarshal([]byte(v.(string)), &req); err != nil {
			return errors.WrapIf(err, "request is not valid JSON")
		}

		return specValidator.ValidateSpec(req.Spec)
	}
}
