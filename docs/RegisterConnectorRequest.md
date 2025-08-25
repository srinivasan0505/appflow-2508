

# RegisterConnectorRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**connectorLabel** | **String** |  The name of the connector. The name is unique for each &lt;code&gt;ConnectorRegistration&lt;/code&gt; in your Amazon Web Services account. |  [optional] |
|**description** | **String** | A description about the connector that&#39;s being registered. |  [optional] |
|**connectorProvisioningType** | [**ConnectorProvisioningTypeEnum**](#ConnectorProvisioningTypeEnum) | The type of provisioning that the connector supports, such as Lambda. |  [optional] |
|**connectorProvisioningConfig** | [**RegisterConnectorRequestConnectorProvisioningConfig**](RegisterConnectorRequestConnectorProvisioningConfig.md) |  |  [optional] |
|**clientToken** | **String** | &lt;p&gt;The &lt;code&gt;clientToken&lt;/code&gt; parameter is an idempotency token. It ensures that your &lt;code&gt;RegisterConnector&lt;/code&gt; request completes only once. You choose the value to pass. For example, if you don&#39;t receive a response from your request, you can safely retry the request with the same &lt;code&gt;clientToken&lt;/code&gt; parameter value.&lt;/p&gt; &lt;p&gt;If you omit a &lt;code&gt;clientToken&lt;/code&gt; value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.&lt;/p&gt; &lt;p&gt;If you specify input parameters that differ from your first request, an error occurs. If you use a different value for &lt;code&gt;clientToken&lt;/code&gt;, Amazon AppFlow considers it a new call to &lt;code&gt;RegisterConnector&lt;/code&gt;. The token is active for 8 hours.&lt;/p&gt; |  [optional] |



## Enum: ConnectorProvisioningTypeEnum

| Name | Value |
|---- | -----|
| LAMBDA | &quot;LAMBDA&quot; |



