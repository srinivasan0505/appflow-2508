

# ListConnectorEntitiesRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**connectorProfileName** | **String** |  The name of the connector profile. The name is unique for each &lt;code&gt;ConnectorProfile&lt;/code&gt; in the Amazon Web Services account, and is used to query the downstream connector.  |  [optional] |
|**connectorType** | [**ConnectorTypeEnum**](#ConnectorTypeEnum) |  The type of connector, such as Salesforce, Amplitude, and so on.  |  [optional] |
|**entitiesPath** | **String** |  This optional parameter is specific to connector implementation. Some connectors support multiple levels or categories of entities. You can find out the list of roots for such providers by sending a request without the &lt;code&gt;entitiesPath&lt;/code&gt; parameter. If the connector supports entities at different roots, this initial request returns the list of roots. Otherwise, this request returns all entities supported by the provider.  |  [optional] |
|**apiVersion** | **String** | The version of the API that&#39;s used by the connector. |  [optional] |
|**maxResults** | **Integer** | The maximum number of items that the operation returns in the response. |  [optional] |
|**nextToken** | **String** | A token that was provided by your prior &lt;code&gt;ListConnectorEntities&lt;/code&gt; operation if the response was too big for the page size. You specify this token to get the next page of results in paginated response. |  [optional] |



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



