

# ConnectorProfile

 Describes an instance of a connector. This includes the provided name, credentials ARN, connection-mode, and so on. To keep the API intuitive and extensible, the fields that are common to all types of connector profiles are explicitly specified at the top level. The rest of the connector-specific properties are available via the <code>connectorProfileProperties</code> field. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**connectorProfileArn** | [**String**](String.md) |  |  [optional] |
|**connectorProfileName** | [**String**](String.md) |  |  [optional] |
|**connectorType** | [**ConnectorType**](ConnectorType.md) |  |  [optional] |
|**connectorLabel** | [**String**](String.md) |  |  [optional] |
|**connectionMode** | [**ConnectionMode**](ConnectionMode.md) |  |  [optional] |
|**credentialsArn** | [**String**](String.md) |  |  [optional] |
|**connectorProfileProperties** | [**CreateConnectorProfileRequestConnectorProfileConfigConnectorProfileProperties**](CreateConnectorProfileRequestConnectorProfileConfigConnectorProfileProperties.md) |  |  [optional] |
|**createdAt** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**lastUpdatedAt** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**privateConnectionProvisioningState** | [**ConnectorProfilePrivateConnectionProvisioningState**](ConnectorProfilePrivateConnectionProvisioningState.md) |  |  [optional] |



