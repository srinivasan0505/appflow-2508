

# ResetConnectorMetadataCacheRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**connectorProfileName** | **String** | &lt;p&gt;The name of the connector profile that you want to reset cached metadata for.&lt;/p&gt; &lt;p&gt;You can omit this parameter if you&#39;re resetting the cache for any of the following connectors: Amazon Connect, Amazon EventBridge, Amazon Lookout for Metrics, Amazon S3, or Upsolver. If you&#39;re resetting the cache for any other connector, you must include this parameter in your request.&lt;/p&gt; |  [optional] |
|**connectorType** | [**ConnectorTypeEnum**](#ConnectorTypeEnum) | &lt;p&gt;The type of connector to reset cached metadata for.&lt;/p&gt; &lt;p&gt;You must include this parameter in your request if you&#39;re resetting the cache for any of the following connectors: Amazon Connect, Amazon EventBridge, Amazon Lookout for Metrics, Amazon S3, or Upsolver. If you&#39;re resetting the cache for any other connector, you can omit this parameter from your request. &lt;/p&gt; |  [optional] |
|**connectorEntityName** | **String** | &lt;p&gt;Use this parameter if you want to reset cached metadata about the details for an individual entity.&lt;/p&gt; &lt;p&gt;If you don&#39;t include this parameter in your request, Amazon AppFlow only resets cached metadata about entity names, not entity details.&lt;/p&gt; |  [optional] |
|**entitiesPath** | **String** | &lt;p&gt;Use this parameter only if you’re resetting the cached metadata about a nested entity. Only some connectors support nested entities. A nested entity is one that has another entity as a parent. To use this parameter, specify the name of the parent entity.&lt;/p&gt; &lt;p&gt;To look up the parent-child relationship of entities, you can send a ListConnectorEntities request that omits the entitiesPath parameter. Amazon AppFlow will return a list of top-level entities. For each one, it indicates whether the entity has nested entities. Then, in a subsequent ListConnectorEntities request, you can specify a parent entity name for the entitiesPath parameter. Amazon AppFlow will return a list of the child entities for that parent.&lt;/p&gt; |  [optional] |
|**apiVersion** | **String** | &lt;p&gt;The API version that you specified in the connector profile that you’re resetting cached metadata for. You must use this parameter only if the connector supports multiple API versions or if the connector type is CustomConnector.&lt;/p&gt; &lt;p&gt;To look up how many versions a connector supports, use the DescribeConnectors action. In the response, find the value that Amazon AppFlow returns for the connectorVersion parameter.&lt;/p&gt; &lt;p&gt;To look up the connector type, use the DescribeConnectorProfiles action. In the response, find the value that Amazon AppFlow returns for the connectorType parameter.&lt;/p&gt; &lt;p&gt;To look up the API version that you specified in a connector profile, use the DescribeConnectorProfiles action.&lt;/p&gt; |  [optional] |



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



