package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-appflow/mcp-server/config"
	"github.com/amazon-appflow/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ResetconnectormetadatacacheHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ResetConnectorMetadataCacherequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/reset-connector-metadata-cache", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Content-Sha256"]; ok {
			req.Header.Set("X-Amz-Content-Sha256", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Date"]; ok {
			req.Header.Set("X-Amz-Date", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Algorithm"]; ok {
			req.Header.Set("X-Amz-Algorithm", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Credential"]; ok {
			req.Header.Set("X-Amz-Credential", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Security-Token"]; ok {
			req.Header.Set("X-Amz-Security-Token", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Signature"]; ok {
			req.Header.Set("X-Amz-Signature", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-SignedHeaders"]; ok {
			req.Header.Set("X-Amz-SignedHeaders", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.ResetConnectorMetadataCacheResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateResetconnectormetadatacacheTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_reset-connector-metadata-cache",
		mcp.WithDescription("<p>Resets metadata about your connector entities that Amazon AppFlow stored in its cache. Use this action when you want Amazon AppFlow to return the latest information about the data that you have in a source application.</p> <p>Amazon AppFlow returns metadata about your entities when you use the ListConnectorEntities or DescribeConnectorEntities actions. Following these actions, Amazon AppFlow caches the metadata to reduce the number of API requests that it must send to the source application. Amazon AppFlow automatically resets the cache once every hour, but you can use this action when you want to get the latest metadata right away.</p>"),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithString("connectorType", mcp.Description("Input parameter: <p>The type of connector to reset cached metadata for.</p> <p>You must include this parameter in your request if you're resetting the cache for any of the following connectors: Amazon Connect, Amazon EventBridge, Amazon Lookout for Metrics, Amazon S3, or Upsolver. If you're resetting the cache for any other connector, you can omit this parameter from your request. </p>")),
		mcp.WithString("entitiesPath", mcp.Description("Input parameter: <p>Use this parameter only if you’re resetting the cached metadata about a nested entity. Only some connectors support nested entities. A nested entity is one that has another entity as a parent. To use this parameter, specify the name of the parent entity.</p> <p>To look up the parent-child relationship of entities, you can send a ListConnectorEntities request that omits the entitiesPath parameter. Amazon AppFlow will return a list of top-level entities. For each one, it indicates whether the entity has nested entities. Then, in a subsequent ListConnectorEntities request, you can specify a parent entity name for the entitiesPath parameter. Amazon AppFlow will return a list of the child entities for that parent.</p>")),
		mcp.WithString("apiVersion", mcp.Description("Input parameter: <p>The API version that you specified in the connector profile that you’re resetting cached metadata for. You must use this parameter only if the connector supports multiple API versions or if the connector type is CustomConnector.</p> <p>To look up how many versions a connector supports, use the DescribeConnectors action. In the response, find the value that Amazon AppFlow returns for the connectorVersion parameter.</p> <p>To look up the connector type, use the DescribeConnectorProfiles action. In the response, find the value that Amazon AppFlow returns for the connectorType parameter.</p> <p>To look up the API version that you specified in a connector profile, use the DescribeConnectorProfiles action.</p>")),
		mcp.WithString("connectorEntityName", mcp.Description("Input parameter: <p>Use this parameter if you want to reset cached metadata about the details for an individual entity.</p> <p>If you don't include this parameter in your request, Amazon AppFlow only resets cached metadata about entity names, not entity details.</p>")),
		mcp.WithString("connectorProfileName", mcp.Description("Input parameter: <p>The name of the connector profile that you want to reset cached metadata for.</p> <p>You can omit this parameter if you're resetting the cache for any of the following connectors: Amazon Connect, Amazon EventBridge, Amazon Lookout for Metrics, Amazon S3, or Upsolver. If you're resetting the cache for any other connector, you must include this parameter in your request.</p>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ResetconnectormetadatacacheHandler(cfg),
	}
}
