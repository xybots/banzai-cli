/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// ClustergroupDeploymentsApiService ClustergroupDeploymentsApi service
type ClustergroupDeploymentsApiService service

// ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNameDeleteOpts Optional parameters for the method 'ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNameDelete'
type ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNameDeleteOpts struct {
    Force optional.Bool
}

/*
ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNameDelete Delete Cluster Group Deployment
deletes a cluster group deployment, also deleting deployments from member clusters
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization identifier
 * @param clusterGroupId Cluster Group ID
 * @param deploymentName release name of a cluster group deployment
 * @param optional nil or *ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNameDeleteOpts - Optional Parameters:
 * @param "Force" (optional.Bool) -  if true cluster group deployment gets deleted even if some deployments can not be deleted from each target cluster
@return DeploymentTargetClusterStatus
*/
func (a *ClustergroupDeploymentsApiService) ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNameDelete(ctx _context.Context, orgId int32, clusterGroupId int32, deploymentName string, localVarOptionals *ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNameDeleteOpts) (DeploymentTargetClusterStatus, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeploymentTargetClusterStatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{orgid}/clustergroups/{clusterGroupId}/deployments/{deploymentName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"clusterGroupId"+"}", _neturl.QueryEscape(parameterToString(clusterGroupId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"deploymentName"+"}", _neturl.QueryEscape(parameterToString(deploymentName, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Force.IsSet() {
		localVarQueryParams.Add("force", parameterToString(localVarOptionals.Force.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v CommonError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNameGet Get Cluster Group Deployment
retrieve info about a cluster group deployment and it&#39;s status on each member cluster
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization identifier
 * @param clusterGroupId Cluster Group ID
 * @param deploymentName release name of a cluster group deployment
@return DeploymentDeploymentInfo
*/
func (a *ClustergroupDeploymentsApiService) ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNameGet(ctx _context.Context, orgId int32, clusterGroupId int32, deploymentName string) (DeploymentDeploymentInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeploymentDeploymentInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{orgid}/clustergroups/{clusterGroupId}/deployments/{deploymentName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"clusterGroupId"+"}", _neturl.QueryEscape(parameterToString(clusterGroupId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"deploymentName"+"}", _neturl.QueryEscape(parameterToString(deploymentName, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v CommonError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNamePut Update Cluster Group Deployment
updates a cluster group deployment, installs or upgrades deployment on each member cluster accordingly
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization identifier
 * @param clusterGroupId Cluster Group ID
 * @param deploymentName release name of a cluster group deployment
 * @param deploymentClusterGroupDeployment Deployment Update Request
@return DeploymentCreateUpdateDeploymentResponse
*/
func (a *ClustergroupDeploymentsApiService) ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNamePut(ctx _context.Context, orgId int32, clusterGroupId int32, deploymentName string, deploymentClusterGroupDeployment DeploymentClusterGroupDeployment) (DeploymentCreateUpdateDeploymentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeploymentCreateUpdateDeploymentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{orgid}/clustergroups/{clusterGroupId}/deployments/{deploymentName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"clusterGroupId"+"}", _neturl.QueryEscape(parameterToString(clusterGroupId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"deploymentName"+"}", _neturl.QueryEscape(parameterToString(deploymentName, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &deploymentClusterGroupDeployment
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v CommonError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNameSyncPut Synchronize Cluster Group Deployment
install / upgrade deployment to target clusters where deployment is not found or has wrong
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization identifier
 * @param clusterGroupId Cluster Group ID
 * @param deploymentName release name of a cluster group deployment
@return DeploymentTargetClusterStatus
*/
func (a *ClustergroupDeploymentsApiService) ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsDeploymentNameSyncPut(ctx _context.Context, orgId int32, clusterGroupId int32, deploymentName string) (DeploymentTargetClusterStatus, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeploymentTargetClusterStatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{orgid}/clustergroups/{clusterGroupId}/deployments/{deploymentName}/sync"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"clusterGroupId"+"}", _neturl.QueryEscape(parameterToString(clusterGroupId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"deploymentName"+"}", _neturl.QueryEscape(parameterToString(deploymentName, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v CommonError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsGet Get All Deployments of a Cluster Group
retrieve all deployments from a cluster group
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization identifier
 * @param clusterGroupId Cluster Group ID
@return []DeploymentListDeploymentResponse
*/
func (a *ClustergroupDeploymentsApiService) ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsGet(ctx _context.Context, orgId int32, clusterGroupId int32) ([]DeploymentListDeploymentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []DeploymentListDeploymentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{orgid}/clustergroups/{clusterGroupId}/deployments"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"clusterGroupId"+"}", _neturl.QueryEscape(parameterToString(clusterGroupId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v CommonError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsPost Create Cluster Group Deployment
creates a new cluster group deployment, installs or upgrades deployment on each member cluster accordingly
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization identifier
 * @param clusterGroupId Cluster Group ID
 * @param deploymentClusterGroupDeployment Deployment Create Request
@return DeploymentCreateUpdateDeploymentResponse
*/
func (a *ClustergroupDeploymentsApiService) ApiV1OrgsOrgidClustergroupsClusterGroupIdDeploymentsPost(ctx _context.Context, orgId int32, clusterGroupId int32, deploymentClusterGroupDeployment DeploymentClusterGroupDeployment) (DeploymentCreateUpdateDeploymentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeploymentCreateUpdateDeploymentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{orgid}/clustergroups/{clusterGroupId}/deployments"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"clusterGroupId"+"}", _neturl.QueryEscape(parameterToString(clusterGroupId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &deploymentClusterGroupDeployment
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v CommonError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
