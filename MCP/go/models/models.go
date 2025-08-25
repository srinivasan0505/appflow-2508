package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// FlowDefinition represents the FlowDefinition schema from the OpenAPI specification
type FlowDefinition struct {
	Flowstatus interface{} `json:"flowStatus,omitempty"`
	Sourceconnectortype interface{} `json:"sourceConnectorType,omitempty"`
	Destinationconnectortype interface{} `json:"destinationConnectorType,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Flowname interface{} `json:"flowName,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Lastupdatedby interface{} `json:"lastUpdatedBy,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Lastrunexecutiondetails DescribeFlowResponselastRunExecutionDetails `json:"lastRunExecutionDetails,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Sourceconnectorlabel interface{} `json:"sourceConnectorLabel,omitempty"`
	Triggertype interface{} `json:"triggerType,omitempty"`
	Destinationconnectorlabel interface{} `json:"destinationConnectorLabel,omitempty"`
	Flowarn interface{} `json:"flowArn,omitempty"`
}

// IncrementalPullConfig represents the IncrementalPullConfig schema from the OpenAPI specification
type IncrementalPullConfig struct {
	Datetimetypefieldname interface{} `json:"datetimeTypeFieldName,omitempty"`
}

// UnregisterConnectorrequest represents the UnregisterConnectorrequest schema from the OpenAPI specification
type UnregisterConnectorrequest struct {
	Connectorlabel string `json:"connectorLabel"` // The label of the connector. The label is unique for each <code>ConnectorRegistration</code> in your Amazon Web Services account.
	Forcedelete bool `json:"forceDelete,omitempty"` // Indicates whether Amazon AppFlow should unregister the connector, even if it is currently in use in one or more connector profiles. The default value is false.
}

// ConnectorConfigurationconnectorProvisioningConfig represents the ConnectorConfigurationconnectorProvisioningConfig schema from the OpenAPI specification
type ConnectorConfigurationconnectorProvisioningConfig struct {
	Lambda RegisterConnectorrequestconnectorProvisioningConfiglambda `json:"lambda,omitempty"`
}

// FieldTypeDetails represents the FieldTypeDetails schema from the OpenAPI specification
type FieldTypeDetails struct {
	Fieldtype interface{} `json:"fieldType"`
	Fieldvaluerange FieldTypeDetailsfieldValueRange `json:"fieldValueRange,omitempty"`
	Filteroperators interface{} `json:"filterOperators"`
	Supporteddateformat interface{} `json:"supportedDateFormat,omitempty"`
	Supportedvalues interface{} `json:"supportedValues,omitempty"`
	Valueregexpattern interface{} `json:"valueRegexPattern,omitempty"`
	Fieldlengthrange FieldTypeDetailsfieldLengthRange `json:"fieldLengthRange,omitempty"`
}

// SAPODataConnectorProfileProperties represents the SAPODataConnectorProfileProperties schema from the OpenAPI specification
type SAPODataConnectorProfileProperties struct {
	Disablesso interface{} `json:"disableSSO,omitempty"`
	Logonlanguage interface{} `json:"logonLanguage,omitempty"`
	Oauthproperties SAPODataConnectorProfilePropertiesoAuthProperties `json:"oAuthProperties,omitempty"`
	Portnumber interface{} `json:"portNumber"`
	Privatelinkservicename interface{} `json:"privateLinkServiceName,omitempty"`
	Applicationhosturl interface{} `json:"applicationHostUrl"`
	Applicationservicepath interface{} `json:"applicationServicePath"`
	Clientnumber interface{} `json:"clientNumber"`
}

// SourceConnectorPropertiesMarketo represents the SourceConnectorPropertiesMarketo schema from the OpenAPI specification
type SourceConnectorPropertiesMarketo struct {
	Object interface{} `json:"object"`
}

// CreateFlowrequest represents the CreateFlowrequest schema from the OpenAPI specification
type CreateFlowrequest struct {
	Tasks []Task `json:"tasks"` // A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.
	Triggerconfig CreateFlowrequesttriggerConfig `json:"triggerConfig"` // The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
	Clienttoken string `json:"clientToken,omitempty"` // <p>The <code>clientToken</code> parameter is an idempotency token. It ensures that your <code>CreateFlow</code> request completes only once. You choose the value to pass. For example, if you don't receive a response from your request, you can safely retry the request with the same <code>clientToken</code> parameter value.</p> <p>If you omit a <code>clientToken</code> value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.</p> <p>If you specify input parameters that differ from your first request, an error occurs. If you use a different value for <code>clientToken</code>, Amazon AppFlow considers it a new call to <code>CreateFlow</code>. The token is active for 8 hours.</p>
	Destinationflowconfiglist []DestinationFlowConfig `json:"destinationFlowConfigList"` // The configuration that controls how Amazon AppFlow places data in the destination connector.
	Flowname string `json:"flowName"` // The specified name of the flow. Spaces are not allowed. Use underscores (_) or hyphens (-) only.
	Tags map[string]interface{} `json:"tags,omitempty"` // The tags used to organize, track, or control access for your flow.
	Kmsarn string `json:"kmsArn,omitempty"` // The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption. This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	Description string `json:"description,omitempty"` // A description of the flow you want to create.
	Metadatacatalogconfig CreateFlowrequestmetadataCatalogConfig `json:"metadataCatalogConfig,omitempty"` // Specifies the configuration that Amazon AppFlow uses when it catalogs your data. When Amazon AppFlow catalogs your data, it stores metadata in a data catalog.
	Sourceflowconfig CreateFlowrequestsourceFlowConfig `json:"sourceFlowConfig"` // Contains information about the configuration of the source connector used in the flow.
}

// Task represents the Task schema from the OpenAPI specification
type Task struct {
	Destinationfield interface{} `json:"destinationField,omitempty"`
	Sourcefields interface{} `json:"sourceFields"`
	Taskproperties interface{} `json:"taskProperties,omitempty"`
	Tasktype interface{} `json:"taskType"`
	Connectoroperator TaskconnectorOperator `json:"connectorOperator,omitempty"`
}

// DeleteConnectorProfilerequest represents the DeleteConnectorProfilerequest schema from the OpenAPI specification
type DeleteConnectorProfilerequest struct {
	Connectorprofilename string `json:"connectorProfileName"` // The name of the connector profile. The name is unique for each <code>ConnectorProfile</code> in your account.
	Forcedelete bool `json:"forceDelete,omitempty"` // Indicates whether Amazon AppFlow should delete the profile, even if it is currently in use in one or more flows.
}

// ConnectorEntity represents the ConnectorEntity schema from the OpenAPI specification
type ConnectorEntity struct {
	Hasnestedentities interface{} `json:"hasNestedEntities,omitempty"`
	Label interface{} `json:"label,omitempty"`
	Name interface{} `json:"name"`
}

// DeleteConnectorProfileResponse represents the DeleteConnectorProfileResponse schema from the OpenAPI specification
type DeleteConnectorProfileResponse struct {
}

// TrendmicroMetadata represents the TrendmicroMetadata schema from the OpenAPI specification
type TrendmicroMetadata struct {
}

// ConnectorProfileCredentialsTrendmicro represents the ConnectorProfileCredentialsTrendmicro schema from the OpenAPI specification
type ConnectorProfileCredentialsTrendmicro struct {
	Apisecretkey interface{} `json:"apiSecretKey"`
}

// CreateFlowrequestsourceFlowConfigincrementalPullConfig represents the CreateFlowrequestsourceFlowConfigincrementalPullConfig schema from the OpenAPI specification
type CreateFlowrequestsourceFlowConfigincrementalPullConfig struct {
	Datetimetypefieldname interface{} `json:"datetimeTypeFieldName,omitempty"`
}

// SourceConnectorPropertiesZendesk represents the SourceConnectorPropertiesZendesk schema from the OpenAPI specification
type SourceConnectorPropertiesZendesk struct {
	Object interface{} `json:"object"`
}

// SnowflakeConnectorProfileCredentials represents the SnowflakeConnectorProfileCredentials schema from the OpenAPI specification
type SnowflakeConnectorProfileCredentials struct {
	Username interface{} `json:"username"`
	Password interface{} `json:"password"`
}

// ProfilePropertiesMap represents the ProfilePropertiesMap schema from the OpenAPI specification
type ProfilePropertiesMap struct {
}

// CreateConnectorProfileRequestconnectorProfileConfig represents the CreateConnectorProfileRequestconnectorProfileConfig schema from the OpenAPI specification
type CreateConnectorProfileRequestconnectorProfileConfig struct {
	Connectorprofileproperties CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileProperties `json:"connectorProfileProperties"`
	Connectorprofilecredentials CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileCredentials `json:"connectorProfileCredentials,omitempty"`
}

// CustomConnectorSourceProperties represents the CustomConnectorSourceProperties schema from the OpenAPI specification
type CustomConnectorSourceProperties struct {
	Entityname interface{} `json:"entityName"`
	Customproperties interface{} `json:"customProperties,omitempty"`
	Datatransferapi CustomConnectorSourcePropertiesdataTransferApi `json:"dataTransferApi,omitempty"`
}

// ConnectorProfileCredentialsSalesforce represents the ConnectorProfileCredentialsSalesforce schema from the OpenAPI specification
type ConnectorProfileCredentialsSalesforce struct {
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientcredentialsarn interface{} `json:"clientCredentialsArn,omitempty"`
	Jwttoken interface{} `json:"jwtToken,omitempty"`
	Oauth2granttype interface{} `json:"oAuth2GrantType,omitempty"`
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
}

// HoneycodeMetadata represents the HoneycodeMetadata schema from the OpenAPI specification
type HoneycodeMetadata struct {
	Oauthscopes interface{} `json:"oAuthScopes,omitempty"`
}

// DestinationFlowConfigdestinationConnectorProperties represents the DestinationFlowConfigdestinationConnectorProperties schema from the OpenAPI specification
type DestinationFlowConfigdestinationConnectorProperties struct {
	S3 DestinationConnectorPropertiesS3 `json:"S3,omitempty"`
	Upsolver DestinationConnectorPropertiesUpsolver `json:"Upsolver,omitempty"`
	Customerprofiles DestinationConnectorPropertiesCustomerProfiles `json:"CustomerProfiles,omitempty"`
	Honeycode DestinationConnectorPropertiesHoneycode `json:"Honeycode,omitempty"`
	Sapodata DestinationConnectorPropertiesSAPOData `json:"SAPOData,omitempty"`
	Zendesk DestinationConnectorPropertiesZendesk `json:"Zendesk,omitempty"`
	Marketo DestinationConnectorPropertiesMarketo `json:"Marketo,omitempty"`
	Salesforce DestinationConnectorPropertiesSalesforce `json:"Salesforce,omitempty"`
	Snowflake DestinationConnectorPropertiesSnowflake `json:"Snowflake,omitempty"`
	Lookoutmetrics interface{} `json:"LookoutMetrics,omitempty"`
	Customconnector DestinationConnectorPropertiesCustomConnector `json:"CustomConnector,omitempty"`
	Eventbridge DestinationConnectorPropertiesEventBridge `json:"EventBridge,omitempty"`
	Redshift DestinationConnectorPropertiesRedshift `json:"Redshift,omitempty"`
}

// RegisterConnectorRequest represents the RegisterConnectorRequest schema from the OpenAPI specification
type RegisterConnectorRequest struct {
	Connectorlabel interface{} `json:"connectorLabel,omitempty"`
	Connectorprovisioningconfig RegisterConnectorRequestconnectorProvisioningConfig `json:"connectorProvisioningConfig,omitempty"`
	Connectorprovisioningtype interface{} `json:"connectorProvisioningType,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// SourceConnectorPropertiesSalesforce represents the SourceConnectorPropertiesSalesforce schema from the OpenAPI specification
type SourceConnectorPropertiesSalesforce struct {
	Enabledynamicfieldupdate interface{} `json:"enableDynamicFieldUpdate,omitempty"`
	Includedeletedrecords interface{} `json:"includeDeletedRecords,omitempty"`
	Object interface{} `json:"object"`
	Datatransferapi interface{} `json:"dataTransferApi,omitempty"`
}

// TrendmicroSourceProperties represents the TrendmicroSourceProperties schema from the OpenAPI specification
type TrendmicroSourceProperties struct {
	Object interface{} `json:"object"`
}

// DestinationConnectorPropertiesSalesforce represents the DestinationConnectorPropertiesSalesforce schema from the OpenAPI specification
type DestinationConnectorPropertiesSalesforce struct {
	Errorhandlingconfig SalesforceDestinationPropertieserrorHandlingConfig `json:"errorHandlingConfig,omitempty"`
	Idfieldnames interface{} `json:"idFieldNames,omitempty"`
	Object interface{} `json:"object"`
	Writeoperationtype interface{} `json:"writeOperationType,omitempty"`
	Datatransferapi interface{} `json:"dataTransferApi,omitempty"`
}

// ConnectorProfileprivateConnectionProvisioningState represents the ConnectorProfileprivateConnectionProvisioningState schema from the OpenAPI specification
type ConnectorProfileprivateConnectionProvisioningState struct {
	Failuremessage interface{} `json:"failureMessage,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Failurecause interface{} `json:"failureCause,omitempty"`
}

// MarketoMetadata represents the MarketoMetadata schema from the OpenAPI specification
type MarketoMetadata struct {
}

// ConnectorEntityFieldsourceProperties represents the ConnectorEntityFieldsourceProperties schema from the OpenAPI specification
type ConnectorEntityFieldsourceProperties struct {
	Isqueryable interface{} `json:"isQueryable,omitempty"`
	Isretrievable interface{} `json:"isRetrievable,omitempty"`
	Istimestampfieldforincrementalqueries interface{} `json:"isTimestampFieldForIncrementalQueries,omitempty"`
}

// ConnectorMetadataGoogleAnalytics represents the ConnectorMetadataGoogleAnalytics schema from the OpenAPI specification
type ConnectorMetadataGoogleAnalytics struct {
	Oauthscopes interface{} `json:"oAuthScopes,omitempty"`
}

// ResetConnectorMetadataCacheRequest represents the ResetConnectorMetadataCacheRequest schema from the OpenAPI specification
type ResetConnectorMetadataCacheRequest struct {
	Connectorprofilename interface{} `json:"connectorProfileName,omitempty"`
	Connectortype interface{} `json:"connectorType,omitempty"`
	Entitiespath interface{} `json:"entitiesPath,omitempty"`
	Apiversion interface{} `json:"apiVersion,omitempty"`
	Connectorentityname interface{} `json:"connectorEntityName,omitempty"`
}

// CreateFlowrequestmetadataCatalogConfig represents the CreateFlowrequestmetadataCatalogConfig schema from the OpenAPI specification
type CreateFlowrequestmetadataCatalogConfig struct {
	Gluedatacatalog CreateFlowrequestmetadataCatalogConfigglueDataCatalog `json:"glueDataCatalog,omitempty"`
}

// ConnectorProfilePropertiesSnowflake represents the ConnectorProfilePropertiesSnowflake schema from the OpenAPI specification
type ConnectorProfilePropertiesSnowflake struct {
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Privatelinkservicename interface{} `json:"privateLinkServiceName,omitempty"`
	Region interface{} `json:"region,omitempty"`
	Stage interface{} `json:"stage"`
	Warehouse interface{} `json:"warehouse"`
	Accountname interface{} `json:"accountName,omitempty"`
	Bucketname interface{} `json:"bucketName"`
}

// SAPODataConnectorProfileCredentials represents the SAPODataConnectorProfileCredentials schema from the OpenAPI specification
type SAPODataConnectorProfileCredentials struct {
	Basicauthcredentials SAPODataConnectorProfileCredentialsbasicAuthCredentials `json:"basicAuthCredentials,omitempty"`
	Oauthcredentials SAPODataConnectorProfileCredentialsoAuthCredentials `json:"oAuthCredentials,omitempty"`
}

// ConnectorProfileCredentials represents the ConnectorProfileCredentials schema from the OpenAPI specification
type ConnectorProfileCredentials struct {
	Infornexus ConnectorProfileCredentialsInforNexus `json:"InforNexus,omitempty"`
	Amplitude ConnectorProfileCredentialsAmplitude `json:"Amplitude,omitempty"`
	Sapodata SAPODataConnectorProfileCredentials `json:"SAPOData,omitempty"` // The connector-specific profile credentials required when using SAPOData.
	Veeva ConnectorProfileCredentialsVeeva `json:"Veeva,omitempty"`
	Honeycode ConnectorProfileCredentialsHoneycode `json:"Honeycode,omitempty"`
	Trendmicro ConnectorProfileCredentialsTrendmicro `json:"Trendmicro,omitempty"`
	Salesforce ConnectorProfileCredentialsSalesforce `json:"Salesforce,omitempty"`
	Pardot ConnectorProfileCredentialsPardot `json:"Pardot,omitempty"`
	Dynatrace ConnectorProfileCredentialsDynatrace `json:"Dynatrace,omitempty"`
	Customconnector CustomConnectorProfileCredentials `json:"CustomConnector,omitempty"` // The connector-specific profile credentials that are required when using the custom connector.
	Servicenow ConnectorProfileCredentialsServiceNow `json:"ServiceNow,omitempty"`
	Slack ConnectorProfileCredentialsSlack `json:"Slack,omitempty"`
	Marketo ConnectorProfileCredentialsMarketo `json:"Marketo,omitempty"`
	Datadog ConnectorProfileCredentialsDatadog `json:"Datadog,omitempty"`
	Googleanalytics ConnectorProfileCredentialsGoogleAnalytics `json:"GoogleAnalytics,omitempty"`
	Zendesk ConnectorProfileCredentialsZendesk `json:"Zendesk,omitempty"`
	Redshift ConnectorProfileCredentialsRedshift `json:"Redshift,omitempty"`
	Singular ConnectorProfileCredentialsSingular `json:"Singular,omitempty"`
	Snowflake ConnectorProfileCredentialsSnowflake `json:"Snowflake,omitempty"`
}

// DescribeFlowRequest represents the DescribeFlowRequest schema from the OpenAPI specification
type DescribeFlowRequest struct {
	Flowname interface{} `json:"flowName"`
}

// MarketoSourceProperties represents the MarketoSourceProperties schema from the OpenAPI specification
type MarketoSourceProperties struct {
	Object interface{} `json:"object"`
}

// VeevaConnectorProfileProperties represents the VeevaConnectorProfileProperties schema from the OpenAPI specification
type VeevaConnectorProfileProperties struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// DatadogSourceProperties represents the DatadogSourceProperties schema from the OpenAPI specification
type DatadogSourceProperties struct {
	Object interface{} `json:"object"`
}

// DescribeConnectorRequest represents the DescribeConnectorRequest schema from the OpenAPI specification
type DescribeConnectorRequest struct {
	Connectorlabel interface{} `json:"connectorLabel,omitempty"`
	Connectortype interface{} `json:"connectorType"`
}

// DataTransferApi represents the DataTransferApi schema from the OpenAPI specification
type DataTransferApi struct {
	Name interface{} `json:"Name,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// TagResourcerequest represents the TagResourcerequest schema from the OpenAPI specification
type TagResourcerequest struct {
	Tags map[string]interface{} `json:"tags"` // The tags used to organize, track, or control access for your flow.
}

// SourceConnectorPropertiesPardot represents the SourceConnectorPropertiesPardot schema from the OpenAPI specification
type SourceConnectorPropertiesPardot struct {
	Object interface{} `json:"object"`
}

// CustomProperties represents the CustomProperties schema from the OpenAPI specification
type CustomProperties struct {
}

// SlackConnectorProfileProperties represents the SlackConnectorProfileProperties schema from the OpenAPI specification
type SlackConnectorProfileProperties struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// DescribeConnectorsrequest represents the DescribeConnectorsrequest schema from the OpenAPI specification
type DescribeConnectorsrequest struct {
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that should be returned in the result set. The default is 20.
	Nexttoken string `json:"nextToken,omitempty"` // The pagination token for the next page of data.
	Connectortypes []string `json:"connectorTypes,omitempty"` // The type of connector, such as Salesforce, Amplitude, and so on.
}

// TrendmicroConnectorProfileCredentials represents the TrendmicroConnectorProfileCredentials schema from the OpenAPI specification
type TrendmicroConnectorProfileCredentials struct {
	Apisecretkey interface{} `json:"apiSecretKey"`
}

// ConnectorProfilePropertiesDatadog represents the ConnectorProfilePropertiesDatadog schema from the OpenAPI specification
type ConnectorProfilePropertiesDatadog struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// SourceConnectorPropertiesServiceNow represents the SourceConnectorPropertiesServiceNow schema from the OpenAPI specification
type SourceConnectorPropertiesServiceNow struct {
	Object interface{} `json:"object"`
}

// DescribeConnectorEntityResponse represents the DescribeConnectorEntityResponse schema from the OpenAPI specification
type DescribeConnectorEntityResponse struct {
	Connectorentityfields interface{} `json:"connectorEntityFields"`
}

// SAPODataDestinationPropertiessuccessResponseHandlingConfig represents the SAPODataDestinationPropertiessuccessResponseHandlingConfig schema from the OpenAPI specification
type SAPODataDestinationPropertiessuccessResponseHandlingConfig struct {
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Bucketname interface{} `json:"bucketName,omitempty"`
}

// RegisterConnectorrequest represents the RegisterConnectorrequest schema from the OpenAPI specification
type RegisterConnectorrequest struct {
	Connectorprovisioningtype string `json:"connectorProvisioningType,omitempty"` // The type of provisioning that the connector supports, such as Lambda.
	Description string `json:"description,omitempty"` // A description about the connector that's being registered.
	Clienttoken string `json:"clientToken,omitempty"` // <p>The <code>clientToken</code> parameter is an idempotency token. It ensures that your <code>RegisterConnector</code> request completes only once. You choose the value to pass. For example, if you don't receive a response from your request, you can safely retry the request with the same <code>clientToken</code> parameter value.</p> <p>If you omit a <code>clientToken</code> value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.</p> <p>If you specify input parameters that differ from your first request, an error occurs. If you use a different value for <code>clientToken</code>, Amazon AppFlow considers it a new call to <code>RegisterConnector</code>. The token is active for 8 hours.</p>
	Connectorlabel string `json:"connectorLabel,omitempty"` // The name of the connector. The name is unique for each <code>ConnectorRegistration</code> in your Amazon Web Services account.
	Connectorprovisioningconfig RegisterConnectorrequestconnectorProvisioningConfig `json:"connectorProvisioningConfig,omitempty"` // Contains information about the configuration of the connector being registered.
}

// SnowflakeMetadata represents the SnowflakeMetadata schema from the OpenAPI specification
type SnowflakeMetadata struct {
	Supportedregions interface{} `json:"supportedRegions,omitempty"`
}

// DescribeFlowExecutionRecordsRequest represents the DescribeFlowExecutionRecordsRequest schema from the OpenAPI specification
type DescribeFlowExecutionRecordsRequest struct {
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Flowname interface{} `json:"flowName"`
}

// AuthenticationConfigoAuth2Defaults represents the AuthenticationConfigoAuth2Defaults schema from the OpenAPI specification
type AuthenticationConfigoAuth2Defaults struct {
	Oauthscopes interface{} `json:"oauthScopes,omitempty"`
	Tokenurls interface{} `json:"tokenUrls,omitempty"`
	Authcodeurls interface{} `json:"authCodeUrls,omitempty"`
	Oauth2customproperties interface{} `json:"oauth2CustomProperties,omitempty"`
	Oauth2granttypessupported interface{} `json:"oauth2GrantTypesSupported,omitempty"`
}

// RedshiftDestinationProperties represents the RedshiftDestinationProperties schema from the OpenAPI specification
type RedshiftDestinationProperties struct {
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Errorhandlingconfig RedshiftDestinationPropertieserrorHandlingConfig `json:"errorHandlingConfig,omitempty"`
	Intermediatebucketname interface{} `json:"intermediateBucketName"`
	Object interface{} `json:"object"`
}

// SourceFlowConfig represents the SourceFlowConfig schema from the OpenAPI specification
type SourceFlowConfig struct {
	Sourceconnectorproperties CreateFlowrequestsourceFlowConfigsourceConnectorProperties `json:"sourceConnectorProperties"`
	Apiversion interface{} `json:"apiVersion,omitempty"`
	Connectorprofilename interface{} `json:"connectorProfileName,omitempty"`
	Connectortype interface{} `json:"connectorType"`
	Incrementalpullconfig CreateFlowrequestsourceFlowConfigincrementalPullConfig `json:"incrementalPullConfig,omitempty"`
}

// SalesforceMetadata represents the SalesforceMetadata schema from the OpenAPI specification
type SalesforceMetadata struct {
	Datatransferapis interface{} `json:"dataTransferApis,omitempty"`
	Oauthscopes interface{} `json:"oAuthScopes,omitempty"`
	Oauth2granttypessupported interface{} `json:"oauth2GrantTypesSupported,omitempty"`
}

// DynatraceMetadata represents the DynatraceMetadata schema from the OpenAPI specification
type DynatraceMetadata struct {
}

// SourceConnectorPropertiesInforNexus represents the SourceConnectorPropertiesInforNexus schema from the OpenAPI specification
type SourceConnectorPropertiesInforNexus struct {
	Object interface{} `json:"object"`
}

// MarketoConnectorProfileProperties represents the MarketoConnectorProfileProperties schema from the OpenAPI specification
type MarketoConnectorProfileProperties struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// CreateFlowrequestsourceFlowConfig represents the CreateFlowrequestsourceFlowConfig schema from the OpenAPI specification
type CreateFlowrequestsourceFlowConfig struct {
	Apiversion interface{} `json:"apiVersion,omitempty"`
	Connectorprofilename interface{} `json:"connectorProfileName,omitempty"`
	Connectortype interface{} `json:"connectorType,omitempty"`
	Incrementalpullconfig CreateFlowrequestsourceFlowConfigincrementalPullConfig `json:"incrementalPullConfig,omitempty"`
	Sourceconnectorproperties CreateFlowrequestsourceFlowConfigsourceConnectorProperties `json:"sourceConnectorProperties,omitempty"`
}

// DestinationConnectorPropertiesS3 represents the DestinationConnectorPropertiesS3 schema from the OpenAPI specification
type DestinationConnectorPropertiesS3 struct {
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	S3outputformatconfig S3OutputFormatConfig `json:"s3OutputFormatConfig,omitempty"` // The configuration that determines how Amazon AppFlow should format the flow output data when Amazon S3 is used as the destination.
	Bucketname interface{} `json:"bucketName"`
}

// CreateFlowrequesttriggerConfig represents the CreateFlowrequesttriggerConfig schema from the OpenAPI specification
type CreateFlowrequesttriggerConfig struct {
	Triggertype interface{} `json:"triggerType,omitempty"`
	Triggerproperties CreateFlowrequesttriggerConfigtriggerProperties `json:"triggerProperties,omitempty"`
}

// CustomAuthConfig represents the CustomAuthConfig schema from the OpenAPI specification
type CustomAuthConfig struct {
	Authparameters interface{} `json:"authParameters,omitempty"`
	Customauthenticationtype interface{} `json:"customAuthenticationType,omitempty"`
}

// DestinationConnectorPropertiesSnowflake represents the DestinationConnectorPropertiesSnowflake schema from the OpenAPI specification
type DestinationConnectorPropertiesSnowflake struct {
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Errorhandlingconfig SnowflakeDestinationPropertieserrorHandlingConfig `json:"errorHandlingConfig,omitempty"`
	Intermediatebucketname interface{} `json:"intermediateBucketName"`
	Object interface{} `json:"object"`
}

// SourceConnectorPropertiesGoogleAnalytics represents the SourceConnectorPropertiesGoogleAnalytics schema from the OpenAPI specification
type SourceConnectorPropertiesGoogleAnalytics struct {
	Object interface{} `json:"object"`
}

// ZendeskConnectorProfileCredentials represents the ZendeskConnectorProfileCredentials schema from the OpenAPI specification
type ZendeskConnectorProfileCredentials struct {
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientid interface{} `json:"clientId"`
	Clientsecret interface{} `json:"clientSecret"`
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
}

// CustomConnectorProfileCredentials represents the CustomConnectorProfileCredentials schema from the OpenAPI specification
type CustomConnectorProfileCredentials struct {
	Authenticationtype interface{} `json:"authenticationType"`
	Basic CustomConnectorProfileCredentialsbasic `json:"basic,omitempty"`
	Custom CustomConnectorProfileCredentialscustom `json:"custom,omitempty"`
	Oauth2 CustomConnectorProfileCredentialsoauth2 `json:"oauth2,omitempty"`
	Apikey CustomConnectorProfileCredentialsapiKey `json:"apiKey,omitempty"`
}

// ConnectorOperator represents the ConnectorOperator schema from the OpenAPI specification
type ConnectorOperator struct {
	Customconnector interface{} `json:"CustomConnector,omitempty"`
	Dynatrace interface{} `json:"Dynatrace,omitempty"`
	S3 interface{} `json:"S3,omitempty"`
	Slack interface{} `json:"Slack,omitempty"`
	Zendesk interface{} `json:"Zendesk,omitempty"`
	Infornexus interface{} `json:"InforNexus,omitempty"`
	Salesforce interface{} `json:"Salesforce,omitempty"`
	Servicenow interface{} `json:"ServiceNow,omitempty"`
	Veeva interface{} `json:"Veeva,omitempty"`
	Amplitude interface{} `json:"Amplitude,omitempty"`
	Datadog interface{} `json:"Datadog,omitempty"`
	Pardot interface{} `json:"Pardot,omitempty"`
	Singular interface{} `json:"Singular,omitempty"`
	Googleanalytics interface{} `json:"GoogleAnalytics,omitempty"`
	Marketo interface{} `json:"Marketo,omitempty"`
	Sapodata interface{} `json:"SAPOData,omitempty"`
	Trendmicro interface{} `json:"Trendmicro,omitempty"`
}

// ConnectorProfilePropertiesServiceNow represents the ConnectorProfilePropertiesServiceNow schema from the OpenAPI specification
type ConnectorProfilePropertiesServiceNow struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// ConnectorMetadataHoneycode represents the ConnectorMetadataHoneycode schema from the OpenAPI specification
type ConnectorMetadataHoneycode struct {
	Oauthscopes interface{} `json:"oAuthScopes,omitempty"`
}

// DescribeConnectorResponseconnectorConfiguration represents the DescribeConnectorResponseconnectorConfiguration schema from the OpenAPI specification
type DescribeConnectorResponseconnectorConfiguration struct {
	Isprivatelinkendpointurlrequired interface{} `json:"isPrivateLinkEndpointUrlRequired,omitempty"`
	Supporteddatatransfertypes interface{} `json:"supportedDataTransferTypes,omitempty"`
	Connectorprovisioningconfig ConnectorConfigurationconnectorProvisioningConfig `json:"connectorProvisioningConfig,omitempty"`
	Registeredat interface{} `json:"registeredAt,omitempty"`
	Supportedapiversions interface{} `json:"supportedApiVersions,omitempty"`
	Supportedtriggertypes interface{} `json:"supportedTriggerTypes,omitempty"`
	Authenticationconfig ConnectorConfigurationauthenticationConfig `json:"authenticationConfig,omitempty"`
	Supportedoperators interface{} `json:"supportedOperators,omitempty"`
	Supportedschedulingfrequencies interface{} `json:"supportedSchedulingFrequencies,omitempty"`
	Connectorname interface{} `json:"connectorName,omitempty"`
	Supporteddatatransferapis interface{} `json:"supportedDataTransferApis,omitempty"`
	Connectorprovisioningtype interface{} `json:"connectorProvisioningType,omitempty"`
	Connectorruntimesettings interface{} `json:"connectorRuntimeSettings,omitempty"`
	Registeredby interface{} `json:"registeredBy,omitempty"`
	Supportedwriteoperations interface{} `json:"supportedWriteOperations,omitempty"`
	Logourl interface{} `json:"logoURL,omitempty"`
	Canuseassource interface{} `json:"canUseAsSource,omitempty"`
	Connectorowner interface{} `json:"connectorOwner,omitempty"`
	Canuseasdestination interface{} `json:"canUseAsDestination,omitempty"`
	Connectorarn interface{} `json:"connectorArn,omitempty"`
	Connectordescription interface{} `json:"connectorDescription,omitempty"`
	Isprivatelinkenabled interface{} `json:"isPrivateLinkEnabled,omitempty"`
	Connectormetadata ConnectorConfigurationconnectorMetadata `json:"connectorMetadata,omitempty"`
	Connectorversion interface{} `json:"connectorVersion,omitempty"`
	Supporteddestinationconnectors interface{} `json:"supportedDestinationConnectors,omitempty"`
	Connectorlabel interface{} `json:"connectorLabel,omitempty"`
	Connectormodes interface{} `json:"connectorModes,omitempty"`
	Connectortype interface{} `json:"connectorType,omitempty"`
}

// ConnectorDetail represents the ConnectorDetail schema from the OpenAPI specification
type ConnectorDetail struct {
	Connectorlabel interface{} `json:"connectorLabel,omitempty"`
	Applicationtype interface{} `json:"applicationType,omitempty"`
	Connectordescription interface{} `json:"connectorDescription,omitempty"`
	Connectormodes interface{} `json:"connectorModes,omitempty"`
	Connectorversion interface{} `json:"connectorVersion,omitempty"`
	Registeredby interface{} `json:"registeredBy,omitempty"`
	Connectorname interface{} `json:"connectorName,omitempty"`
	Connectorowner interface{} `json:"connectorOwner,omitempty"`
	Connectortype interface{} `json:"connectorType,omitempty"`
	Supporteddatatransfertypes interface{} `json:"supportedDataTransferTypes,omitempty"`
	Connectorprovisioningtype interface{} `json:"connectorProvisioningType,omitempty"`
	Registeredat interface{} `json:"registeredAt,omitempty"`
}

// ExecutionRecordexecutionResult represents the ExecutionRecordexecutionResult schema from the OpenAPI specification
type ExecutionRecordexecutionResult struct {
	Errorinfo ExecutionResulterrorInfo `json:"errorInfo,omitempty"`
	Recordsprocessed interface{} `json:"recordsProcessed,omitempty"`
	Bytesprocessed interface{} `json:"bytesProcessed,omitempty"`
	Byteswritten interface{} `json:"bytesWritten,omitempty"`
}

// CustomConnectorProfileCredentialscustom represents the CustomConnectorProfileCredentialscustom schema from the OpenAPI specification
type CustomConnectorProfileCredentialscustom struct {
	Credentialsmap interface{} `json:"credentialsMap,omitempty"`
	Customauthenticationtype interface{} `json:"customAuthenticationType"`
}

// ListConnectorsRequest represents the ListConnectorsRequest schema from the OpenAPI specification
type ListConnectorsRequest struct {
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// PardotSourceProperties represents the PardotSourceProperties schema from the OpenAPI specification
type PardotSourceProperties struct {
	Object interface{} `json:"object"`
}

// Range represents the Range schema from the OpenAPI specification
type Range struct {
	Maximum interface{} `json:"maximum,omitempty"`
	Minimum interface{} `json:"minimum,omitempty"`
}

// PrivateConnectionProvisioningState represents the PrivateConnectionProvisioningState schema from the OpenAPI specification
type PrivateConnectionProvisioningState struct {
	Failurecause interface{} `json:"failureCause,omitempty"`
	Failuremessage interface{} `json:"failureMessage,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// ServiceNowConnectorProfileCredentials represents the ServiceNowConnectorProfileCredentials schema from the OpenAPI specification
type ServiceNowConnectorProfileCredentials struct {
	Password interface{} `json:"password"`
	Username interface{} `json:"username"`
}

// SourceConnectorPropertiesDatadog represents the SourceConnectorPropertiesDatadog schema from the OpenAPI specification
type SourceConnectorPropertiesDatadog struct {
	Object interface{} `json:"object"`
}

// CreateFlowrequestsourceFlowConfigsourceConnectorProperties represents the CreateFlowrequestsourceFlowConfigsourceConnectorProperties schema from the OpenAPI specification
type CreateFlowrequestsourceFlowConfigsourceConnectorProperties struct {
	Googleanalytics SourceConnectorPropertiesGoogleAnalytics `json:"GoogleAnalytics,omitempty"`
	Servicenow SourceConnectorPropertiesServiceNow `json:"ServiceNow,omitempty"`
	Dynatrace SourceConnectorPropertiesDynatrace `json:"Dynatrace,omitempty"`
	Pardot SourceConnectorPropertiesPardot `json:"Pardot,omitempty"`
	Sapodata SAPODataSourceProperties `json:"SAPOData,omitempty"` // The properties that are applied when using SAPOData as a flow source.
	Datadog SourceConnectorPropertiesDatadog `json:"Datadog,omitempty"`
	Marketo SourceConnectorPropertiesMarketo `json:"Marketo,omitempty"`
	Amplitude SourceConnectorPropertiesAmplitude `json:"Amplitude,omitempty"`
	Salesforce SourceConnectorPropertiesSalesforce `json:"Salesforce,omitempty"`
	Slack SourceConnectorPropertiesSlack `json:"Slack,omitempty"`
	Customconnector CustomConnectorSourceProperties `json:"CustomConnector,omitempty"` // The properties that are applied when the custom connector is being used as a source.
	Singular SourceConnectorPropertiesSingular `json:"Singular,omitempty"`
	Zendesk SourceConnectorPropertiesZendesk `json:"Zendesk,omitempty"`
	Infornexus SourceConnectorPropertiesInforNexus `json:"InforNexus,omitempty"`
	Veeva SourceConnectorPropertiesVeeva `json:"Veeva,omitempty"`
	S3 SourceConnectorPropertiesS3 `json:"S3,omitempty"`
	Trendmicro SourceConnectorPropertiesTrendmicro `json:"Trendmicro,omitempty"`
}

// ConnectorProfilePropertiesCustomConnector represents the ConnectorProfilePropertiesCustomConnector schema from the OpenAPI specification
type ConnectorProfilePropertiesCustomConnector struct {
	Profileproperties interface{} `json:"profileProperties,omitempty"`
	Oauth2properties OAuth2Properties `json:"oAuth2Properties,omitempty"` // The OAuth 2.0 properties required for OAuth 2.0 authentication.
}

// CustomConnectorProfileCredentialsoauth2 represents the CustomConnectorProfileCredentialsoauth2 schema from the OpenAPI specification
type CustomConnectorProfileCredentialsoauth2 struct {
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientid interface{} `json:"clientId,omitempty"`
	Clientsecret interface{} `json:"clientSecret,omitempty"`
	Oauthrequest ConnectorOAuthRequest `json:"oAuthRequest,omitempty"` // Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
}

// GoogleAnalyticsSourceProperties represents the GoogleAnalyticsSourceProperties schema from the OpenAPI specification
type GoogleAnalyticsSourceProperties struct {
	Object interface{} `json:"object"`
}

// ConnectorConfiguration represents the ConnectorConfiguration schema from the OpenAPI specification
type ConnectorConfiguration struct {
	Authenticationconfig ConnectorConfigurationauthenticationConfig `json:"authenticationConfig,omitempty"`
	Supportedoperators interface{} `json:"supportedOperators,omitempty"`
	Supportedschedulingfrequencies interface{} `json:"supportedSchedulingFrequencies,omitempty"`
	Connectorname interface{} `json:"connectorName,omitempty"`
	Supporteddatatransferapis interface{} `json:"supportedDataTransferApis,omitempty"`
	Connectorprovisioningtype interface{} `json:"connectorProvisioningType,omitempty"`
	Connectorruntimesettings interface{} `json:"connectorRuntimeSettings,omitempty"`
	Registeredby interface{} `json:"registeredBy,omitempty"`
	Supportedwriteoperations interface{} `json:"supportedWriteOperations,omitempty"`
	Logourl interface{} `json:"logoURL,omitempty"`
	Canuseassource interface{} `json:"canUseAsSource,omitempty"`
	Connectorowner interface{} `json:"connectorOwner,omitempty"`
	Canuseasdestination interface{} `json:"canUseAsDestination,omitempty"`
	Connectorarn interface{} `json:"connectorArn,omitempty"`
	Connectordescription interface{} `json:"connectorDescription,omitempty"`
	Isprivatelinkenabled interface{} `json:"isPrivateLinkEnabled,omitempty"`
	Connectormetadata ConnectorConfigurationconnectorMetadata `json:"connectorMetadata,omitempty"`
	Connectorversion interface{} `json:"connectorVersion,omitempty"`
	Supporteddestinationconnectors interface{} `json:"supportedDestinationConnectors,omitempty"`
	Connectorlabel interface{} `json:"connectorLabel,omitempty"`
	Connectormodes interface{} `json:"connectorModes,omitempty"`
	Connectortype interface{} `json:"connectorType,omitempty"`
	Isprivatelinkendpointurlrequired interface{} `json:"isPrivateLinkEndpointUrlRequired,omitempty"`
	Supporteddatatransfertypes interface{} `json:"supportedDataTransferTypes,omitempty"`
	Connectorprovisioningconfig ConnectorConfigurationconnectorProvisioningConfig `json:"connectorProvisioningConfig,omitempty"`
	Registeredat interface{} `json:"registeredAt,omitempty"`
	Supportedapiversions interface{} `json:"supportedApiVersions,omitempty"`
	Supportedtriggertypes interface{} `json:"supportedTriggerTypes,omitempty"`
}

// ResetConnectorMetadataCacherequest represents the ResetConnectorMetadataCacherequest schema from the OpenAPI specification
type ResetConnectorMetadataCacherequest struct {
	Apiversion string `json:"apiVersion,omitempty"` // <p>The API version that you specified in the connector profile that you’re resetting cached metadata for. You must use this parameter only if the connector supports multiple API versions or if the connector type is CustomConnector.</p> <p>To look up how many versions a connector supports, use the DescribeConnectors action. In the response, find the value that Amazon AppFlow returns for the connectorVersion parameter.</p> <p>To look up the connector type, use the DescribeConnectorProfiles action. In the response, find the value that Amazon AppFlow returns for the connectorType parameter.</p> <p>To look up the API version that you specified in a connector profile, use the DescribeConnectorProfiles action.</p>
	Connectorentityname string `json:"connectorEntityName,omitempty"` // <p>Use this parameter if you want to reset cached metadata about the details for an individual entity.</p> <p>If you don't include this parameter in your request, Amazon AppFlow only resets cached metadata about entity names, not entity details.</p>
	Connectorprofilename string `json:"connectorProfileName,omitempty"` // <p>The name of the connector profile that you want to reset cached metadata for.</p> <p>You can omit this parameter if you're resetting the cache for any of the following connectors: Amazon Connect, Amazon EventBridge, Amazon Lookout for Metrics, Amazon S3, or Upsolver. If you're resetting the cache for any other connector, you must include this parameter in your request.</p>
	Connectortype string `json:"connectorType,omitempty"` // <p>The type of connector to reset cached metadata for.</p> <p>You must include this parameter in your request if you're resetting the cache for any of the following connectors: Amazon Connect, Amazon EventBridge, Amazon Lookout for Metrics, Amazon S3, or Upsolver. If you're resetting the cache for any other connector, you can omit this parameter from your request. </p>
	Entitiespath string `json:"entitiesPath,omitempty"` // <p>Use this parameter only if you’re resetting the cached metadata about a nested entity. Only some connectors support nested entities. A nested entity is one that has another entity as a parent. To use this parameter, specify the name of the parent entity.</p> <p>To look up the parent-child relationship of entities, you can send a ListConnectorEntities request that omits the entitiesPath parameter. Amazon AppFlow will return a list of top-level entities. For each one, it indicates whether the entity has nested entities. Then, in a subsequent ListConnectorEntities request, you can specify a parent entity name for the entitiesPath parameter. Amazon AppFlow will return a list of the child entities for that parent.</p>
}

// ConnectorMetadataSlack represents the ConnectorMetadataSlack schema from the OpenAPI specification
type ConnectorMetadataSlack struct {
	Oauthscopes interface{} `json:"oAuthScopes,omitempty"`
}

// ConnectorProfileCredentialsSlack represents the ConnectorProfileCredentialsSlack schema from the OpenAPI specification
type ConnectorProfileCredentialsSlack struct {
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientid interface{} `json:"clientId"`
	Clientsecret interface{} `json:"clientSecret"`
}

// RegisterConnectorRequestconnectorProvisioningConfig represents the RegisterConnectorRequestconnectorProvisioningConfig schema from the OpenAPI specification
type RegisterConnectorRequestconnectorProvisioningConfig struct {
	Lambda RegisterConnectorrequestconnectorProvisioningConfiglambda `json:"lambda,omitempty"`
}

// TaskPropertiesMap represents the TaskPropertiesMap schema from the OpenAPI specification
type TaskPropertiesMap struct {
}

// SingularConnectorProfileProperties represents the SingularConnectorProfileProperties schema from the OpenAPI specification
type SingularConnectorProfileProperties struct {
}

// DescribeConnectorResponse represents the DescribeConnectorResponse schema from the OpenAPI specification
type DescribeConnectorResponse struct {
	Connectorconfiguration DescribeConnectorResponseconnectorConfiguration `json:"connectorConfiguration,omitempty"`
}

// DestinationConnectorPropertiesUpsolver represents the DestinationConnectorPropertiesUpsolver schema from the OpenAPI specification
type DestinationConnectorPropertiesUpsolver struct {
	Bucketname interface{} `json:"bucketName"`
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	S3outputformatconfig UpsolverDestinationPropertiess3OutputFormatConfig `json:"s3OutputFormatConfig"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags"`
}

// StopFlowResponse represents the StopFlowResponse schema from the OpenAPI specification
type StopFlowResponse struct {
	Flowstatus interface{} `json:"flowStatus,omitempty"`
	Flowarn interface{} `json:"flowArn,omitempty"`
}

// BasicAuthCredentials represents the BasicAuthCredentials schema from the OpenAPI specification
type BasicAuthCredentials struct {
	Password interface{} `json:"password"`
	Username interface{} `json:"username"`
}

// ConnectorProfilePropertiesPardot represents the ConnectorProfilePropertiesPardot schema from the OpenAPI specification
type ConnectorProfilePropertiesPardot struct {
	Businessunitid interface{} `json:"businessUnitId,omitempty"`
	Instanceurl interface{} `json:"instanceUrl,omitempty"`
	Issandboxenvironment interface{} `json:"isSandboxEnvironment,omitempty"`
}

// CredentialsMap represents the CredentialsMap schema from the OpenAPI specification
type CredentialsMap struct {
}

// MetadataCatalogConfig represents the MetadataCatalogConfig schema from the OpenAPI specification
type MetadataCatalogConfig struct {
	Gluedatacatalog CreateFlowrequestmetadataCatalogConfigglueDataCatalog `json:"glueDataCatalog,omitempty"`
}

// SAPODataConnectorProfileCredentialsbasicAuthCredentials represents the SAPODataConnectorProfileCredentialsbasicAuthCredentials schema from the OpenAPI specification
type SAPODataConnectorProfileCredentialsbasicAuthCredentials struct {
	Password interface{} `json:"password"`
	Username interface{} `json:"username"`
}

// LookoutMetricsDestinationProperties represents the LookoutMetricsDestinationProperties schema from the OpenAPI specification
type LookoutMetricsDestinationProperties struct {
}

// MetadataCatalogDetailpartitionRegistrationOutput represents the MetadataCatalogDetailpartitionRegistrationOutput schema from the OpenAPI specification
type MetadataCatalogDetailpartitionRegistrationOutput struct {
	Message interface{} `json:"message,omitempty"`
	Result interface{} `json:"result,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// EventBridgeDestinationProperties represents the EventBridgeDestinationProperties schema from the OpenAPI specification
type EventBridgeDestinationProperties struct {
	Errorhandlingconfig ErrorHandlingConfig `json:"errorHandlingConfig,omitempty"` // The settings that determine how Amazon AppFlow handles an error when placing data in the destination. For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. <code>ErrorHandlingConfig</code> is a part of the destination connector details.
	Object interface{} `json:"object"`
}

// SAPODataSourceProperties represents the SAPODataSourceProperties schema from the OpenAPI specification
type SAPODataSourceProperties struct {
	Objectpath interface{} `json:"objectPath,omitempty"`
}

// ConnectorProfileCredentialsRedshift represents the ConnectorProfileCredentialsRedshift schema from the OpenAPI specification
type ConnectorProfileCredentialsRedshift struct {
	Password interface{} `json:"password,omitempty"`
	Username interface{} `json:"username,omitempty"`
}

// ConnectorProfileCredentialsInforNexus represents the ConnectorProfileCredentialsInforNexus schema from the OpenAPI specification
type ConnectorProfileCredentialsInforNexus struct {
	Accesskeyid interface{} `json:"accessKeyId"`
	Datakey interface{} `json:"datakey"`
	Secretaccesskey interface{} `json:"secretAccessKey"`
	Userid interface{} `json:"userId"`
}

// EventBridgeMetadata represents the EventBridgeMetadata schema from the OpenAPI specification
type EventBridgeMetadata struct {
}

// DestinationConnectorPropertiesCustomConnector represents the DestinationConnectorPropertiesCustomConnector schema from the OpenAPI specification
type DestinationConnectorPropertiesCustomConnector struct {
	Customproperties interface{} `json:"customProperties,omitempty"`
	Entityname interface{} `json:"entityName"`
	Errorhandlingconfig CustomConnectorDestinationPropertieserrorHandlingConfig `json:"errorHandlingConfig,omitempty"`
	Idfieldnames interface{} `json:"idFieldNames,omitempty"`
	Writeoperationtype interface{} `json:"writeOperationType,omitempty"`
}

// OAuth2CustomParameter represents the OAuth2CustomParameter schema from the OpenAPI specification
type OAuth2CustomParameter struct {
	Description interface{} `json:"description,omitempty"`
	Isrequired interface{} `json:"isRequired,omitempty"`
	Issensitivefield interface{} `json:"isSensitiveField,omitempty"`
	Key interface{} `json:"key,omitempty"`
	Label interface{} `json:"label,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Connectorsuppliedvalues interface{} `json:"connectorSuppliedValues,omitempty"`
}

// DestinationFieldProperties represents the DestinationFieldProperties schema from the OpenAPI specification
type DestinationFieldProperties struct {
	Isupsertable interface{} `json:"isUpsertable,omitempty"`
	Supportedwriteoperations interface{} `json:"supportedWriteOperations,omitempty"`
	Iscreatable interface{} `json:"isCreatable,omitempty"`
	Isdefaultedoncreate interface{} `json:"isDefaultedOnCreate,omitempty"`
	Isnullable interface{} `json:"isNullable,omitempty"`
	Isupdatable interface{} `json:"isUpdatable,omitempty"`
}

// DescribeFlowExecutionRecordsResponse represents the DescribeFlowExecutionRecordsResponse schema from the OpenAPI specification
type DescribeFlowExecutionRecordsResponse struct {
	Flowexecutions interface{} `json:"flowExecutions,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// CancelFlowExecutionsRequest represents the CancelFlowExecutionsRequest schema from the OpenAPI specification
type CancelFlowExecutionsRequest struct {
	Executionids interface{} `json:"executionIds,omitempty"`
	Flowname interface{} `json:"flowName"`
}

// SourceConnectorProperties represents the SourceConnectorProperties schema from the OpenAPI specification
type SourceConnectorProperties struct {
	Datadog SourceConnectorPropertiesDatadog `json:"Datadog,omitempty"`
	Marketo SourceConnectorPropertiesMarketo `json:"Marketo,omitempty"`
	Amplitude SourceConnectorPropertiesAmplitude `json:"Amplitude,omitempty"`
	Salesforce SourceConnectorPropertiesSalesforce `json:"Salesforce,omitempty"`
	Slack SourceConnectorPropertiesSlack `json:"Slack,omitempty"`
	Customconnector CustomConnectorSourceProperties `json:"CustomConnector,omitempty"` // The properties that are applied when the custom connector is being used as a source.
	Singular SourceConnectorPropertiesSingular `json:"Singular,omitempty"`
	Zendesk SourceConnectorPropertiesZendesk `json:"Zendesk,omitempty"`
	Infornexus SourceConnectorPropertiesInforNexus `json:"InforNexus,omitempty"`
	Veeva SourceConnectorPropertiesVeeva `json:"Veeva,omitempty"`
	S3 SourceConnectorPropertiesS3 `json:"S3,omitempty"`
	Trendmicro SourceConnectorPropertiesTrendmicro `json:"Trendmicro,omitempty"`
	Googleanalytics SourceConnectorPropertiesGoogleAnalytics `json:"GoogleAnalytics,omitempty"`
	Servicenow SourceConnectorPropertiesServiceNow `json:"ServiceNow,omitempty"`
	Dynatrace SourceConnectorPropertiesDynatrace `json:"Dynatrace,omitempty"`
	Pardot SourceConnectorPropertiesPardot `json:"Pardot,omitempty"`
	Sapodata SAPODataSourceProperties `json:"SAPOData,omitempty"` // The properties that are applied when using SAPOData as a flow source.
}

// DestinationConnectorPropertiesRedshift represents the DestinationConnectorPropertiesRedshift schema from the OpenAPI specification
type DestinationConnectorPropertiesRedshift struct {
	Intermediatebucketname interface{} `json:"intermediateBucketName"`
	Object interface{} `json:"object"`
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Errorhandlingconfig RedshiftDestinationPropertieserrorHandlingConfig `json:"errorHandlingConfig,omitempty"`
}

// AggregationConfig represents the AggregationConfig schema from the OpenAPI specification
type AggregationConfig struct {
	Aggregationtype interface{} `json:"aggregationType,omitempty"`
	Targetfilesize interface{} `json:"targetFileSize,omitempty"`
}

// DescribeFlowResponsemetadataCatalogConfig represents the DescribeFlowResponsemetadataCatalogConfig schema from the OpenAPI specification
type DescribeFlowResponsemetadataCatalogConfig struct {
	Gluedatacatalog CreateFlowrequestmetadataCatalogConfigglueDataCatalog `json:"glueDataCatalog,omitempty"`
}

// TrendmicroConnectorProfileProperties represents the TrendmicroConnectorProfileProperties schema from the OpenAPI specification
type TrendmicroConnectorProfileProperties struct {
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// ConnectorEntityFielddestinationProperties represents the ConnectorEntityFielddestinationProperties schema from the OpenAPI specification
type ConnectorEntityFielddestinationProperties struct {
	Isdefaultedoncreate interface{} `json:"isDefaultedOnCreate,omitempty"`
	Isnullable interface{} `json:"isNullable,omitempty"`
	Isupdatable interface{} `json:"isUpdatable,omitempty"`
	Isupsertable interface{} `json:"isUpsertable,omitempty"`
	Supportedwriteoperations interface{} `json:"supportedWriteOperations,omitempty"`
	Iscreatable interface{} `json:"isCreatable,omitempty"`
}

// SAPODataMetadata represents the SAPODataMetadata schema from the OpenAPI specification
type SAPODataMetadata struct {
}

// CustomerProfilesDestinationProperties represents the CustomerProfilesDestinationProperties schema from the OpenAPI specification
type CustomerProfilesDestinationProperties struct {
	Domainname interface{} `json:"domainName"`
	Objecttypename interface{} `json:"objectTypeName,omitempty"`
}

// TaskconnectorOperator represents the TaskconnectorOperator schema from the OpenAPI specification
type TaskconnectorOperator struct {
	Infornexus interface{} `json:"InforNexus,omitempty"`
	Salesforce interface{} `json:"Salesforce,omitempty"`
	Servicenow interface{} `json:"ServiceNow,omitempty"`
	Veeva interface{} `json:"Veeva,omitempty"`
	Amplitude interface{} `json:"Amplitude,omitempty"`
	Datadog interface{} `json:"Datadog,omitempty"`
	Pardot interface{} `json:"Pardot,omitempty"`
	Singular interface{} `json:"Singular,omitempty"`
	Googleanalytics interface{} `json:"GoogleAnalytics,omitempty"`
	Marketo interface{} `json:"Marketo,omitempty"`
	Sapodata interface{} `json:"SAPOData,omitempty"`
	Trendmicro interface{} `json:"Trendmicro,omitempty"`
	Customconnector interface{} `json:"CustomConnector,omitempty"`
	Dynatrace interface{} `json:"Dynatrace,omitempty"`
	S3 interface{} `json:"S3,omitempty"`
	Slack interface{} `json:"Slack,omitempty"`
	Zendesk interface{} `json:"Zendesk,omitempty"`
}

// SAPODataConnectorProfilePropertiesoAuthProperties represents the SAPODataConnectorProfilePropertiesoAuthProperties schema from the OpenAPI specification
type SAPODataConnectorProfilePropertiesoAuthProperties struct {
	Oauthscopes interface{} `json:"oAuthScopes"`
	Tokenurl interface{} `json:"tokenUrl"`
	Authcodeurl interface{} `json:"authCodeUrl"`
}

// SourceConnectorPropertiesS3 represents the SourceConnectorPropertiesS3 schema from the OpenAPI specification
type SourceConnectorPropertiesS3 struct {
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	S3inputformatconfig S3InputFormatConfig `json:"s3InputFormatConfig,omitempty"` // When you use Amazon S3 as the source, the configuration format that you provide the flow input data.
	Bucketname interface{} `json:"bucketName"`
}

// ErrorInfo represents the ErrorInfo schema from the OpenAPI specification
type ErrorInfo struct {
	Putfailurescount interface{} `json:"putFailuresCount,omitempty"`
	Executionmessage interface{} `json:"executionMessage,omitempty"`
}

// GoogleAnalyticsConnectorProfileProperties represents the GoogleAnalyticsConnectorProfileProperties schema from the OpenAPI specification
type GoogleAnalyticsConnectorProfileProperties struct {
}

// CreateConnectorProfilerequestconnectorProfileConfig represents the CreateConnectorProfilerequestconnectorProfileConfig schema from the OpenAPI specification
type CreateConnectorProfilerequestconnectorProfileConfig struct {
	Connectorprofileproperties CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileProperties `json:"connectorProfileProperties,omitempty"`
	Connectorprofilecredentials CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileCredentials `json:"connectorProfileCredentials,omitempty"`
}

// DatadogConnectorProfileProperties represents the DatadogConnectorProfileProperties schema from the OpenAPI specification
type DatadogConnectorProfileProperties struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// UpsolverS3OutputFormatConfig represents the UpsolverS3OutputFormatConfig schema from the OpenAPI specification
type UpsolverS3OutputFormatConfig struct {
	Aggregationconfig AggregationConfig `json:"aggregationConfig,omitempty"` // The aggregation settings that you can use to customize the output format of your flow data.
	Filetype interface{} `json:"fileType,omitempty"`
	Prefixconfig PrefixConfig `json:"prefixConfig"` // Specifies elements that Amazon AppFlow includes in the file and folder names in the flow destination.
}

// UpdateFlowrequest represents the UpdateFlowrequest schema from the OpenAPI specification
type UpdateFlowrequest struct {
	Flowname string `json:"flowName"` // The specified name of the flow. Spaces are not allowed. Use underscores (_) or hyphens (-) only.
	Metadatacatalogconfig CreateFlowrequestmetadataCatalogConfig `json:"metadataCatalogConfig,omitempty"` // Specifies the configuration that Amazon AppFlow uses when it catalogs your data. When Amazon AppFlow catalogs your data, it stores metadata in a data catalog.
	Sourceflowconfig CreateFlowrequestsourceFlowConfig `json:"sourceFlowConfig"` // Contains information about the configuration of the source connector used in the flow.
	Tasks []Task `json:"tasks"` // A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.
	Triggerconfig CreateFlowrequesttriggerConfig `json:"triggerConfig"` // The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
	Clienttoken string `json:"clientToken,omitempty"` // <p>The <code>clientToken</code> parameter is an idempotency token. It ensures that your <code>UpdateFlow</code> request completes only once. You choose the value to pass. For example, if you don't receive a response from your request, you can safely retry the request with the same <code>clientToken</code> parameter value.</p> <p>If you omit a <code>clientToken</code> value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.</p> <p>If you specify input parameters that differ from your first request, an error occurs. If you use a different value for <code>clientToken</code>, Amazon AppFlow considers it a new call to <code>UpdateFlow</code>. The token is active for 8 hours.</p>
	Description string `json:"description,omitempty"` // A description of the flow.
	Destinationflowconfiglist []DestinationFlowConfig `json:"destinationFlowConfigList"` // The configuration that controls how Amazon AppFlow transfers data to the destination connector.
}

// DynatraceSourceProperties represents the DynatraceSourceProperties schema from the OpenAPI specification
type DynatraceSourceProperties struct {
	Object interface{} `json:"object"`
}

// VeevaMetadata represents the VeevaMetadata schema from the OpenAPI specification
type VeevaMetadata struct {
}

// ConnectorProfileCredentialsDynatrace represents the ConnectorProfileCredentialsDynatrace schema from the OpenAPI specification
type ConnectorProfileCredentialsDynatrace struct {
	Apitoken interface{} `json:"apiToken"`
}

// DescribeFlowResponsesourceFlowConfig represents the DescribeFlowResponsesourceFlowConfig schema from the OpenAPI specification
type DescribeFlowResponsesourceFlowConfig struct {
	Incrementalpullconfig CreateFlowrequestsourceFlowConfigincrementalPullConfig `json:"incrementalPullConfig,omitempty"`
	Sourceconnectorproperties CreateFlowrequestsourceFlowConfigsourceConnectorProperties `json:"sourceConnectorProperties"`
	Apiversion interface{} `json:"apiVersion,omitempty"`
	Connectorprofilename interface{} `json:"connectorProfileName,omitempty"`
	Connectortype interface{} `json:"connectorType"`
}

// DynatraceConnectorProfileCredentials represents the DynatraceConnectorProfileCredentials schema from the OpenAPI specification
type DynatraceConnectorProfileCredentials struct {
	Apitoken interface{} `json:"apiToken"`
}

// ListFlowsRequest represents the ListFlowsRequest schema from the OpenAPI specification
type ListFlowsRequest struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
}

// AmplitudeConnectorProfileCredentials represents the AmplitudeConnectorProfileCredentials schema from the OpenAPI specification
type AmplitudeConnectorProfileCredentials struct {
	Secretkey interface{} `json:"secretKey"`
	Apikey interface{} `json:"apiKey"`
}

// HoneycodeConnectorProfileProperties represents the HoneycodeConnectorProfileProperties schema from the OpenAPI specification
type HoneycodeConnectorProfileProperties struct {
}

// ConnectorProfileCredentialsAmplitude represents the ConnectorProfileCredentialsAmplitude schema from the OpenAPI specification
type ConnectorProfileCredentialsAmplitude struct {
	Secretkey interface{} `json:"secretKey"`
	Apikey interface{} `json:"apiKey"`
}

// ListFlowsrequest represents the ListFlowsrequest schema from the OpenAPI specification
type ListFlowsrequest struct {
	Nexttoken string `json:"nextToken,omitempty"` // The pagination token for next page of data.
	Maxresults int `json:"maxResults,omitempty"` // Specifies the maximum number of items that should be returned in the result set.
}

// MetadataCatalogDetailtableRegistrationOutput represents the MetadataCatalogDetailtableRegistrationOutput schema from the OpenAPI specification
type MetadataCatalogDetailtableRegistrationOutput struct {
	Message interface{} `json:"message,omitempty"`
	Result interface{} `json:"result,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// ServiceNowMetadata represents the ServiceNowMetadata schema from the OpenAPI specification
type ServiceNowMetadata struct {
}

// SnowflakeDestinationPropertieserrorHandlingConfig represents the SnowflakeDestinationPropertieserrorHandlingConfig schema from the OpenAPI specification
type SnowflakeDestinationPropertieserrorHandlingConfig struct {
	Bucketname interface{} `json:"bucketName,omitempty"`
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Failonfirstdestinationerror interface{} `json:"failOnFirstDestinationError,omitempty"`
}

// ConnectorEntityFieldsupportedFieldTypeDetails represents the ConnectorEntityFieldsupportedFieldTypeDetails schema from the OpenAPI specification
type ConnectorEntityFieldsupportedFieldTypeDetails struct {
	V1 SupportedFieldTypeDetailsv1 `json:"v1"`
}

// GoogleAnalyticsConnectorProfileCredentialsoAuthRequest represents the GoogleAnalyticsConnectorProfileCredentialsoAuthRequest schema from the OpenAPI specification
type GoogleAnalyticsConnectorProfileCredentialsoAuthRequest struct {
	Authcode interface{} `json:"authCode,omitempty"`
	Redirecturi interface{} `json:"redirectUri,omitempty"`
}

// SourceConnectorPropertiesAmplitude represents the SourceConnectorPropertiesAmplitude schema from the OpenAPI specification
type SourceConnectorPropertiesAmplitude struct {
	Object interface{} `json:"object"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// S3OutputFormatConfig represents the S3OutputFormatConfig schema from the OpenAPI specification
type S3OutputFormatConfig struct {
	Aggregationconfig AggregationConfig `json:"aggregationConfig,omitempty"` // The aggregation settings that you can use to customize the output format of your flow data.
	Filetype interface{} `json:"fileType,omitempty"`
	Prefixconfig S3OutputFormatConfigprefixConfig `json:"prefixConfig,omitempty"`
	Preservesourcedatatyping interface{} `json:"preserveSourceDataTyping,omitempty"`
}

// AuthenticationConfig represents the AuthenticationConfig schema from the OpenAPI specification
type AuthenticationConfig struct {
	Customauthconfigs interface{} `json:"customAuthConfigs,omitempty"`
	Isapikeyauthsupported interface{} `json:"isApiKeyAuthSupported,omitempty"`
	Isbasicauthsupported interface{} `json:"isBasicAuthSupported,omitempty"`
	Iscustomauthsupported interface{} `json:"isCustomAuthSupported,omitempty"`
	Isoauth2supported interface{} `json:"isOAuth2Supported,omitempty"`
	Oauth2defaults AuthenticationConfigoAuth2Defaults `json:"oAuth2Defaults,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// AuthParameter represents the AuthParameter schema from the OpenAPI specification
type AuthParameter struct {
	Description interface{} `json:"description,omitempty"`
	Isrequired interface{} `json:"isRequired,omitempty"`
	Issensitivefield interface{} `json:"isSensitiveField,omitempty"`
	Key interface{} `json:"key,omitempty"`
	Label interface{} `json:"label,omitempty"`
	Connectorsuppliedvalues interface{} `json:"connectorSuppliedValues,omitempty"`
}

// S3InputFormatConfig represents the S3InputFormatConfig schema from the OpenAPI specification
type S3InputFormatConfig struct {
	S3inputfiletype interface{} `json:"s3InputFileType,omitempty"`
}

// PardotMetadata represents the PardotMetadata schema from the OpenAPI specification
type PardotMetadata struct {
}

// DatadogConnectorProfileCredentials represents the DatadogConnectorProfileCredentials schema from the OpenAPI specification
type DatadogConnectorProfileCredentials struct {
	Apikey interface{} `json:"apiKey"`
	Applicationkey interface{} `json:"applicationKey"`
}

// DestinationConnectorPropertiesHoneycode represents the DestinationConnectorPropertiesHoneycode schema from the OpenAPI specification
type DestinationConnectorPropertiesHoneycode struct {
	Errorhandlingconfig ErrorHandlingConfig `json:"errorHandlingConfig,omitempty"` // The settings that determine how Amazon AppFlow handles an error when placing data in the destination. For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. <code>ErrorHandlingConfig</code> is a part of the destination connector details.
	Object interface{} `json:"object"`
}

// InforNexusSourceProperties represents the InforNexusSourceProperties schema from the OpenAPI specification
type InforNexusSourceProperties struct {
	Object interface{} `json:"object"`
}

// DescribeFlowResponsetriggerConfig represents the DescribeFlowResponsetriggerConfig schema from the OpenAPI specification
type DescribeFlowResponsetriggerConfig struct {
	Triggerproperties CreateFlowrequesttriggerConfigtriggerProperties `json:"triggerProperties,omitempty"`
	Triggertype interface{} `json:"triggerType"`
}

// AmplitudeMetadata represents the AmplitudeMetadata schema from the OpenAPI specification
type AmplitudeMetadata struct {
}

// ConnectorProvisioningConfig represents the ConnectorProvisioningConfig schema from the OpenAPI specification
type ConnectorProvisioningConfig struct {
	Lambda RegisterConnectorrequestconnectorProvisioningConfiglambda `json:"lambda,omitempty"`
}

// S3Metadata represents the S3Metadata schema from the OpenAPI specification
type S3Metadata struct {
}

// SourceFieldProperties represents the SourceFieldProperties schema from the OpenAPI specification
type SourceFieldProperties struct {
	Isqueryable interface{} `json:"isQueryable,omitempty"`
	Isretrievable interface{} `json:"isRetrievable,omitempty"`
	Istimestampfieldforincrementalqueries interface{} `json:"isTimestampFieldForIncrementalQueries,omitempty"`
}

// DeleteFlowRequest represents the DeleteFlowRequest schema from the OpenAPI specification
type DeleteFlowRequest struct {
	Flowname interface{} `json:"flowName"`
	Forcedelete interface{} `json:"forceDelete,omitempty"`
}

// CreateConnectorProfileRequest represents the CreateConnectorProfileRequest schema from the OpenAPI specification
type CreateConnectorProfileRequest struct {
	Connectorlabel interface{} `json:"connectorLabel,omitempty"`
	Connectorprofileconfig CreateConnectorProfileRequestconnectorProfileConfig `json:"connectorProfileConfig"`
	Connectorprofilename interface{} `json:"connectorProfileName"`
	Connectortype interface{} `json:"connectorType"`
	Kmsarn interface{} `json:"kmsArn,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Connectionmode interface{} `json:"connectionMode"`
}

// ConnectorProfileCredentialsHoneycode represents the ConnectorProfileCredentialsHoneycode schema from the OpenAPI specification
type ConnectorProfileCredentialsHoneycode struct {
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Oauthrequest ConnectorOAuthRequest `json:"oAuthRequest,omitempty"` // Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
}

// ConnectorConfigurationconnectorMetadata represents the ConnectorConfigurationconnectorMetadata schema from the OpenAPI specification
type ConnectorConfigurationconnectorMetadata struct {
	Redshift interface{} `json:"Redshift,omitempty"`
	Snowflake ConnectorMetadataSnowflake `json:"Snowflake,omitempty"`
	Googleanalytics ConnectorMetadataGoogleAnalytics `json:"GoogleAnalytics,omitempty"`
	Slack ConnectorMetadataSlack `json:"Slack,omitempty"`
	Marketo interface{} `json:"Marketo,omitempty"`
	Singular interface{} `json:"Singular,omitempty"`
	Datadog interface{} `json:"Datadog,omitempty"`
	S3 interface{} `json:"S3,omitempty"`
	Sapodata map[string]interface{} `json:"SAPOData,omitempty"` // The connector metadata specific to SAPOData.
	Salesforce ConnectorMetadataSalesforce `json:"Salesforce,omitempty"`
	Infornexus interface{} `json:"InforNexus,omitempty"`
	Trendmicro interface{} `json:"Trendmicro,omitempty"`
	Amplitude interface{} `json:"Amplitude,omitempty"`
	Dynatrace interface{} `json:"Dynatrace,omitempty"`
	Honeycode ConnectorMetadataHoneycode `json:"Honeycode,omitempty"`
	Upsolver interface{} `json:"Upsolver,omitempty"`
	Servicenow interface{} `json:"ServiceNow,omitempty"`
	Zendesk ConnectorMetadataZendesk `json:"Zendesk,omitempty"`
	Veeva interface{} `json:"Veeva,omitempty"`
	Customerprofiles interface{} `json:"CustomerProfiles,omitempty"`
	Eventbridge interface{} `json:"EventBridge,omitempty"`
	Pardot interface{} `json:"Pardot,omitempty"`
}

// CustomConnectorProfileCredentialsapiKey represents the CustomConnectorProfileCredentialsapiKey schema from the OpenAPI specification
type CustomConnectorProfileCredentialsapiKey struct {
	Apikey interface{} `json:"apiKey"`
	Apisecretkey interface{} `json:"apiSecretKey,omitempty"`
}

// ErrorHandlingConfig represents the ErrorHandlingConfig schema from the OpenAPI specification
type ErrorHandlingConfig struct {
	Bucketname interface{} `json:"bucketName,omitempty"`
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Failonfirstdestinationerror interface{} `json:"failOnFirstDestinationError,omitempty"`
}

// AmplitudeConnectorProfileProperties represents the AmplitudeConnectorProfileProperties schema from the OpenAPI specification
type AmplitudeConnectorProfileProperties struct {
}

// CustomerProfilesMetadata represents the CustomerProfilesMetadata schema from the OpenAPI specification
type CustomerProfilesMetadata struct {
}

// ConnectorProfilePropertiesInforNexus represents the ConnectorProfilePropertiesInforNexus schema from the OpenAPI specification
type ConnectorProfilePropertiesInforNexus struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// UpdateFlowResponse represents the UpdateFlowResponse schema from the OpenAPI specification
type UpdateFlowResponse struct {
	Flowstatus interface{} `json:"flowStatus,omitempty"`
}

// GoogleAnalyticsMetadata represents the GoogleAnalyticsMetadata schema from the OpenAPI specification
type GoogleAnalyticsMetadata struct {
	Oauthscopes interface{} `json:"oAuthScopes,omitempty"`
}

// UpdateConnectorRegistrationRequest represents the UpdateConnectorRegistrationRequest schema from the OpenAPI specification
type UpdateConnectorRegistrationRequest struct {
	Connectorlabel interface{} `json:"connectorLabel"`
	Connectorprovisioningconfig ConnectorProvisioningConfig `json:"connectorProvisioningConfig,omitempty"` // Contains information about the configuration of the connector being registered.
	Description interface{} `json:"description,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// HoneycodeDestinationProperties represents the HoneycodeDestinationProperties schema from the OpenAPI specification
type HoneycodeDestinationProperties struct {
	Errorhandlingconfig ErrorHandlingConfig `json:"errorHandlingConfig,omitempty"` // The settings that determine how Amazon AppFlow handles an error when placing data in the destination. For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. <code>ErrorHandlingConfig</code> is a part of the destination connector details.
	Object interface{} `json:"object"`
}

// ListConnectorsrequest represents the ListConnectorsrequest schema from the OpenAPI specification
type ListConnectorsrequest struct {
	Maxresults int `json:"maxResults,omitempty"` // Specifies the maximum number of items that should be returned in the result set. The default for <code>maxResults</code> is 20 (for all paginated API operations).
	Nexttoken string `json:"nextToken,omitempty"` // The pagination token for the next page of data.
}

// SalesforceDestinationPropertieserrorHandlingConfig represents the SalesforceDestinationPropertieserrorHandlingConfig schema from the OpenAPI specification
type SalesforceDestinationPropertieserrorHandlingConfig struct {
	Bucketname interface{} `json:"bucketName,omitempty"`
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Failonfirstdestinationerror interface{} `json:"failOnFirstDestinationError,omitempty"`
}

// SalesforceSourceProperties represents the SalesforceSourceProperties schema from the OpenAPI specification
type SalesforceSourceProperties struct {
	Datatransferapi interface{} `json:"dataTransferApi,omitempty"`
	Enabledynamicfieldupdate interface{} `json:"enableDynamicFieldUpdate,omitempty"`
	Includedeletedrecords interface{} `json:"includeDeletedRecords,omitempty"`
	Object interface{} `json:"object"`
}

// DescribeConnectorsResponse represents the DescribeConnectorsResponse schema from the OpenAPI specification
type DescribeConnectorsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Connectorconfigurations interface{} `json:"connectorConfigurations,omitempty"`
	Connectors interface{} `json:"connectors,omitempty"`
}

// ExecutionResulterrorInfo represents the ExecutionResulterrorInfo schema from the OpenAPI specification
type ExecutionResulterrorInfo struct {
	Putfailurescount interface{} `json:"putFailuresCount,omitempty"`
	Executionmessage interface{} `json:"executionMessage,omitempty"`
}

// ConnectorMetadataSnowflake represents the ConnectorMetadataSnowflake schema from the OpenAPI specification
type ConnectorMetadataSnowflake struct {
	Supportedregions interface{} `json:"supportedRegions,omitempty"`
}

// SalesforceConnectorProfileCredentials represents the SalesforceConnectorProfileCredentials schema from the OpenAPI specification
type SalesforceConnectorProfileCredentials struct {
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientcredentialsarn interface{} `json:"clientCredentialsArn,omitempty"`
	Jwttoken interface{} `json:"jwtToken,omitempty"`
	Oauth2granttype interface{} `json:"oAuth2GrantType,omitempty"`
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
}

// UpdateConnectorProfileRequestconnectorProfileConfig represents the UpdateConnectorProfileRequestconnectorProfileConfig schema from the OpenAPI specification
type UpdateConnectorProfileRequestconnectorProfileConfig struct {
	Connectorprofilecredentials CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileCredentials `json:"connectorProfileCredentials,omitempty"`
	Connectorprofileproperties CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileProperties `json:"connectorProfileProperties"`
}

// DescribeConnectorEntityrequest represents the DescribeConnectorEntityrequest schema from the OpenAPI specification
type DescribeConnectorEntityrequest struct {
	Connectortype string `json:"connectorType,omitempty"` // The type of connector application, such as Salesforce, Amplitude, and so on.
	Apiversion string `json:"apiVersion,omitempty"` // The version of the API that's used by the connector.
	Connectorentityname string `json:"connectorEntityName"` // The entity name for that connector.
	Connectorprofilename string `json:"connectorProfileName,omitempty"` // The name of the connector profile. The name is unique for each <code>ConnectorProfile</code> in the Amazon Web Services account.
}

// MarketoConnectorProfileCredentials represents the MarketoConnectorProfileCredentials schema from the OpenAPI specification
type MarketoConnectorProfileCredentials struct {
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientid interface{} `json:"clientId"`
	Clientsecret interface{} `json:"clientSecret"`
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
}

// ConnectorProfileCredentialsDatadog represents the ConnectorProfileCredentialsDatadog schema from the OpenAPI specification
type ConnectorProfileCredentialsDatadog struct {
	Applicationkey interface{} `json:"applicationKey"`
	Apikey interface{} `json:"apiKey"`
}

// TokenUrlCustomProperties represents the TokenUrlCustomProperties schema from the OpenAPI specification
type TokenUrlCustomProperties struct {
}

// CreateConnectorProfilerequest represents the CreateConnectorProfilerequest schema from the OpenAPI specification
type CreateConnectorProfilerequest struct {
	Connectorlabel string `json:"connectorLabel,omitempty"` // The label of the connector. The label is unique for each <code>ConnectorRegistration</code> in your Amazon Web Services account. Only needed if calling for CUSTOMCONNECTOR connector type/.
	Connectorprofileconfig CreateConnectorProfilerequestconnectorProfileConfig `json:"connectorProfileConfig"` // Defines the connector-specific configuration and credentials for the connector profile.
	Connectorprofilename string `json:"connectorProfileName"` // The name of the connector profile. The name is unique for each <code>ConnectorProfile</code> in your Amazon Web Services account.
	Connectortype string `json:"connectorType"` // The type of connector, such as Salesforce, Amplitude, and so on.
	Kmsarn string `json:"kmsArn,omitempty"` // The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption. This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	Clienttoken string `json:"clientToken,omitempty"` // <p>The <code>clientToken</code> parameter is an idempotency token. It ensures that your <code>CreateConnectorProfile</code> request completes only once. You choose the value to pass. For example, if you don't receive a response from your request, you can safely retry the request with the same <code>clientToken</code> parameter value.</p> <p>If you omit a <code>clientToken</code> value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.</p> <p>If you specify input parameters that differ from your first request, an error occurs. If you use a different value for <code>clientToken</code>, Amazon AppFlow considers it a new call to <code>CreateConnectorProfile</code>. The token is active for 8 hours.</p>
	Connectionmode string `json:"connectionMode"` // Indicates the connection mode and specifies whether it is public or private. Private flows use Amazon Web Services PrivateLink to route data over Amazon Web Services infrastructure without exposing it to the public internet.
}

// StartFlowrequest represents the StartFlowrequest schema from the OpenAPI specification
type StartFlowrequest struct {
	Clienttoken string `json:"clientToken,omitempty"` // <p>The <code>clientToken</code> parameter is an idempotency token. It ensures that your <code>StartFlow</code> request completes only once. You choose the value to pass. For example, if you don't receive a response from your request, you can safely retry the request with the same <code>clientToken</code> parameter value.</p> <p>If you omit a <code>clientToken</code> value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.</p> <p>If you specify input parameters that differ from your first request, an error occurs for flows that run on a schedule or based on an event. However, the error doesn't occur for flows that run on demand. You set the conditions that initiate your flow for the <code>triggerConfig</code> parameter.</p> <p>If you use a different value for <code>clientToken</code>, Amazon AppFlow considers it a new call to <code>StartFlow</code>. The token is active for 8 hours.</p>
	Flowname string `json:"flowName"` // The specified name of the flow. Spaces are not allowed. Use underscores (_) or hyphens (-) only.
}

// DestinationConnectorPropertiesSAPOData represents the DestinationConnectorPropertiesSAPOData schema from the OpenAPI specification
type DestinationConnectorPropertiesSAPOData struct {
	Successresponsehandlingconfig SAPODataDestinationPropertiessuccessResponseHandlingConfig `json:"successResponseHandlingConfig,omitempty"`
	Writeoperationtype string `json:"writeOperationType,omitempty"` // The possible write operations in the destination connector. When this value is not provided, this defaults to the <code>INSERT</code> operation.
	Errorhandlingconfig ErrorHandlingConfig `json:"errorHandlingConfig,omitempty"` // The settings that determine how Amazon AppFlow handles an error when placing data in the destination. For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. <code>ErrorHandlingConfig</code> is a part of the destination connector details.
	Idfieldnames []string `json:"idFieldNames,omitempty"` // A list of field names that can be used as an ID field when performing a write operation.
	Objectpath interface{} `json:"objectPath"`
}

// RedshiftDestinationPropertieserrorHandlingConfig represents the RedshiftDestinationPropertieserrorHandlingConfig schema from the OpenAPI specification
type RedshiftDestinationPropertieserrorHandlingConfig struct {
	Bucketname interface{} `json:"bucketName,omitempty"`
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Failonfirstdestinationerror interface{} `json:"failOnFirstDestinationError,omitempty"`
}

// GlueDataCatalogConfig represents the GlueDataCatalogConfig schema from the OpenAPI specification
type GlueDataCatalogConfig struct {
	Databasename interface{} `json:"databaseName"`
	Rolearn interface{} `json:"roleArn"`
	Tableprefix interface{} `json:"tablePrefix"`
}

// ZendeskDestinationProperties represents the ZendeskDestinationProperties schema from the OpenAPI specification
type ZendeskDestinationProperties struct {
	Errorhandlingconfig ErrorHandlingConfig `json:"errorHandlingConfig,omitempty"` // The settings that determine how Amazon AppFlow handles an error when placing data in the destination. For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. <code>ErrorHandlingConfig</code> is a part of the destination connector details.
	Idfieldnames []string `json:"idFieldNames,omitempty"` // A list of field names that can be used as an ID field when performing a write operation.
	Object interface{} `json:"object"`
	Writeoperationtype string `json:"writeOperationType,omitempty"` // The possible write operations in the destination connector. When this value is not provided, this defaults to the <code>INSERT</code> operation.
}

// DeleteConnectorProfileRequest represents the DeleteConnectorProfileRequest schema from the OpenAPI specification
type DeleteConnectorProfileRequest struct {
	Connectorprofilename interface{} `json:"connectorProfileName"`
	Forcedelete interface{} `json:"forceDelete,omitempty"`
}

// RegisterConnectorrequestconnectorProvisioningConfiglambda represents the RegisterConnectorrequestconnectorProvisioningConfiglambda schema from the OpenAPI specification
type RegisterConnectorrequestconnectorProvisioningConfiglambda struct {
	Lambdaarn interface{} `json:"lambdaArn"`
}

// CreateConnectorProfileResponse represents the CreateConnectorProfileResponse schema from the OpenAPI specification
type CreateConnectorProfileResponse struct {
	Connectorprofilearn interface{} `json:"connectorProfileArn,omitempty"`
}

// UpsolverDestinationPropertiess3OutputFormatConfig represents the UpsolverDestinationPropertiess3OutputFormatConfig schema from the OpenAPI specification
type UpsolverDestinationPropertiess3OutputFormatConfig struct {
	Aggregationconfig AggregationConfig `json:"aggregationConfig,omitempty"` // The aggregation settings that you can use to customize the output format of your flow data.
	Filetype interface{} `json:"fileType,omitempty"`
	Prefixconfig PrefixConfig `json:"prefixConfig"` // Specifies elements that Amazon AppFlow includes in the file and folder names in the flow destination.
}

// ConnectorConfigurationauthenticationConfig represents the ConnectorConfigurationauthenticationConfig schema from the OpenAPI specification
type ConnectorConfigurationauthenticationConfig struct {
	Isapikeyauthsupported interface{} `json:"isApiKeyAuthSupported,omitempty"`
	Isbasicauthsupported interface{} `json:"isBasicAuthSupported,omitempty"`
	Iscustomauthsupported interface{} `json:"isCustomAuthSupported,omitempty"`
	Isoauth2supported interface{} `json:"isOAuth2Supported,omitempty"`
	Oauth2defaults AuthenticationConfigoAuth2Defaults `json:"oAuth2Defaults,omitempty"`
	Customauthconfigs interface{} `json:"customAuthConfigs,omitempty"`
}

// DescribeFlowResponselastRunExecutionDetails represents the DescribeFlowResponselastRunExecutionDetails schema from the OpenAPI specification
type DescribeFlowResponselastRunExecutionDetails struct {
	Mostrecentexecutiontime interface{} `json:"mostRecentExecutionTime,omitempty"`
	Mostrecentexecutionmessage interface{} `json:"mostRecentExecutionMessage,omitempty"`
	Mostrecentexecutionstatus interface{} `json:"mostRecentExecutionStatus,omitempty"`
}

// ZendeskSourceProperties represents the ZendeskSourceProperties schema from the OpenAPI specification
type ZendeskSourceProperties struct {
	Object interface{} `json:"object"`
}

// ConnectorEntityMap represents the ConnectorEntityMap schema from the OpenAPI specification
type ConnectorEntityMap struct {
}

// OAuthProperties represents the OAuthProperties schema from the OpenAPI specification
type OAuthProperties struct {
	Oauthscopes interface{} `json:"oAuthScopes"`
	Tokenurl interface{} `json:"tokenUrl"`
	Authcodeurl interface{} `json:"authCodeUrl"`
}

// UpsolverMetadata represents the UpsolverMetadata schema from the OpenAPI specification
type UpsolverMetadata struct {
}

// SnowflakeDestinationProperties represents the SnowflakeDestinationProperties schema from the OpenAPI specification
type SnowflakeDestinationProperties struct {
	Object interface{} `json:"object"`
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Errorhandlingconfig SnowflakeDestinationPropertieserrorHandlingConfig `json:"errorHandlingConfig,omitempty"`
	Intermediatebucketname interface{} `json:"intermediateBucketName"`
}

// ServiceNowConnectorProfileProperties represents the ServiceNowConnectorProfileProperties schema from the OpenAPI specification
type ServiceNowConnectorProfileProperties struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// RedshiftConnectorProfileProperties represents the RedshiftConnectorProfileProperties schema from the OpenAPI specification
type RedshiftConnectorProfileProperties struct {
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Clusteridentifier interface{} `json:"clusterIdentifier,omitempty"`
	Databasename interface{} `json:"databaseName,omitempty"`
	Dataapirolearn interface{} `json:"dataApiRoleArn,omitempty"`
	Databaseurl interface{} `json:"databaseUrl,omitempty"`
	Isredshiftserverless interface{} `json:"isRedshiftServerless,omitempty"`
	Workgroupname interface{} `json:"workgroupName,omitempty"`
	Bucketname interface{} `json:"bucketName"`
	Rolearn interface{} `json:"roleArn"`
}

// ConnectorProfileCredentialsGoogleAnalytics represents the ConnectorProfileCredentialsGoogleAnalytics schema from the OpenAPI specification
type ConnectorProfileCredentialsGoogleAnalytics struct {
	Clientsecret interface{} `json:"clientSecret"`
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientid interface{} `json:"clientId"`
}

// SAPODataDestinationProperties represents the SAPODataDestinationProperties schema from the OpenAPI specification
type SAPODataDestinationProperties struct {
	Errorhandlingconfig ErrorHandlingConfig `json:"errorHandlingConfig,omitempty"` // The settings that determine how Amazon AppFlow handles an error when placing data in the destination. For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. <code>ErrorHandlingConfig</code> is a part of the destination connector details.
	Idfieldnames []string `json:"idFieldNames,omitempty"` // A list of field names that can be used as an ID field when performing a write operation.
	Objectpath interface{} `json:"objectPath"`
	Successresponsehandlingconfig SAPODataDestinationPropertiessuccessResponseHandlingConfig `json:"successResponseHandlingConfig,omitempty"`
	Writeoperationtype string `json:"writeOperationType,omitempty"` // The possible write operations in the destination connector. When this value is not provided, this defaults to the <code>INSERT</code> operation.
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// TriggerPropertiesScheduled represents the TriggerPropertiesScheduled schema from the OpenAPI specification
type TriggerPropertiesScheduled struct {
	Scheduleoffset interface{} `json:"scheduleOffset,omitempty"`
	Schedulestarttime interface{} `json:"scheduleStartTime,omitempty"`
	Timezone interface{} `json:"timezone,omitempty"`
	Datapullmode interface{} `json:"dataPullMode,omitempty"`
	Firstexecutionfrom interface{} `json:"firstExecutionFrom,omitempty"`
	Flowerrordeactivationthreshold interface{} `json:"flowErrorDeactivationThreshold,omitempty"`
	Scheduleendtime interface{} `json:"scheduleEndTime,omitempty"`
	Scheduleexpression interface{} `json:"scheduleExpression"`
}

// OAuth2Properties represents the OAuth2Properties schema from the OpenAPI specification
type OAuth2Properties struct {
	Tokenurlcustomproperties interface{} `json:"tokenUrlCustomProperties,omitempty"`
	Oauth2granttype interface{} `json:"oAuth2GrantType"`
	Tokenurl interface{} `json:"tokenUrl"`
}

// ExecutionRecord represents the ExecutionRecord schema from the OpenAPI specification
type ExecutionRecord struct {
	Startedat interface{} `json:"startedAt,omitempty"`
	Datapullendtime interface{} `json:"dataPullEndTime,omitempty"`
	Datapullstarttime interface{} `json:"dataPullStartTime,omitempty"`
	Executionid interface{} `json:"executionId,omitempty"`
	Executionresult ExecutionRecordexecutionResult `json:"executionResult,omitempty"`
	Executionstatus interface{} `json:"executionStatus,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Metadatacatalogdetails interface{} `json:"metadataCatalogDetails,omitempty"`
}

// ConnectorEntityField represents the ConnectorEntityField schema from the OpenAPI specification
type ConnectorEntityField struct {
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Label interface{} `json:"label,omitempty"`
	Destinationproperties ConnectorEntityFielddestinationProperties `json:"destinationProperties,omitempty"`
	Identifier interface{} `json:"identifier"`
	Parentidentifier interface{} `json:"parentIdentifier,omitempty"`
	Sourceproperties ConnectorEntityFieldsourceProperties `json:"sourceProperties,omitempty"`
	Supportedfieldtypedetails ConnectorEntityFieldsupportedFieldTypeDetails `json:"supportedFieldTypeDetails,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Isdeprecated interface{} `json:"isDeprecated,omitempty"`
	Isprimarykey interface{} `json:"isPrimaryKey,omitempty"`
	Customproperties interface{} `json:"customProperties,omitempty"`
}

// CustomConnectorSourcePropertiesdataTransferApi represents the CustomConnectorSourcePropertiesdataTransferApi schema from the OpenAPI specification
type CustomConnectorSourcePropertiesdataTransferApi struct {
	Name interface{} `json:"Name,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// ZendeskConnectorProfileProperties represents the ZendeskConnectorProfileProperties schema from the OpenAPI specification
type ZendeskConnectorProfileProperties struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// ConnectorProfilePropertiesMarketo represents the ConnectorProfilePropertiesMarketo schema from the OpenAPI specification
type ConnectorProfilePropertiesMarketo struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// PrefixConfig represents the PrefixConfig schema from the OpenAPI specification
type PrefixConfig struct {
	Pathprefixhierarchy interface{} `json:"pathPrefixHierarchy,omitempty"`
	Prefixformat interface{} `json:"prefixFormat,omitempty"`
	Prefixtype interface{} `json:"prefixType,omitempty"`
}

// ConnectorProfileCredentialsSnowflake represents the ConnectorProfileCredentialsSnowflake schema from the OpenAPI specification
type ConnectorProfileCredentialsSnowflake struct {
	Username interface{} `json:"username"`
	Password interface{} `json:"password"`
}

// SupportedFieldTypeDetails represents the SupportedFieldTypeDetails schema from the OpenAPI specification
type SupportedFieldTypeDetails struct {
	V1 SupportedFieldTypeDetailsv1 `json:"v1"`
}

// RedshiftConnectorProfileCredentials represents the RedshiftConnectorProfileCredentials schema from the OpenAPI specification
type RedshiftConnectorProfileCredentials struct {
	Password interface{} `json:"password,omitempty"`
	Username interface{} `json:"username,omitempty"`
}

// OAuthCredentials represents the OAuthCredentials schema from the OpenAPI specification
type OAuthCredentials struct {
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientid interface{} `json:"clientId"`
	Clientsecret interface{} `json:"clientSecret"`
}

// UnregisterConnectorRequest represents the UnregisterConnectorRequest schema from the OpenAPI specification
type UnregisterConnectorRequest struct {
	Connectorlabel interface{} `json:"connectorLabel"`
	Forcedelete interface{} `json:"forceDelete,omitempty"`
}

// FieldTypeDetailsfieldValueRange represents the FieldTypeDetailsfieldValueRange schema from the OpenAPI specification
type FieldTypeDetailsfieldValueRange struct {
	Maximum interface{} `json:"maximum,omitempty"`
	Minimum interface{} `json:"minimum,omitempty"`
}

// ConnectorRuntimeSetting represents the ConnectorRuntimeSetting schema from the OpenAPI specification
type ConnectorRuntimeSetting struct {
	Isrequired interface{} `json:"isRequired,omitempty"`
	Key interface{} `json:"key,omitempty"`
	Label interface{} `json:"label,omitempty"`
	Scope interface{} `json:"scope,omitempty"`
	Connectorsuppliedvalueoptions interface{} `json:"connectorSuppliedValueOptions,omitempty"`
	Datatype interface{} `json:"dataType,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

// RegistrationOutput represents the RegistrationOutput schema from the OpenAPI specification
type RegistrationOutput struct {
	Status interface{} `json:"status,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

// SnowflakeConnectorProfileProperties represents the SnowflakeConnectorProfileProperties schema from the OpenAPI specification
type SnowflakeConnectorProfileProperties struct {
	Privatelinkservicename interface{} `json:"privateLinkServiceName,omitempty"`
	Region interface{} `json:"region,omitempty"`
	Stage interface{} `json:"stage"`
	Warehouse interface{} `json:"warehouse"`
	Accountname interface{} `json:"accountName,omitempty"`
	Bucketname interface{} `json:"bucketName"`
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
}

// ScheduledTriggerProperties represents the ScheduledTriggerProperties schema from the OpenAPI specification
type ScheduledTriggerProperties struct {
	Scheduleendtime interface{} `json:"scheduleEndTime,omitempty"`
	Scheduleexpression interface{} `json:"scheduleExpression"`
	Scheduleoffset interface{} `json:"scheduleOffset,omitempty"`
	Schedulestarttime interface{} `json:"scheduleStartTime,omitempty"`
	Timezone interface{} `json:"timezone,omitempty"`
	Datapullmode interface{} `json:"dataPullMode,omitempty"`
	Firstexecutionfrom interface{} `json:"firstExecutionFrom,omitempty"`
	Flowerrordeactivationthreshold interface{} `json:"flowErrorDeactivationThreshold,omitempty"`
}

// CreateFlowResponse represents the CreateFlowResponse schema from the OpenAPI specification
type CreateFlowResponse struct {
	Flowarn interface{} `json:"flowArn,omitempty"`
	Flowstatus interface{} `json:"flowStatus,omitempty"`
}

// UpdateFlowRequest represents the UpdateFlowRequest schema from the OpenAPI specification
type UpdateFlowRequest struct {
	Triggerconfig DescribeFlowResponsetriggerConfig `json:"triggerConfig"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Destinationflowconfiglist interface{} `json:"destinationFlowConfigList"`
	Flowname interface{} `json:"flowName"`
	Metadatacatalogconfig DescribeFlowResponsemetadataCatalogConfig `json:"metadataCatalogConfig,omitempty"`
	Sourceflowconfig SourceFlowConfig `json:"sourceFlowConfig"` // Contains information about the configuration of the source connector used in the flow.
	Tasks interface{} `json:"tasks"`
}

// DeleteFlowrequest represents the DeleteFlowrequest schema from the OpenAPI specification
type DeleteFlowrequest struct {
	Forcedelete bool `json:"forceDelete,omitempty"` // Indicates whether Amazon AppFlow should delete the flow, even if it is currently in use.
	Flowname string `json:"flowName"` // The specified name of the flow. Spaces are not allowed. Use underscores (_) or hyphens (-) only.
}

// GoogleAnalyticsConnectorProfileCredentials represents the GoogleAnalyticsConnectorProfileCredentials schema from the OpenAPI specification
type GoogleAnalyticsConnectorProfileCredentials struct {
	Clientid interface{} `json:"clientId"`
	Clientsecret interface{} `json:"clientSecret"`
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
}

// DescribeConnectorrequest represents the DescribeConnectorrequest schema from the OpenAPI specification
type DescribeConnectorrequest struct {
	Connectorlabel string `json:"connectorLabel,omitempty"` // The label of the connector. The label is unique for each <code>ConnectorRegistration</code> in your Amazon Web Services account. Only needed if calling for CUSTOMCONNECTOR connector type/.
	Connectortype string `json:"connectorType"` // The connector type, such as CUSTOMCONNECTOR, Saleforce, Marketo. Please choose CUSTOMCONNECTOR for Lambda based custom connectors.
}

// ConnectorProfileConfig represents the ConnectorProfileConfig schema from the OpenAPI specification
type ConnectorProfileConfig struct {
	Connectorprofilecredentials CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileCredentials `json:"connectorProfileCredentials,omitempty"`
	Connectorprofileproperties CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileProperties `json:"connectorProfileProperties"`
}

// PardotConnectorProfileProperties represents the PardotConnectorProfileProperties schema from the OpenAPI specification
type PardotConnectorProfileProperties struct {
	Businessunitid interface{} `json:"businessUnitId,omitempty"`
	Instanceurl interface{} `json:"instanceUrl,omitempty"`
	Issandboxenvironment interface{} `json:"isSandboxEnvironment,omitempty"`
}

// ConnectorMetadataSalesforce represents the ConnectorMetadataSalesforce schema from the OpenAPI specification
type ConnectorMetadataSalesforce struct {
	Datatransferapis interface{} `json:"dataTransferApis,omitempty"`
	Oauthscopes interface{} `json:"oAuthScopes,omitempty"`
	Oauth2granttypessupported interface{} `json:"oauth2GrantTypesSupported,omitempty"`
}

// CustomConnectorDestinationProperties represents the CustomConnectorDestinationProperties schema from the OpenAPI specification
type CustomConnectorDestinationProperties struct {
	Idfieldnames interface{} `json:"idFieldNames,omitempty"`
	Writeoperationtype interface{} `json:"writeOperationType,omitempty"`
	Customproperties interface{} `json:"customProperties,omitempty"`
	Entityname interface{} `json:"entityName"`
	Errorhandlingconfig CustomConnectorDestinationPropertieserrorHandlingConfig `json:"errorHandlingConfig,omitempty"`
}

// SourceConnectorPropertiesDynatrace represents the SourceConnectorPropertiesDynatrace schema from the OpenAPI specification
type SourceConnectorPropertiesDynatrace struct {
	Object interface{} `json:"object"`
}

// CreateFlowrequestmetadataCatalogConfigglueDataCatalog represents the CreateFlowrequestmetadataCatalogConfigglueDataCatalog schema from the OpenAPI specification
type CreateFlowrequestmetadataCatalogConfigglueDataCatalog struct {
	Rolearn interface{} `json:"roleArn"`
	Tableprefix interface{} `json:"tablePrefix"`
	Databasename interface{} `json:"databaseName"`
}

// SlackMetadata represents the SlackMetadata schema from the OpenAPI specification
type SlackMetadata struct {
	Oauthscopes interface{} `json:"oAuthScopes,omitempty"`
}

// StartFlowRequest represents the StartFlowRequest schema from the OpenAPI specification
type StartFlowRequest struct {
	Flowname interface{} `json:"flowName"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// FieldTypeDetailsfieldLengthRange represents the FieldTypeDetailsfieldLengthRange schema from the OpenAPI specification
type FieldTypeDetailsfieldLengthRange struct {
	Maximum interface{} `json:"maximum,omitempty"`
	Minimum interface{} `json:"minimum,omitempty"`
}

// UpsolverDestinationProperties represents the UpsolverDestinationProperties schema from the OpenAPI specification
type UpsolverDestinationProperties struct {
	Bucketname interface{} `json:"bucketName"`
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	S3outputformatconfig UpsolverDestinationPropertiess3OutputFormatConfig `json:"s3OutputFormatConfig"`
}

// CustomConnectorProfileCredentialsbasic represents the CustomConnectorProfileCredentialsbasic schema from the OpenAPI specification
type CustomConnectorProfileCredentialsbasic struct {
	Password interface{} `json:"password"`
	Username interface{} `json:"username"`
}

// ListConnectorsResponse represents the ListConnectorsResponse schema from the OpenAPI specification
type ListConnectorsResponse struct {
	Connectors interface{} `json:"connectors,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// S3DestinationProperties represents the S3DestinationProperties schema from the OpenAPI specification
type S3DestinationProperties struct {
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	S3outputformatconfig S3OutputFormatConfig `json:"s3OutputFormatConfig,omitempty"` // The configuration that determines how Amazon AppFlow should format the flow output data when Amazon S3 is used as the destination.
	Bucketname interface{} `json:"bucketName"`
}

// ConnectorMetadataZendesk represents the ConnectorMetadataZendesk schema from the OpenAPI specification
type ConnectorMetadataZendesk struct {
	Oauthscopes interface{} `json:"oAuthScopes,omitempty"`
}

// S3OutputFormatConfigprefixConfig represents the S3OutputFormatConfigprefixConfig schema from the OpenAPI specification
type S3OutputFormatConfigprefixConfig struct {
	Prefixformat interface{} `json:"prefixFormat,omitempty"`
	Prefixtype interface{} `json:"prefixType,omitempty"`
	Pathprefixhierarchy interface{} `json:"pathPrefixHierarchy,omitempty"`
}

// SourceConnectorPropertiesTrendmicro represents the SourceConnectorPropertiesTrendmicro schema from the OpenAPI specification
type SourceConnectorPropertiesTrendmicro struct {
	Object interface{} `json:"object"`
}

// SuccessResponseHandlingConfig represents the SuccessResponseHandlingConfig schema from the OpenAPI specification
type SuccessResponseHandlingConfig struct {
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Bucketname interface{} `json:"bucketName,omitempty"`
}

// ListConnectorEntitiesrequest represents the ListConnectorEntitiesrequest schema from the OpenAPI specification
type ListConnectorEntitiesrequest struct {
	Nexttoken string `json:"nextToken,omitempty"` // A token that was provided by your prior <code>ListConnectorEntities</code> operation if the response was too big for the page size. You specify this token to get the next page of results in paginated response.
	Apiversion string `json:"apiVersion,omitempty"` // The version of the API that's used by the connector.
	Connectorprofilename string `json:"connectorProfileName,omitempty"` // The name of the connector profile. The name is unique for each <code>ConnectorProfile</code> in the Amazon Web Services account, and is used to query the downstream connector.
	Connectortype string `json:"connectorType,omitempty"` // The type of connector, such as Salesforce, Amplitude, and so on.
	Entitiespath string `json:"entitiesPath,omitempty"` // This optional parameter is specific to connector implementation. Some connectors support multiple levels or categories of entities. You can find out the list of roots for such providers by sending a request without the <code>entitiesPath</code> parameter. If the connector supports entities at different roots, this initial request returns the list of roots. Otherwise, this request returns all entities supported by the provider.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of items that the operation returns in the response.
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// DeleteFlowResponse represents the DeleteFlowResponse schema from the OpenAPI specification
type DeleteFlowResponse struct {
}

// SingularSourceProperties represents the SingularSourceProperties schema from the OpenAPI specification
type SingularSourceProperties struct {
	Object interface{} `json:"object"`
}

// ApiKeyCredentials represents the ApiKeyCredentials schema from the OpenAPI specification
type ApiKeyCredentials struct {
	Apikey interface{} `json:"apiKey"`
	Apisecretkey interface{} `json:"apiSecretKey,omitempty"`
}

// CustomAuthCredentials represents the CustomAuthCredentials schema from the OpenAPI specification
type CustomAuthCredentials struct {
	Credentialsmap interface{} `json:"credentialsMap,omitempty"`
	Customauthenticationtype interface{} `json:"customAuthenticationType"`
}

// DescribeFlowResponse represents the DescribeFlowResponse schema from the OpenAPI specification
type DescribeFlowResponse struct {
	Description interface{} `json:"description,omitempty"`
	Flowstatus interface{} `json:"flowStatus,omitempty"`
	Flowname interface{} `json:"flowName,omitempty"`
	Metadatacatalogconfig DescribeFlowResponsemetadataCatalogConfig `json:"metadataCatalogConfig,omitempty"`
	Kmsarn interface{} `json:"kmsArn,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Destinationflowconfiglist interface{} `json:"destinationFlowConfigList,omitempty"`
	Flowarn interface{} `json:"flowArn,omitempty"`
	Lastrunexecutiondetails DescribeFlowResponselastRunExecutionDetails `json:"lastRunExecutionDetails,omitempty"`
	Lastrunmetadatacatalogdetails interface{} `json:"lastRunMetadataCatalogDetails,omitempty"`
	Schemaversion interface{} `json:"schemaVersion,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Triggerconfig DescribeFlowResponsetriggerConfig `json:"triggerConfig,omitempty"`
	Lastupdatedby interface{} `json:"lastUpdatedBy,omitempty"`
	Flowstatusmessage interface{} `json:"flowStatusMessage,omitempty"`
	Sourceflowconfig DescribeFlowResponsesourceFlowConfig `json:"sourceFlowConfig,omitempty"`
	Tasks interface{} `json:"tasks,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
}

// CreateFlowRequest represents the CreateFlowRequest schema from the OpenAPI specification
type CreateFlowRequest struct {
	Kmsarn interface{} `json:"kmsArn,omitempty"`
	Sourceflowconfig DescribeFlowResponsesourceFlowConfig `json:"sourceFlowConfig"`
	Tasks interface{} `json:"tasks"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Destinationflowconfiglist interface{} `json:"destinationFlowConfigList"`
	Tags interface{} `json:"tags,omitempty"`
	Triggerconfig DescribeFlowResponsetriggerConfig `json:"triggerConfig"`
	Description interface{} `json:"description,omitempty"`
	Flowname interface{} `json:"flowName"`
	Metadatacatalogconfig DescribeFlowResponsemetadataCatalogConfig `json:"metadataCatalogConfig,omitempty"`
}

// DestinationConnectorPropertiesEventBridge represents the DestinationConnectorPropertiesEventBridge schema from the OpenAPI specification
type DestinationConnectorPropertiesEventBridge struct {
	Errorhandlingconfig ErrorHandlingConfig `json:"errorHandlingConfig,omitempty"` // The settings that determine how Amazon AppFlow handles an error when placing data in the destination. For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. <code>ErrorHandlingConfig</code> is a part of the destination connector details.
	Object interface{} `json:"object"`
}

// TriggerConfig represents the TriggerConfig schema from the OpenAPI specification
type TriggerConfig struct {
	Triggerproperties CreateFlowrequesttriggerConfigtriggerProperties `json:"triggerProperties,omitempty"`
	Triggertype interface{} `json:"triggerType"`
}

// DescribeConnectorsRequest represents the DescribeConnectorsRequest schema from the OpenAPI specification
type DescribeConnectorsRequest struct {
	Connectortypes interface{} `json:"connectorTypes,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// UpdateConnectorRegistrationrequest represents the UpdateConnectorRegistrationrequest schema from the OpenAPI specification
type UpdateConnectorRegistrationrequest struct {
	Description string `json:"description,omitempty"` // A description about the update that you're applying to the connector.
	Clienttoken string `json:"clientToken,omitempty"` // <p>The <code>clientToken</code> parameter is an idempotency token. It ensures that your <code>UpdateConnectorRegistration</code> request completes only once. You choose the value to pass. For example, if you don't receive a response from your request, you can safely retry the request with the same <code>clientToken</code> parameter value.</p> <p>If you omit a <code>clientToken</code> value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.</p> <p>If you specify input parameters that differ from your first request, an error occurs. If you use a different value for <code>clientToken</code>, Amazon AppFlow considers it a new call to <code>UpdateConnectorRegistration</code>. The token is active for 8 hours.</p>
	Connectorlabel string `json:"connectorLabel"` // The name of the connector. The name is unique for each connector registration in your AWS account.
	Connectorprovisioningconfig RegisterConnectorrequestconnectorProvisioningConfig `json:"connectorProvisioningConfig,omitempty"` // Contains information about the configuration of the connector being registered.
}

// CustomConnectorDestinationPropertieserrorHandlingConfig represents the CustomConnectorDestinationPropertieserrorHandlingConfig schema from the OpenAPI specification
type CustomConnectorDestinationPropertieserrorHandlingConfig struct {
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	Failonfirstdestinationerror interface{} `json:"failOnFirstDestinationError,omitempty"`
	Bucketname interface{} `json:"bucketName,omitempty"`
}

// ConnectorProfilePropertiesVeeva represents the ConnectorProfilePropertiesVeeva schema from the OpenAPI specification
type ConnectorProfilePropertiesVeeva struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// DestinationConnectorProperties represents the DestinationConnectorProperties schema from the OpenAPI specification
type DestinationConnectorProperties struct {
	S3 DestinationConnectorPropertiesS3 `json:"S3,omitempty"`
	Upsolver DestinationConnectorPropertiesUpsolver `json:"Upsolver,omitempty"`
	Customerprofiles DestinationConnectorPropertiesCustomerProfiles `json:"CustomerProfiles,omitempty"`
	Honeycode DestinationConnectorPropertiesHoneycode `json:"Honeycode,omitempty"`
	Sapodata DestinationConnectorPropertiesSAPOData `json:"SAPOData,omitempty"`
	Zendesk DestinationConnectorPropertiesZendesk `json:"Zendesk,omitempty"`
	Marketo DestinationConnectorPropertiesMarketo `json:"Marketo,omitempty"`
	Salesforce DestinationConnectorPropertiesSalesforce `json:"Salesforce,omitempty"`
	Snowflake DestinationConnectorPropertiesSnowflake `json:"Snowflake,omitempty"`
	Lookoutmetrics interface{} `json:"LookoutMetrics,omitempty"`
	Customconnector DestinationConnectorPropertiesCustomConnector `json:"CustomConnector,omitempty"`
	Eventbridge DestinationConnectorPropertiesEventBridge `json:"EventBridge,omitempty"`
	Redshift DestinationConnectorPropertiesRedshift `json:"Redshift,omitempty"`
}

// ServiceNowSourceProperties represents the ServiceNowSourceProperties schema from the OpenAPI specification
type ServiceNowSourceProperties struct {
	Object interface{} `json:"object"`
}

// MetadataCatalogDetail represents the MetadataCatalogDetail schema from the OpenAPI specification
type MetadataCatalogDetail struct {
	Catalogtype interface{} `json:"catalogType,omitempty"`
	Partitionregistrationoutput MetadataCatalogDetailpartitionRegistrationOutput `json:"partitionRegistrationOutput,omitempty"`
	Tablename interface{} `json:"tableName,omitempty"`
	Tableregistrationoutput MetadataCatalogDetailtableRegistrationOutput `json:"tableRegistrationOutput,omitempty"`
}

// DestinationConnectorPropertiesMarketo represents the DestinationConnectorPropertiesMarketo schema from the OpenAPI specification
type DestinationConnectorPropertiesMarketo struct {
	Errorhandlingconfig ErrorHandlingConfig `json:"errorHandlingConfig,omitempty"` // The settings that determine how Amazon AppFlow handles an error when placing data in the destination. For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. <code>ErrorHandlingConfig</code> is a part of the destination connector details.
	Object interface{} `json:"object"`
}

// DescribeFlowrequest represents the DescribeFlowrequest schema from the OpenAPI specification
type DescribeFlowrequest struct {
	Flowname string `json:"flowName"` // The specified name of the flow. Spaces are not allowed. Use underscores (_) or hyphens (-) only.
}

// ListFlowsResponse represents the ListFlowsResponse schema from the OpenAPI specification
type ListFlowsResponse struct {
	Flows interface{} `json:"flows,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// SlackSourceProperties represents the SlackSourceProperties schema from the OpenAPI specification
type SlackSourceProperties struct {
	Object interface{} `json:"object"`
}

// ConnectorProfileCredentialsSingular represents the ConnectorProfileCredentialsSingular schema from the OpenAPI specification
type ConnectorProfileCredentialsSingular struct {
	Apikey interface{} `json:"apiKey"`
}

// ConnectorMetadata represents the ConnectorMetadata schema from the OpenAPI specification
type ConnectorMetadata struct {
	Sapodata map[string]interface{} `json:"SAPOData,omitempty"` // The connector metadata specific to SAPOData.
	Salesforce ConnectorMetadataSalesforce `json:"Salesforce,omitempty"`
	Infornexus interface{} `json:"InforNexus,omitempty"`
	Trendmicro interface{} `json:"Trendmicro,omitempty"`
	Amplitude interface{} `json:"Amplitude,omitempty"`
	Dynatrace interface{} `json:"Dynatrace,omitempty"`
	Honeycode ConnectorMetadataHoneycode `json:"Honeycode,omitempty"`
	Upsolver interface{} `json:"Upsolver,omitempty"`
	Servicenow interface{} `json:"ServiceNow,omitempty"`
	Zendesk ConnectorMetadataZendesk `json:"Zendesk,omitempty"`
	Veeva interface{} `json:"Veeva,omitempty"`
	Customerprofiles interface{} `json:"CustomerProfiles,omitempty"`
	Eventbridge interface{} `json:"EventBridge,omitempty"`
	Pardot interface{} `json:"Pardot,omitempty"`
	Redshift interface{} `json:"Redshift,omitempty"`
	Snowflake ConnectorMetadataSnowflake `json:"Snowflake,omitempty"`
	Googleanalytics ConnectorMetadataGoogleAnalytics `json:"GoogleAnalytics,omitempty"`
	Slack ConnectorMetadataSlack `json:"Slack,omitempty"`
	Marketo interface{} `json:"Marketo,omitempty"`
	Singular interface{} `json:"Singular,omitempty"`
	Datadog interface{} `json:"Datadog,omitempty"`
	S3 interface{} `json:"S3,omitempty"`
}

// SalesforceDestinationProperties represents the SalesforceDestinationProperties schema from the OpenAPI specification
type SalesforceDestinationProperties struct {
	Errorhandlingconfig SalesforceDestinationPropertieserrorHandlingConfig `json:"errorHandlingConfig,omitempty"`
	Idfieldnames interface{} `json:"idFieldNames,omitempty"`
	Object interface{} `json:"object"`
	Writeoperationtype interface{} `json:"writeOperationType,omitempty"`
	Datatransferapi interface{} `json:"dataTransferApi,omitempty"`
}

// UnregisterConnectorResponse represents the UnregisterConnectorResponse schema from the OpenAPI specification
type UnregisterConnectorResponse struct {
}

// RegisterConnectorResponse represents the RegisterConnectorResponse schema from the OpenAPI specification
type RegisterConnectorResponse struct {
	Connectorarn interface{} `json:"connectorArn,omitempty"`
}

// ConnectorProfileCredentialsPardot represents the ConnectorProfileCredentialsPardot schema from the OpenAPI specification
type ConnectorProfileCredentialsPardot struct {
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientcredentialsarn interface{} `json:"clientCredentialsArn,omitempty"`
	Oauthrequest ConnectorOAuthRequest `json:"oAuthRequest,omitempty"` // Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
}

// SalesforceConnectorProfileProperties represents the SalesforceConnectorProfileProperties schema from the OpenAPI specification
type SalesforceConnectorProfileProperties struct {
	Useprivatelinkformetadataandauthorization interface{} `json:"usePrivateLinkForMetadataAndAuthorization,omitempty"`
	Instanceurl interface{} `json:"instanceUrl,omitempty"`
	Issandboxenvironment interface{} `json:"isSandboxEnvironment,omitempty"`
}

// AmplitudeSourceProperties represents the AmplitudeSourceProperties schema from the OpenAPI specification
type AmplitudeSourceProperties struct {
	Object interface{} `json:"object"`
}

// DestinationConnectorPropertiesZendesk represents the DestinationConnectorPropertiesZendesk schema from the OpenAPI specification
type DestinationConnectorPropertiesZendesk struct {
	Object interface{} `json:"object"`
	Writeoperationtype string `json:"writeOperationType,omitempty"` // The possible write operations in the destination connector. When this value is not provided, this defaults to the <code>INSERT</code> operation.
	Errorhandlingconfig ErrorHandlingConfig `json:"errorHandlingConfig,omitempty"` // The settings that determine how Amazon AppFlow handles an error when placing data in the destination. For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. <code>ErrorHandlingConfig</code> is a part of the destination connector details.
	Idfieldnames []string `json:"idFieldNames,omitempty"` // A list of field names that can be used as an ID field when performing a write operation.
}

// InforNexusConnectorProfileProperties represents the InforNexusConnectorProfileProperties schema from the OpenAPI specification
type InforNexusConnectorProfileProperties struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// DescribeConnectorEntityRequest represents the DescribeConnectorEntityRequest schema from the OpenAPI specification
type DescribeConnectorEntityRequest struct {
	Connectortype interface{} `json:"connectorType,omitempty"`
	Apiversion interface{} `json:"apiVersion,omitempty"`
	Connectorentityname interface{} `json:"connectorEntityName"`
	Connectorprofilename interface{} `json:"connectorProfileName,omitempty"`
}

// ConnectorProfileCredentialsZendesk represents the ConnectorProfileCredentialsZendesk schema from the OpenAPI specification
type ConnectorProfileCredentialsZendesk struct {
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientid interface{} `json:"clientId"`
	Clientsecret interface{} `json:"clientSecret"`
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
}

// ConnectorProfilePropertiesSlack represents the ConnectorProfilePropertiesSlack schema from the OpenAPI specification
type ConnectorProfilePropertiesSlack struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// ConnectorProfilePropertiesSalesforce represents the ConnectorProfilePropertiesSalesforce schema from the OpenAPI specification
type ConnectorProfilePropertiesSalesforce struct {
	Useprivatelinkformetadataandauthorization interface{} `json:"usePrivateLinkForMetadataAndAuthorization,omitempty"`
	Instanceurl interface{} `json:"instanceUrl,omitempty"`
	Issandboxenvironment interface{} `json:"isSandboxEnvironment,omitempty"`
}

// SlackConnectorProfileCredentials represents the SlackConnectorProfileCredentials schema from the OpenAPI specification
type SlackConnectorProfileCredentials struct {
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientid interface{} `json:"clientId"`
	Clientsecret interface{} `json:"clientSecret"`
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
}

// ResetConnectorMetadataCacheResponse represents the ResetConnectorMetadataCacheResponse schema from the OpenAPI specification
type ResetConnectorMetadataCacheResponse struct {
}

// ConnectorConfigurationsMap represents the ConnectorConfigurationsMap schema from the OpenAPI specification
type ConnectorConfigurationsMap struct {
}

// PardotConnectorProfileCredentials represents the PardotConnectorProfileCredentials schema from the OpenAPI specification
type PardotConnectorProfileCredentials struct {
	Clientcredentialsarn interface{} `json:"clientCredentialsArn,omitempty"`
	Oauthrequest ConnectorOAuthRequest `json:"oAuthRequest,omitempty"` // Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
}

// UpdateConnectorProfileRequest represents the UpdateConnectorProfileRequest schema from the OpenAPI specification
type UpdateConnectorProfileRequest struct {
	Connectorprofilename interface{} `json:"connectorProfileName"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Connectionmode interface{} `json:"connectionMode"`
	Connectorprofileconfig UpdateConnectorProfileRequestconnectorProfileConfig `json:"connectorProfileConfig"`
}

// ListConnectorEntitiesResponse represents the ListConnectorEntitiesResponse schema from the OpenAPI specification
type ListConnectorEntitiesResponse struct {
	Connectorentitymap interface{} `json:"connectorEntityMap"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileCredentials represents the CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileCredentials schema from the OpenAPI specification
type CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileCredentials struct {
	Customconnector CustomConnectorProfileCredentials `json:"CustomConnector,omitempty"` // The connector-specific profile credentials that are required when using the custom connector.
	Servicenow ConnectorProfileCredentialsServiceNow `json:"ServiceNow,omitempty"`
	Slack ConnectorProfileCredentialsSlack `json:"Slack,omitempty"`
	Marketo ConnectorProfileCredentialsMarketo `json:"Marketo,omitempty"`
	Datadog ConnectorProfileCredentialsDatadog `json:"Datadog,omitempty"`
	Googleanalytics ConnectorProfileCredentialsGoogleAnalytics `json:"GoogleAnalytics,omitempty"`
	Zendesk ConnectorProfileCredentialsZendesk `json:"Zendesk,omitempty"`
	Redshift ConnectorProfileCredentialsRedshift `json:"Redshift,omitempty"`
	Singular ConnectorProfileCredentialsSingular `json:"Singular,omitempty"`
	Snowflake ConnectorProfileCredentialsSnowflake `json:"Snowflake,omitempty"`
	Infornexus ConnectorProfileCredentialsInforNexus `json:"InforNexus,omitempty"`
	Amplitude ConnectorProfileCredentialsAmplitude `json:"Amplitude,omitempty"`
	Sapodata SAPODataConnectorProfileCredentials `json:"SAPOData,omitempty"` // The connector-specific profile credentials required when using SAPOData.
	Veeva ConnectorProfileCredentialsVeeva `json:"Veeva,omitempty"`
	Honeycode ConnectorProfileCredentialsHoneycode `json:"Honeycode,omitempty"`
	Trendmicro ConnectorProfileCredentialsTrendmicro `json:"Trendmicro,omitempty"`
	Salesforce ConnectorProfileCredentialsSalesforce `json:"Salesforce,omitempty"`
	Pardot ConnectorProfileCredentialsPardot `json:"Pardot,omitempty"`
	Dynatrace ConnectorProfileCredentialsDynatrace `json:"Dynatrace,omitempty"`
}

// OAuth2Credentials represents the OAuth2Credentials schema from the OpenAPI specification
type OAuth2Credentials struct {
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientid interface{} `json:"clientId,omitempty"`
	Clientsecret interface{} `json:"clientSecret,omitempty"`
	Oauthrequest ConnectorOAuthRequest `json:"oAuthRequest,omitempty"` // Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
}

// ConnectorProfilePropertiesZendesk represents the ConnectorProfilePropertiesZendesk schema from the OpenAPI specification
type ConnectorProfilePropertiesZendesk struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// ListConnectorEntitiesRequest represents the ListConnectorEntitiesRequest schema from the OpenAPI specification
type ListConnectorEntitiesRequest struct {
	Apiversion interface{} `json:"apiVersion,omitempty"`
	Connectorprofilename interface{} `json:"connectorProfileName,omitempty"`
	Connectortype interface{} `json:"connectorType,omitempty"`
	Entitiespath interface{} `json:"entitiesPath,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// CreateFlowrequesttriggerConfigtriggerProperties represents the CreateFlowrequesttriggerConfigtriggerProperties schema from the OpenAPI specification
type CreateFlowrequesttriggerConfigtriggerProperties struct {
	Scheduled TriggerPropertiesScheduled `json:"Scheduled,omitempty"`
}

// VeevaConnectorProfileCredentials represents the VeevaConnectorProfileCredentials schema from the OpenAPI specification
type VeevaConnectorProfileCredentials struct {
	Username interface{} `json:"username"`
	Password interface{} `json:"password"`
}

// RegisterConnectorrequestconnectorProvisioningConfig represents the RegisterConnectorrequestconnectorProvisioningConfig schema from the OpenAPI specification
type RegisterConnectorrequestconnectorProvisioningConfig struct {
	Lambda RegisterConnectorrequestconnectorProvisioningConfiglambda `json:"lambda,omitempty"`
}

// DatadogMetadata represents the DatadogMetadata schema from the OpenAPI specification
type DatadogMetadata struct {
}

// SourceConnectorPropertiesSlack represents the SourceConnectorPropertiesSlack schema from the OpenAPI specification
type SourceConnectorPropertiesSlack struct {
	Object interface{} `json:"object"`
}

// ConnectorProfileProperties represents the ConnectorProfileProperties schema from the OpenAPI specification
type ConnectorProfileProperties struct {
	Trendmicro interface{} `json:"Trendmicro,omitempty"`
	Infornexus ConnectorProfilePropertiesInforNexus `json:"InforNexus,omitempty"`
	Marketo ConnectorProfilePropertiesMarketo `json:"Marketo,omitempty"`
	Servicenow ConnectorProfilePropertiesServiceNow `json:"ServiceNow,omitempty"`
	Zendesk ConnectorProfilePropertiesZendesk `json:"Zendesk,omitempty"`
	Sapodata SAPODataConnectorProfileProperties `json:"SAPOData,omitempty"` // The connector-specific profile properties required when using SAPOData.
	Pardot ConnectorProfilePropertiesPardot `json:"Pardot,omitempty"`
	Salesforce ConnectorProfilePropertiesSalesforce `json:"Salesforce,omitempty"`
	Veeva ConnectorProfilePropertiesVeeva `json:"Veeva,omitempty"`
	Honeycode interface{} `json:"Honeycode,omitempty"`
	Redshift ConnectorProfilePropertiesRedshift `json:"Redshift,omitempty"`
	Slack ConnectorProfilePropertiesSlack `json:"Slack,omitempty"`
	Snowflake ConnectorProfilePropertiesSnowflake `json:"Snowflake,omitempty"`
	Customconnector ConnectorProfilePropertiesCustomConnector `json:"CustomConnector,omitempty"`
	Amplitude interface{} `json:"Amplitude,omitempty"`
	Dynatrace ConnectorProfilePropertiesDynatrace `json:"Dynatrace,omitempty"`
	Googleanalytics interface{} `json:"GoogleAnalytics,omitempty"`
	Datadog ConnectorProfilePropertiesDatadog `json:"Datadog,omitempty"`
	Singular interface{} `json:"Singular,omitempty"`
}

// SupportedFieldTypeDetailsv1 represents the SupportedFieldTypeDetailsv1 schema from the OpenAPI specification
type SupportedFieldTypeDetailsv1 struct {
	Fieldtype interface{} `json:"fieldType"`
	Fieldvaluerange FieldTypeDetailsfieldValueRange `json:"fieldValueRange,omitempty"`
	Filteroperators interface{} `json:"filterOperators"`
	Supporteddateformat interface{} `json:"supportedDateFormat,omitempty"`
	Supportedvalues interface{} `json:"supportedValues,omitempty"`
	Valueregexpattern interface{} `json:"valueRegexPattern,omitempty"`
	Fieldlengthrange FieldTypeDetailsfieldLengthRange `json:"fieldLengthRange,omitempty"`
}

// ExecutionResult represents the ExecutionResult schema from the OpenAPI specification
type ExecutionResult struct {
	Recordsprocessed interface{} `json:"recordsProcessed,omitempty"`
	Bytesprocessed interface{} `json:"bytesProcessed,omitempty"`
	Byteswritten interface{} `json:"bytesWritten,omitempty"`
	Errorinfo ExecutionResulterrorInfo `json:"errorInfo,omitempty"`
}

// CustomConnectorProfileProperties represents the CustomConnectorProfileProperties schema from the OpenAPI specification
type CustomConnectorProfileProperties struct {
	Profileproperties interface{} `json:"profileProperties,omitempty"`
	Oauth2properties OAuth2Properties `json:"oAuth2Properties,omitempty"` // The OAuth 2.0 properties required for OAuth 2.0 authentication.
}

// UpdateConnectorProfileResponse represents the UpdateConnectorProfileResponse schema from the OpenAPI specification
type UpdateConnectorProfileResponse struct {
	Connectorprofilearn interface{} `json:"connectorProfileArn,omitempty"`
}

// SourceConnectorPropertiesVeeva represents the SourceConnectorPropertiesVeeva schema from the OpenAPI specification
type SourceConnectorPropertiesVeeva struct {
	Object interface{} `json:"object"`
	Documenttype interface{} `json:"documentType,omitempty"`
	Includeallversions interface{} `json:"includeAllVersions,omitempty"`
	Includerenditions interface{} `json:"includeRenditions,omitempty"`
	Includesourcefiles interface{} `json:"includeSourceFiles,omitempty"`
}

// DescribeConnectorProfilesResponse represents the DescribeConnectorProfilesResponse schema from the OpenAPI specification
type DescribeConnectorProfilesResponse struct {
	Connectorprofiledetails interface{} `json:"connectorProfileDetails,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// SingularMetadata represents the SingularMetadata schema from the OpenAPI specification
type SingularMetadata struct {
}

// ConnectorProfilePropertiesDynatrace represents the ConnectorProfilePropertiesDynatrace schema from the OpenAPI specification
type ConnectorProfilePropertiesDynatrace struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// DestinationConnectorPropertiesCustomerProfiles represents the DestinationConnectorPropertiesCustomerProfiles schema from the OpenAPI specification
type DestinationConnectorPropertiesCustomerProfiles struct {
	Domainname interface{} `json:"domainName"`
	Objecttypename interface{} `json:"objectTypeName,omitempty"`
}

// CancelFlowExecutionsrequest represents the CancelFlowExecutionsrequest schema from the OpenAPI specification
type CancelFlowExecutionsrequest struct {
	Executionids []string `json:"executionIds,omitempty"` // <p>The ID of each active run to cancel. These runs must belong to the flow you specify in your request.</p> <p>If you omit this parameter, your request ends all active runs that belong to the flow.</p>
	Flowname string `json:"flowName"` // The name of a flow with active runs that you want to cancel.
}

// CancelFlowExecutionsResponse represents the CancelFlowExecutionsResponse schema from the OpenAPI specification
type CancelFlowExecutionsResponse struct {
	Invalidexecutions interface{} `json:"invalidExecutions,omitempty"`
}

// RedshiftMetadata represents the RedshiftMetadata schema from the OpenAPI specification
type RedshiftMetadata struct {
}

// DescribeConnectorProfilesrequest represents the DescribeConnectorProfilesrequest schema from the OpenAPI specification
type DescribeConnectorProfilesrequest struct {
	Nexttoken string `json:"nextToken,omitempty"` // The pagination token for the next page of data.
	Connectorlabel string `json:"connectorLabel,omitempty"` // The name of the connector. The name is unique for each <code>ConnectorRegistration</code> in your Amazon Web Services account. Only needed if calling for CUSTOMCONNECTOR connector type/.
	Connectorprofilenames []string `json:"connectorProfileNames,omitempty"` // The name of the connector profile. The name is unique for each <code>ConnectorProfile</code> in the Amazon Web Services account.
	Connectortype string `json:"connectorType,omitempty"` // The type of connector, such as Salesforce, Amplitude, and so on.
	Maxresults int `json:"maxResults,omitempty"` // Specifies the maximum number of items that should be returned in the result set. The default for <code>maxResults</code> is 20 (for all paginated API operations).
}

// SAPODataConnectorProfileCredentialsoAuthCredentials represents the SAPODataConnectorProfileCredentialsoAuthCredentials schema from the OpenAPI specification
type SAPODataConnectorProfileCredentialsoAuthCredentials struct {
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientid interface{} `json:"clientId"`
	Clientsecret interface{} `json:"clientSecret"`
}

// VeevaSourceProperties represents the VeevaSourceProperties schema from the OpenAPI specification
type VeevaSourceProperties struct {
	Documenttype interface{} `json:"documentType,omitempty"`
	Includeallversions interface{} `json:"includeAllVersions,omitempty"`
	Includerenditions interface{} `json:"includeRenditions,omitempty"`
	Includesourcefiles interface{} `json:"includeSourceFiles,omitempty"`
	Object interface{} `json:"object"`
}

// SourceConnectorPropertiesSingular represents the SourceConnectorPropertiesSingular schema from the OpenAPI specification
type SourceConnectorPropertiesSingular struct {
	Object interface{} `json:"object"`
}

// MarketoDestinationProperties represents the MarketoDestinationProperties schema from the OpenAPI specification
type MarketoDestinationProperties struct {
	Errorhandlingconfig ErrorHandlingConfig `json:"errorHandlingConfig,omitempty"` // The settings that determine how Amazon AppFlow handles an error when placing data in the destination. For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. <code>ErrorHandlingConfig</code> is a part of the destination connector details.
	Object interface{} `json:"object"`
}

// ConnectorOAuthRequest represents the ConnectorOAuthRequest schema from the OpenAPI specification
type ConnectorOAuthRequest struct {
	Authcode interface{} `json:"authCode,omitempty"`
	Redirecturi interface{} `json:"redirectUri,omitempty"`
}

// DynatraceConnectorProfileProperties represents the DynatraceConnectorProfileProperties schema from the OpenAPI specification
type DynatraceConnectorProfileProperties struct {
	Instanceurl interface{} `json:"instanceUrl"`
}

// ConnectorProfileCredentialsServiceNow represents the ConnectorProfileCredentialsServiceNow schema from the OpenAPI specification
type ConnectorProfileCredentialsServiceNow struct {
	Password interface{} `json:"password"`
	Username interface{} `json:"username"`
}

// InforNexusMetadata represents the InforNexusMetadata schema from the OpenAPI specification
type InforNexusMetadata struct {
}

// StopFlowRequest represents the StopFlowRequest schema from the OpenAPI specification
type StopFlowRequest struct {
	Flowname interface{} `json:"flowName"`
}

// InforNexusConnectorProfileCredentials represents the InforNexusConnectorProfileCredentials schema from the OpenAPI specification
type InforNexusConnectorProfileCredentials struct {
	Userid interface{} `json:"userId"`
	Accesskeyid interface{} `json:"accessKeyId"`
	Datakey interface{} `json:"datakey"`
	Secretaccesskey interface{} `json:"secretAccessKey"`
}

// UpdateConnectorProfilerequest represents the UpdateConnectorProfilerequest schema from the OpenAPI specification
type UpdateConnectorProfilerequest struct {
	Clienttoken string `json:"clientToken,omitempty"` // <p>The <code>clientToken</code> parameter is an idempotency token. It ensures that your <code>UpdateConnectorProfile</code> request completes only once. You choose the value to pass. For example, if you don't receive a response from your request, you can safely retry the request with the same <code>clientToken</code> parameter value.</p> <p>If you omit a <code>clientToken</code> value, the Amazon Web Services SDK that you are using inserts a value for you. This way, the SDK can safely retry requests multiple times after a network error. You must provide your own value for other use cases.</p> <p>If you specify input parameters that differ from your first request, an error occurs. If you use a different value for <code>clientToken</code>, Amazon AppFlow considers it a new call to <code>UpdateConnectorProfile</code>. The token is active for 8 hours.</p>
	Connectionmode string `json:"connectionMode"` // Indicates the connection mode and if it is public or private.
	Connectorprofileconfig CreateConnectorProfilerequestconnectorProfileConfig `json:"connectorProfileConfig"` // Defines the connector-specific configuration and credentials for the connector profile.
	Connectorprofilename string `json:"connectorProfileName"` // The name of the connector profile and is unique for each <code>ConnectorProfile</code> in the Amazon Web Services account.
}

// StartFlowResponse represents the StartFlowResponse schema from the OpenAPI specification
type StartFlowResponse struct {
	Flowarn interface{} `json:"flowArn,omitempty"`
	Flowstatus interface{} `json:"flowStatus,omitempty"`
	Executionid interface{} `json:"executionId,omitempty"`
}

// ConnectorProfileCredentialsMarketo represents the ConnectorProfileCredentialsMarketo schema from the OpenAPI specification
type ConnectorProfileCredentialsMarketo struct {
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Clientid interface{} `json:"clientId"`
	Clientsecret interface{} `json:"clientSecret"`
	Oauthrequest GoogleAnalyticsConnectorProfileCredentialsoAuthRequest `json:"oAuthRequest,omitempty"`
}

// LambdaConnectorProvisioningConfig represents the LambdaConnectorProvisioningConfig schema from the OpenAPI specification
type LambdaConnectorProvisioningConfig struct {
	Lambdaarn interface{} `json:"lambdaArn"`
}

// ZendeskMetadata represents the ZendeskMetadata schema from the OpenAPI specification
type ZendeskMetadata struct {
	Oauthscopes interface{} `json:"oAuthScopes,omitempty"`
}

// HoneycodeConnectorProfileCredentials represents the HoneycodeConnectorProfileCredentials schema from the OpenAPI specification
type HoneycodeConnectorProfileCredentials struct {
	Refreshtoken interface{} `json:"refreshToken,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Oauthrequest ConnectorOAuthRequest `json:"oAuthRequest,omitempty"` // Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
}

// S3SourceProperties represents the S3SourceProperties schema from the OpenAPI specification
type S3SourceProperties struct {
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
	S3inputformatconfig S3InputFormatConfig `json:"s3InputFormatConfig,omitempty"` // When you use Amazon S3 as the source, the configuration format that you provide the flow input data.
	Bucketname interface{} `json:"bucketName"`
}

// DestinationFlowConfig represents the DestinationFlowConfig schema from the OpenAPI specification
type DestinationFlowConfig struct {
	Destinationconnectorproperties DestinationFlowConfigdestinationConnectorProperties `json:"destinationConnectorProperties"`
	Apiversion interface{} `json:"apiVersion,omitempty"`
	Connectorprofilename interface{} `json:"connectorProfileName,omitempty"`
	Connectortype interface{} `json:"connectorType"`
}

// DescribeFlowExecutionRecordsrequest represents the DescribeFlowExecutionRecordsrequest schema from the OpenAPI specification
type DescribeFlowExecutionRecordsrequest struct {
	Nexttoken string `json:"nextToken,omitempty"` // The pagination token for the next page of data.
	Flowname string `json:"flowName"` // The specified name of the flow. Spaces are not allowed. Use underscores (_) or hyphens (-) only.
	Maxresults int `json:"maxResults,omitempty"` // Specifies the maximum number of items that should be returned in the result set. The default for <code>maxResults</code> is 20 (for all paginated API operations).
}

// TriggerProperties represents the TriggerProperties schema from the OpenAPI specification
type TriggerProperties struct {
	Scheduled TriggerPropertiesScheduled `json:"Scheduled,omitempty"`
}

// ConnectorProfileCredentialsVeeva represents the ConnectorProfileCredentialsVeeva schema from the OpenAPI specification
type ConnectorProfileCredentialsVeeva struct {
	Username interface{} `json:"username"`
	Password interface{} `json:"password"`
}

// UpdateConnectorRegistrationResponse represents the UpdateConnectorRegistrationResponse schema from the OpenAPI specification
type UpdateConnectorRegistrationResponse struct {
	Connectorarn interface{} `json:"connectorArn,omitempty"`
}

// SingularConnectorProfileCredentials represents the SingularConnectorProfileCredentials schema from the OpenAPI specification
type SingularConnectorProfileCredentials struct {
	Apikey interface{} `json:"apiKey"`
}

// ConnectorProfile represents the ConnectorProfile schema from the OpenAPI specification
type ConnectorProfile struct {
	Connectorprofileproperties CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileProperties `json:"connectorProfileProperties,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Connectionmode interface{} `json:"connectionMode,omitempty"`
	Connectorlabel interface{} `json:"connectorLabel,omitempty"`
	Connectorprofilearn interface{} `json:"connectorProfileArn,omitempty"`
	Connectorprofilename interface{} `json:"connectorProfileName,omitempty"`
	Connectortype interface{} `json:"connectorType,omitempty"`
	Credentialsarn interface{} `json:"credentialsArn,omitempty"`
	Lastupdatedat interface{} `json:"lastUpdatedAt,omitempty"`
	Privateconnectionprovisioningstate ConnectorProfileprivateConnectionProvisioningState `json:"privateConnectionProvisioningState,omitempty"`
}

// CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileProperties represents the CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileProperties schema from the OpenAPI specification
type CreateConnectorProfilerequestconnectorProfileConfigconnectorProfileProperties struct {
	Customconnector ConnectorProfilePropertiesCustomConnector `json:"CustomConnector,omitempty"`
	Amplitude interface{} `json:"Amplitude,omitempty"`
	Dynatrace ConnectorProfilePropertiesDynatrace `json:"Dynatrace,omitempty"`
	Googleanalytics interface{} `json:"GoogleAnalytics,omitempty"`
	Datadog ConnectorProfilePropertiesDatadog `json:"Datadog,omitempty"`
	Singular interface{} `json:"Singular,omitempty"`
	Trendmicro interface{} `json:"Trendmicro,omitempty"`
	Infornexus ConnectorProfilePropertiesInforNexus `json:"InforNexus,omitempty"`
	Marketo ConnectorProfilePropertiesMarketo `json:"Marketo,omitempty"`
	Servicenow ConnectorProfilePropertiesServiceNow `json:"ServiceNow,omitempty"`
	Zendesk ConnectorProfilePropertiesZendesk `json:"Zendesk,omitempty"`
	Sapodata SAPODataConnectorProfileProperties `json:"SAPOData,omitempty"` // The connector-specific profile properties required when using SAPOData.
	Pardot ConnectorProfilePropertiesPardot `json:"Pardot,omitempty"`
	Salesforce ConnectorProfilePropertiesSalesforce `json:"Salesforce,omitempty"`
	Veeva ConnectorProfilePropertiesVeeva `json:"Veeva,omitempty"`
	Honeycode interface{} `json:"Honeycode,omitempty"`
	Redshift ConnectorProfilePropertiesRedshift `json:"Redshift,omitempty"`
	Slack ConnectorProfilePropertiesSlack `json:"Slack,omitempty"`
	Snowflake ConnectorProfilePropertiesSnowflake `json:"Snowflake,omitempty"`
}

// ConnectorProfilePropertiesRedshift represents the ConnectorProfilePropertiesRedshift schema from the OpenAPI specification
type ConnectorProfilePropertiesRedshift struct {
	Clusteridentifier interface{} `json:"clusterIdentifier,omitempty"`
	Databasename interface{} `json:"databaseName,omitempty"`
	Dataapirolearn interface{} `json:"dataApiRoleArn,omitempty"`
	Databaseurl interface{} `json:"databaseUrl,omitempty"`
	Isredshiftserverless interface{} `json:"isRedshiftServerless,omitempty"`
	Workgroupname interface{} `json:"workgroupName,omitempty"`
	Bucketname interface{} `json:"bucketName"`
	Rolearn interface{} `json:"roleArn"`
	Bucketprefix interface{} `json:"bucketPrefix,omitempty"`
}

// OAuth2Defaults represents the OAuth2Defaults schema from the OpenAPI specification
type OAuth2Defaults struct {
	Authcodeurls interface{} `json:"authCodeUrls,omitempty"`
	Oauth2customproperties interface{} `json:"oauth2CustomProperties,omitempty"`
	Oauth2granttypessupported interface{} `json:"oauth2GrantTypesSupported,omitempty"`
	Oauthscopes interface{} `json:"oauthScopes,omitempty"`
	Tokenurls interface{} `json:"tokenUrls,omitempty"`
}

// DescribeConnectorProfilesRequest represents the DescribeConnectorProfilesRequest schema from the OpenAPI specification
type DescribeConnectorProfilesRequest struct {
	Connectortype interface{} `json:"connectorType,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Connectorlabel interface{} `json:"connectorLabel,omitempty"`
	Connectorprofilenames interface{} `json:"connectorProfileNames,omitempty"`
}

// ExecutionDetails represents the ExecutionDetails schema from the OpenAPI specification
type ExecutionDetails struct {
	Mostrecentexecutionstatus interface{} `json:"mostRecentExecutionStatus,omitempty"`
	Mostrecentexecutiontime interface{} `json:"mostRecentExecutionTime,omitempty"`
	Mostrecentexecutionmessage interface{} `json:"mostRecentExecutionMessage,omitempty"`
}
