/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

import (
	"time"
)

// Compute task attributes which include values derived from the TaskSpec, as well as plugin-specific data and task metadata.
type AdminTaskClosure struct {
	// Represents the compiled representation of the task from the specification provided.
	CompiledTask *CoreCompiledTask `json:"compiled_task,omitempty"`
	// Time at which the task was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
}
