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
	"time"
)
// ReportPkeNodeStatusRequest struct for ReportPkeNodeStatusRequest
type ReportPkeNodeStatusRequest struct {
	// name of node
	Name string `json:"name,omitempty"`
	// name of nodepool
	NodePool string `json:"nodePool,omitempty"`
	// ip address of node (where the other nodes can reach it)
	Ip string `json:"ip,omitempty"`
	// detailed description about the current bootstrapping status (including the cause of the failure)
	Message string `json:"message,omitempty"`
	// the current phase of the bootstrap process
	Phase string `json:"phase,omitempty"`
	// if this is the final status report, that describes the conclusion of the whole process
	Final bool `json:"final,omitempty"`
	Status ProcessStatus `json:"status,omitempty"`
	// exact time of event
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// ID of the process registered earlier (register new process if empty)
	ProcessId string `json:"processId,omitempty"`
}
