

# DescribeConnectorProfilesRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**connectorProfileNames** | **List&lt;String&gt;** |  The name of the connector profile. The name is unique for each &lt;code&gt;ConnectorProfile&lt;/code&gt; in the Amazon Web Services account.  |  [optional] |
|**connectorType** | [**ConnectorTypeEnum**](#ConnectorTypeEnum) |  The type of connector, such as Salesforce, Amplitude, and so on.  |  [optional] |
|**connectorLabel** | **String** | The name of the connector. The name is unique for each &lt;code&gt;ConnectorRegistration&lt;/code&gt; in your Amazon Web Services account. Only needed if calling for CUSTOMCONNECTOR connector type/. |  [optional] |
|**maxResults** | **Integer** |  Specifies the maximum number of items that should be returned in the result set. The default for &lt;code&gt;maxResults&lt;/code&gt; is 20 (for all paginated API operations).  |  [optional] |
|**nextToken** | **String** |  The pagination token for the next page of data.  |  [optional] |



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



