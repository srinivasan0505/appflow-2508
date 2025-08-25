package main

import (
	"github.com/amazon-appflow/mcp-server/config"
	"github.com/amazon-appflow/mcp-server/models"
	tools_tags "github.com/amazon-appflow/mcp-server/tools/tags"
	tools_list_connector_entities "github.com/amazon-appflow/mcp-server/tools/list_connector_entities"
	tools_describe_connectors "github.com/amazon-appflow/mcp-server/tools/describe_connectors"
	tools_start_flow "github.com/amazon-appflow/mcp-server/tools/start_flow"
	tools_stop_flow "github.com/amazon-appflow/mcp-server/tools/stop_flow"
	tools_delete_connector_profile "github.com/amazon-appflow/mcp-server/tools/delete_connector_profile"
	tools_update_flow "github.com/amazon-appflow/mcp-server/tools/update_flow"
	tools_cancel_flow_executions "github.com/amazon-appflow/mcp-server/tools/cancel_flow_executions"
	tools_create_flow "github.com/amazon-appflow/mcp-server/tools/create_flow"
	tools_register_connector "github.com/amazon-appflow/mcp-server/tools/register_connector"
	tools_update_connector_profile "github.com/amazon-appflow/mcp-server/tools/update_connector_profile"
	tools_delete_flow "github.com/amazon-appflow/mcp-server/tools/delete_flow"
	tools_create_connector_profile "github.com/amazon-appflow/mcp-server/tools/create_connector_profile"
	tools_unregister_connector "github.com/amazon-appflow/mcp-server/tools/unregister_connector"
	tools_update_connector_registration "github.com/amazon-appflow/mcp-server/tools/update_connector_registration"
	tools_describe_flow "github.com/amazon-appflow/mcp-server/tools/describe_flow"
	tools_describe_flow_execution_records "github.com/amazon-appflow/mcp-server/tools/describe_flow_execution_records"
	tools_list_connectors "github.com/amazon-appflow/mcp-server/tools/list_connectors"
	tools_list_flows "github.com/amazon-appflow/mcp-server/tools/list_flows"
	tools_describe_connector_entity "github.com/amazon-appflow/mcp-server/tools/describe_connector_entity"
	tools_describe_connector_profiles "github.com/amazon-appflow/mcp-server/tools/describe_connector_profiles"
	tools_describe_connector "github.com/amazon-appflow/mcp-server/tools/describe_connector"
	tools_reset_connector_metadata_cache "github.com/amazon-appflow/mcp-server/tools/reset_connector_metadata_cache"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_list_connector_entities.CreateListconnectorentitiesTool(cfg),
		tools_describe_connectors.CreateDescribeconnectorsTool(cfg),
		tools_start_flow.CreateStartflowTool(cfg),
		tools_stop_flow.CreateStopflowTool(cfg),
		tools_delete_connector_profile.CreateDeleteconnectorprofileTool(cfg),
		tools_update_flow.CreateUpdateflowTool(cfg),
		tools_cancel_flow_executions.CreateCancelflowexecutionsTool(cfg),
		tools_create_flow.CreateCreateflowTool(cfg),
		tools_register_connector.CreateRegisterconnectorTool(cfg),
		tools_update_connector_profile.CreateUpdateconnectorprofileTool(cfg),
		tools_delete_flow.CreateDeleteflowTool(cfg),
		tools_create_connector_profile.CreateCreateconnectorprofileTool(cfg),
		tools_unregister_connector.CreateUnregisterconnectorTool(cfg),
		tools_update_connector_registration.CreateUpdateconnectorregistrationTool(cfg),
		tools_describe_flow.CreateDescribeflowTool(cfg),
		tools_describe_flow_execution_records.CreateDescribeflowexecutionrecordsTool(cfg),
		tools_list_connectors.CreateListconnectorsTool(cfg),
		tools_list_flows.CreateListflowsTool(cfg),
		tools_describe_connector_entity.CreateDescribeconnectorentityTool(cfg),
		tools_describe_connector_profiles.CreateDescribeconnectorprofilesTool(cfg),
		tools_describe_connector.CreateDescribeconnectorTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_reset_connector_metadata_cache.CreateResetconnectormetadatacacheTool(cfg),
	}
}
