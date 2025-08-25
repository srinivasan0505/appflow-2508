

# StartFlowRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**flowName** | **String** |  The specified name of the flow. Spaces are not allowed. Use underscores (_) or hyphens (-) only.  |  |
|**clientToken** | **String** | &lt;p&gt;The &lt;code&gt;clientToken&lt;/code&gt; parameter is an idempotency token. It ensures that your &lt;code&gt;StartFlow&lt;/code&gt; request completes only once. You choose the value to pass. For example, if you don&#39;t receive a response from your request, you can safely retry the request with the same &lt;code&gt;clientToken&lt;/code&gt; parameter value.&lt;/p&gt; &lt;p&gt;If you omit a &lt;code&gt;clientToken&lt;/code&gt; value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.&lt;/p&gt; &lt;p&gt;If you specify input parameters that differ from your first request, an error occurs for flows that run on a schedule or based on an event. However, the error doesn&#39;t occur for flows that run on demand. You set the conditions that initiate your flow for the &lt;code&gt;triggerConfig&lt;/code&gt; parameter.&lt;/p&gt; &lt;p&gt;If you use a different value for &lt;code&gt;clientToken&lt;/code&gt;, Amazon AppFlow considers it a new call to &lt;code&gt;StartFlow&lt;/code&gt;. The token is active for 8 hours.&lt;/p&gt; |  [optional] |



