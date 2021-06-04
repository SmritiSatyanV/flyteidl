/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type AdminNodeExecutionMetaData struct {
	// Node executions are grouped depending on retries of the parent Retry group is unique within the context of a parent node.
	RetryGroup   string `json:"retry_group,omitempty"`
	IsParentNode bool   `json:"is_parent_node,omitempty"`
	SpecNodeId   string `json:"spec_node_id,omitempty"`
}
