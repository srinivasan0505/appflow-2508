

# DescribeConnectorEntityRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**connectorEntityName** | **String** |  The entity name for that connector.  |  |
|**connectorType** | [**ConnectorTypeEnum**](#ConnectorTypeEnum) |  The type of connector application, such as Salesforce, Amplitude, and so on.  |  [optional] |
|**connectorProfileName** | **String** |  The name of the connector profile. The name is unique for each &lt;code&gt;ConnectorProfile&lt;/code&gt; in the Amazon Web Services account.  |  [optional] |
|**apiVersion** | **String** | The version of the API that&#39;s used by the connector. |  [optional] |



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



