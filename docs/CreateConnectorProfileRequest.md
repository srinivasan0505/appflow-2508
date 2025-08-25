

# CreateConnectorProfileRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**connectorProfileName** | **String** |  The name of the connector profile. The name is unique for each &lt;code&gt;ConnectorProfile&lt;/code&gt; in your Amazon Web Services account.  |  |
|**kmsArn** | **String** |  The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption. This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don&#39;t provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.  |  [optional] |
|**connectorType** | [**ConnectorTypeEnum**](#ConnectorTypeEnum) |  The type of connector, such as Salesforce, Amplitude, and so on.  |  |
|**connectorLabel** | **String** | The label of the connector. The label is unique for each &lt;code&gt;ConnectorRegistration&lt;/code&gt; in your Amazon Web Services account. Only needed if calling for CUSTOMCONNECTOR connector type/. |  [optional] |
|**connectionMode** | [**ConnectionModeEnum**](#ConnectionModeEnum) |  Indicates the connection mode and specifies whether it is public or private. Private flows use Amazon Web Services PrivateLink to route data over Amazon Web Services infrastructure without exposing it to the public internet.  |  |
|**connectorProfileConfig** | [**CreateConnectorProfileRequestConnectorProfileConfig**](CreateConnectorProfileRequestConnectorProfileConfig.md) |  |  |
|**clientToken** | **String** | &lt;p&gt;The &lt;code&gt;clientToken&lt;/code&gt; parameter is an idempotency token. It ensures that your &lt;code&gt;CreateConnectorProfile&lt;/code&gt; request completes only once. You choose the value to pass. For example, if you don&#39;t receive a response from your request, you can safely retry the request with the same &lt;code&gt;clientToken&lt;/code&gt; parameter value.&lt;/p&gt; &lt;p&gt;If you omit a &lt;code&gt;clientToken&lt;/code&gt; value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.&lt;/p&gt; &lt;p&gt;If you specify input parameters that differ from your first request, an error occurs. If you use a different value for &lt;code&gt;clientToken&lt;/code&gt;, Amazon AppFlow considers it a new call to &lt;code&gt;CreateConnectorProfile&lt;/code&gt;. The token is active for 8 hours.&lt;/p&gt; |  [optional] |



## Enum: ConnectorTypeEnum

| Name | Value |
|---- | -----|
| SALESFORCE | &quot;Salesforce&quot; |
| SINGULAR | &quot;Singular&quot; |
| SLACK | &quot;Slack&quot; |
| REDSHIFT | &quot;Redshift&quot; |
| S3 | &quot;S3&quot; |
| MARKETO | &quot;Marketo&quot; |
| GOOGLEANALYTICS | &quot;Googleanalytics&quot; |
| ZENDESK | &quot;Zendesk&quot; |
| SERVICENOW | &quot;Servicenow&quot; |
| DATADOG | &quot;Datadog&quot; |
| TRENDMICRO | &quot;Trendmicro&quot; |
| SNOWFLAKE | &quot;Snowflake&quot; |
| DYNATRACE | &quot;Dynatrace&quot; |
| INFORNEXUS | &quot;Infornexus&quot; |
| AMPLITUDE | &quot;Amplitude&quot; |
| VEEVA | &quot;Veeva&quot; |
| EVENT_BRIDGE | &quot;EventBridge&quot; |
| LOOKOUT_METRICS | &quot;LookoutMetrics&quot; |
| UPSOLVER | &quot;Upsolver&quot; |
| HONEYCODE | &quot;Honeycode&quot; |
| CUSTOMER_PROFILES | &quot;CustomerProfiles&quot; |
| SAPO_DATA | &quot;SAPOData&quot; |
| CUSTOM_CONNECTOR | &quot;CustomConnector&quot; |
| PARDOT | &quot;Pardot&quot; |



## Enum: ConnectionModeEnum

| Name | Value |
|---- | -----|
| PUBLIC | &quot;Public&quot; |
| PRIVATE | &quot;Private&quot; |



