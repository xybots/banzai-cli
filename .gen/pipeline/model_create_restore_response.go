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
// CreateRestoreResponse struct for CreateRestoreResponse
type CreateRestoreResponse struct {
	Restore RestoreResponse `json:"restore,omitempty"`
	Status int32 `json:"status,omitempty"`
}
