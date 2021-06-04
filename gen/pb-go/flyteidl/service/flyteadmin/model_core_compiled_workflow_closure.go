/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// A Compiled Workflow Closure contains all the information required to start a new execution, or to visualize a workflow and its details. The CompiledWorkflowClosure should always contain a primary workflow, that is the main workflow that will being the execution. All subworkflows are denormalized. WorkflowNodes refer to the workflow identifiers of compiled subworkflows.
type CoreCompiledWorkflowClosure struct {
	Primary      *CoreCompiledWorkflow  `json:"primary,omitempty"`
	SubWorkflows []CoreCompiledWorkflow `json:"sub_workflows,omitempty"`
	Tasks        []CoreCompiledTask     `json:"tasks,omitempty"`
}
