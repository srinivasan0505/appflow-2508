

# CreateFlowRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**flowName** | **String** |  The specified name of the flow. Spaces are not allowed. Use underscores (_) or hyphens (-) only.  |  |
|**description** | **String** |  A description of the flow you want to create.  |  [optional] |
|**kmsArn** | **String** |  The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption. This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don&#39;t provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.  |  [optional] |
|**triggerConfig** | [**CreateFlowRequestTriggerConfig**](CreateFlowRequestTriggerConfig.md) |  |  |
|**sourceFlowConfig** | [**CreateFlowRequestSourceFlowConfig**](CreateFlowRequestSourceFlowConfig.md) |  |  |
|**destinationFlowConfigList** | [**List&lt;DestinationFlowConfig&gt;**](DestinationFlowConfig.md) |  The configuration that controls how Amazon AppFlow places data in the destination connector.  |  |
|**tasks** | [**List&lt;Task&gt;**](Task.md) |  A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.  |  |
|**tags** | **Map&lt;String, String&gt;** |  The tags used to organize, track, or control access for your flow.  |  [optional] |
|**metadataCatalogConfig** | [**CreateFlowRequestMetadataCatalogConfig**](CreateFlowRequestMetadataCatalogConfig.md) |  |  [optional] |
|**clientToken** | **String** | &lt;p&gt;The &lt;code&gt;clientToken&lt;/code&gt; parameter is an idempotency token. It ensures that your &lt;code&gt;CreateFlow&lt;/code&gt; request completes only once. You choose the value to pass. For example, if you don&#39;t receive a response from your request, you can safely retry the request with the same &lt;code&gt;clientToken&lt;/code&gt; parameter value.&lt;/p&gt; &lt;p&gt;If you omit a &lt;code&gt;clientToken&lt;/code&gt; value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.&lt;/p&gt; &lt;p&gt;If you specify input parameters that differ from your first request, an error occurs. If you use a different value for &lt;code&gt;clientToken&lt;/code&gt;, Amazon AppFlow considers it a new call to &lt;code&gt;CreateFlow&lt;/code&gt;. The token is active for 8 hours.&lt;/p&gt; |  [optional] |



