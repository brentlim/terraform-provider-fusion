/*
Copyright 2022 Pure Storage Inc
SPDX-License-Identifier: Apache-2.0
*/

// Code generated DO NOT EDIT.
/*
 * Pure Fusion API
 *
 * Pure Fusion is fully API-driven. Most APIs which change the system (POST, PATCH, DELETE) return an Operation in status \"Pending\" or \"Running\". You can poll (GET) the operation to check its status, waiting for it to change to \"Succeeded\" or \"Failed\".
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package fusion

type PlacementGroupPost struct {
	// The name of the resource, supplied by the user at creation, and used in the URI path of a resource.
	Name string `json:"name"`
	// The display name of the resource.
	DisplayName string `json:"display_name,omitempty"`
	// The name of the Availability Zone that this PG should be created.
	AvailabilityZone string `json:"availability_zone"`
	// The name of the Region associated with the mentioned Availability Zone.
	Region string `json:"region"`
	// The name of the parent Storage Service
	StorageService string `json:"storage_service"`
}
