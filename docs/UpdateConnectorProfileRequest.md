

# UpdateConnectorProfileRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**connectorProfileName** | **String** |  The name of the connector profile and is unique for each &lt;code&gt;ConnectorProfile&lt;/code&gt; in the Amazon Web Services account.  |  |
|**connectionMode** | [**ConnectionModeEnum**](#ConnectionModeEnum) |  Indicates the connection mode and if it is public or private.  |  |
|**connectorProfileConfig** | [**CreateConnectorProfileRequestConnectorProfileConfig**](CreateConnectorProfileRequestConnectorProfileConfig.md) |  |  |
|**clientToken** | **String** | &lt;p&gt;The &lt;code&gt;clientToken&lt;/code&gt; parameter is an idempotency token. It ensures that your &lt;code&gt;UpdateConnectorProfile&lt;/code&gt; request completes only once. You choose the value to pass. For example, if you don&#39;t receive a response from your request, you can safely retry the request with the same &lt;code&gt;clientToken&lt;/code&gt; parameter value.&lt;/p&gt; &lt;p&gt;If you omit a &lt;code&gt;clientToken&lt;/code&gt; value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.&lt;/p&gt; &lt;p&gt;If you specify input parameters that differ from your first request, an error occurs. If you use a different value for &lt;code&gt;clientToken&lt;/code&gt;, Amazon AppFlow considers it a new call to &lt;code&gt;UpdateConnectorProfile&lt;/code&gt;. The token is active for 8 hours.&lt;/p&gt; |  [optional] |



## Enum: ConnectionModeEnum

| Name | Value |
|---- | -----|
| PUBLIC | &quot;Public&quot; |
| PRIVATE | &quot;Private&quot; |



