

# UpdateFlowRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**flowName** | **String** |  The specified name of the flow. Spaces are not allowed. Use underscores (_) or hyphens (-) only.  |  |
|**description** | **String** |  A description of the flow.  |  [optional] |
|**triggerConfig** | [**CreateFlowRequestTriggerConfig**](CreateFlowRequestTriggerConfig.md) |  |  |
|**sourceFlowConfig** | [**CreateFlowRequestSourceFlowConfig**](CreateFlowRequestSourceFlowConfig.md) |  |  |
|**destinationFlowConfigList** | [**List&lt;DestinationFlowConfig&gt;**](DestinationFlowConfig.md) |  The configuration that controls how Amazon AppFlow transfers data to the destination connector.  |  |
|**tasks** | [**List&lt;Task&gt;**](Task.md) |  A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.  |  |
|**metadataCatalogConfig** | [**CreateFlowRequestMetadataCatalogConfig**](CreateFlowRequestMetadataCatalogConfig.md) |  |  [optional] |
|**clientToken** | **String** | &lt;p&gt;The &lt;code&gt;clientToken&lt;/code&gt; parameter is an idempotency token. It ensures that your &lt;code&gt;UpdateFlow&lt;/code&gt; request completes only once. You choose the value to pass. For example, if you don&#39;t receive a response from your request, you can safely retry the request with the same &lt;code&gt;clientToken&lt;/code&gt; parameter value.&lt;/p&gt; &lt;p&gt;If you omit a &lt;code&gt;clientToken&lt;/code&gt; value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.&lt;/p&gt; &lt;p&gt;If you specify input parameters that differ from your first request, an error occurs. If you use a different value for &lt;code&gt;clientToken&lt;/code&gt;, Amazon AppFlow considers it a new call to &lt;code&gt;UpdateFlow&lt;/code&gt;. The token is active for 8 hours.&lt;/p&gt; |  [optional] |



