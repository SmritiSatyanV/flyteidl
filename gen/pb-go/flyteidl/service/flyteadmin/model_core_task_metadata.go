/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type CoreTaskMetadata struct {
	// Indicates whether the system should attempt to lookup this task's output to avoid duplication of work.
	Discoverable bool `json:"discoverable,omitempty"`
	// Runtime information about the task.
	Runtime *CoreRuntimeMetadata `json:"runtime,omitempty"`
	// The overall timeout of a task including user-triggered retries.
	Timeout string `json:"timeout,omitempty"`
	// Number of retries per task.
	Retries *CoreRetryStrategy `json:"retries,omitempty"`
	// Indicates a logical version to apply to this task for the purpose of discovery.
	DiscoveryVersion string `json:"discovery_version,omitempty"`
	// If set, this indicates that this task is deprecated.  This will enable owners of tasks to notify consumers of the ending of support for a given task.
	DeprecatedErrorMessage string `json:"deprecated_error_message,omitempty"`
	Interruptible          bool   `json:"interruptible,omitempty"`
}
