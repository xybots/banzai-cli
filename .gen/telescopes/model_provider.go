/*
 * Cluster Recommender.
 *
 * This project can be used to recommend instance type groups on different cloud providers consisting of regular and spot/preemptible instances. The main goal is to provide and continuously manage a cost-effective but still stable cluster layout that's built up from a diverse set of regular and spot instances.
 *
 * API version: 0.5.3
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package telescopes
// Provider struct for Provider
type Provider struct {
	Provider string `json:"provider,omitempty"`
	Services []string `json:"services,omitempty"`
}