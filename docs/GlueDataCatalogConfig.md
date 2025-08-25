

# GlueDataCatalogConfig

<p>Specifies the configuration that Amazon AppFlow uses when it catalogs your data with the Glue Data Catalog. When Amazon AppFlow catalogs your data, it stores metadata in Data Catalog tables. This metadata represents the data that's transferred by the flow that you configure with these settings.</p> <note> <p>You can configure a flow with these settings only when the flow destination is Amazon S3.</p> </note>

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**roleArn** | [**String**](String.md) |  |  |
|**databaseName** | [**String**](String.md) |  |  |
|**tablePrefix** | [**String**](String.md) |  |  |



