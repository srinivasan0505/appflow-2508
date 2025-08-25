/**
 * MCP Server function for Updates an existing flow.
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Post_Update_FlowMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_Update_FlowHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("X-Amz-Content-Sha256")) {
            queryParams.add("X-Amz-Content-Sha256=" + args.get("X-Amz-Content-Sha256"));
        }
        if (args.containsKey("X-Amz-Date")) {
            queryParams.add("X-Amz-Date=" + args.get("X-Amz-Date"));
        }
        if (args.containsKey("X-Amz-Algorithm")) {
            queryParams.add("X-Amz-Algorithm=" + args.get("X-Amz-Algorithm"));
        }
        if (args.containsKey("X-Amz-Credential")) {
            queryParams.add("X-Amz-Credential=" + args.get("X-Amz-Credential"));
        }
        if (args.containsKey("X-Amz-Security-Token")) {
            queryParams.add("X-Amz-Security-Token=" + args.get("X-Amz-Security-Token"));
        }
        if (args.containsKey("X-Amz-Signature")) {
            queryParams.add("X-Amz-Signature=" + args.get("X-Amz-Signature"));
        }
        if (args.containsKey("X-Amz-SignedHeaders")) {
            queryParams.add("X-Amz-SignedHeaders=" + args.get("X-Amz-SignedHeaders"));
        }
        if (args.containsKey("clientToken")) {
            queryParams.add("clientToken=" + args.get("clientToken"));
        }
        if (args.containsKey("description")) {
            queryParams.add("description=" + args.get("description"));
        }
        if (args.containsKey("flowName")) {
            queryParams.add("flowName=" + args.get("flowName"));
        }
        if (args.containsKey("sourceFlowConfig")) {
            queryParams.add("sourceFlowConfig=" + args.get("sourceFlowConfig"));
        }
        if (args.containsKey("triggerConfig")) {
            queryParams.add("triggerConfig=" + args.get("triggerConfig"));
        }
        if (args.containsKey("metadataCatalogConfig")) {
            queryParams.add("metadataCatalogConfig=" + args.get("metadataCatalogConfig"));
        }
        if (args.containsKey("tasks")) {
            queryParams.add("tasks=" + args.get("tasks"));
        }
        if (args.containsKey("destinationFlowConfigList")) {
            queryParams.add("destinationFlowConfigList=" + args.get("destinationFlowConfigList"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_update_flow" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createPost_Update_FlowTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> X-Amz-Content-Sha256Property = new HashMap<>();
        X-Amz-Content-Sha256Property.put("type", "string");
        X-Amz-Content-Sha256Property.put("required", false);
        X-Amz-Content-Sha256Property.put("description", "");
        properties.put("X-Amz-Content-Sha256", X-Amz-Content-Sha256Property);
        Map<String, Object> X-Amz-DateProperty = new HashMap<>();
        X-Amz-DateProperty.put("type", "string");
        X-Amz-DateProperty.put("required", false);
        X-Amz-DateProperty.put("description", "");
        properties.put("X-Amz-Date", X-Amz-DateProperty);
        Map<String, Object> X-Amz-AlgorithmProperty = new HashMap<>();
        X-Amz-AlgorithmProperty.put("type", "string");
        X-Amz-AlgorithmProperty.put("required", false);
        X-Amz-AlgorithmProperty.put("description", "");
        properties.put("X-Amz-Algorithm", X-Amz-AlgorithmProperty);
        Map<String, Object> X-Amz-CredentialProperty = new HashMap<>();
        X-Amz-CredentialProperty.put("type", "string");
        X-Amz-CredentialProperty.put("required", false);
        X-Amz-CredentialProperty.put("description", "");
        properties.put("X-Amz-Credential", X-Amz-CredentialProperty);
        Map<String, Object> X-Amz-Security-TokenProperty = new HashMap<>();
        X-Amz-Security-TokenProperty.put("type", "string");
        X-Amz-Security-TokenProperty.put("required", false);
        X-Amz-Security-TokenProperty.put("description", "");
        properties.put("X-Amz-Security-Token", X-Amz-Security-TokenProperty);
        Map<String, Object> X-Amz-SignatureProperty = new HashMap<>();
        X-Amz-SignatureProperty.put("type", "string");
        X-Amz-SignatureProperty.put("required", false);
        X-Amz-SignatureProperty.put("description", "");
        properties.put("X-Amz-Signature", X-Amz-SignatureProperty);
        Map<String, Object> X-Amz-SignedHeadersProperty = new HashMap<>();
        X-Amz-SignedHeadersProperty.put("type", "string");
        X-Amz-SignedHeadersProperty.put("required", false);
        X-Amz-SignedHeadersProperty.put("description", "");
        properties.put("X-Amz-SignedHeaders", X-Amz-SignedHeadersProperty);
        Map<String, Object> clientTokenProperty = new HashMap<>();
        clientTokenProperty.put("type", "string");
        clientTokenProperty.put("required", false);
        clientTokenProperty.put("description", "Input parameter: <p>The <code>clientToken</code> parameter is an idempotency token. It ensures that your <code>UpdateFlow</code> request completes only once. You choose the value to pass. For example, if you don't receive a response from your request, you can safely retry the request with the same <code>clientToken</code> parameter value.</p> <p>If you omit a <code>clientToken</code> value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.</p> <p>If you specify input parameters that differ from your first request, an error occurs. If you use a different value for <code>clientToken</code>, Amazon AppFlow considers it a new call to <code>UpdateFlow</code>. The token is active for 8 hours.</p>");
        properties.put("clientToken", clientTokenProperty);
        Map<String, Object> descriptionProperty = new HashMap<>();
        descriptionProperty.put("type", "string");
        descriptionProperty.put("required", false);
        descriptionProperty.put("description", "Input parameter: A description of the flow.");
        properties.put("description", descriptionProperty);
        Map<String, Object> flowNameProperty = new HashMap<>();
        flowNameProperty.put("type", "string");
        flowNameProperty.put("required", true);
        flowNameProperty.put("description", "Input parameter: The specified name of the flow. Spaces are not allowed. Use underscores (_) or hyphens (-) only.");
        properties.put("flowName", flowNameProperty);
        Map<String, Object> sourceFlowConfigProperty = new HashMap<>();
        sourceFlowConfigProperty.put("type", "string");
        sourceFlowConfigProperty.put("required", true);
        sourceFlowConfigProperty.put("description", "Input parameter: Contains information about the configuration of the source connector used in the flow.");
        properties.put("sourceFlowConfig", sourceFlowConfigProperty);
        Map<String, Object> triggerConfigProperty = new HashMap<>();
        triggerConfigProperty.put("type", "string");
        triggerConfigProperty.put("required", true);
        triggerConfigProperty.put("description", "Input parameter: The trigger settings that determine how and when Amazon AppFlow runs the specified flow.");
        properties.put("triggerConfig", triggerConfigProperty);
        Map<String, Object> metadataCatalogConfigProperty = new HashMap<>();
        metadataCatalogConfigProperty.put("type", "string");
        metadataCatalogConfigProperty.put("required", false);
        metadataCatalogConfigProperty.put("description", "Input parameter: Specifies the configuration that Amazon AppFlow uses when it catalogs your data. When Amazon AppFlow catalogs your data, it stores metadata in a data catalog.");
        properties.put("metadataCatalogConfig", metadataCatalogConfigProperty);
        Map<String, Object> tasksProperty = new HashMap<>();
        tasksProperty.put("type", "string");
        tasksProperty.put("required", true);
        tasksProperty.put("description", "Input parameter: A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.");
        properties.put("tasks", tasksProperty);
        Map<String, Object> destinationFlowConfigListProperty = new HashMap<>();
        destinationFlowConfigListProperty.put("type", "string");
        destinationFlowConfigListProperty.put("required", true);
        destinationFlowConfigListProperty.put("description", "Input parameter: The configuration that controls how Amazon AppFlow transfers data to the destination connector.");
        properties.put("destinationFlowConfigList", destinationFlowConfigListProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_update_flow",
            "Updates an existing flow.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_Update_FlowHandler(config));
    }
    
}