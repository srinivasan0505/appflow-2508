

# RegistrationOutput

<p>Describes the status of an attempt from Amazon AppFlow to register a resource.</p> <p>When you run a flow that you've configured to use a metadata catalog, Amazon AppFlow registers a metadata table and data partitions with that catalog. This operation provides the status of that registration attempt. The operation also indicates how many related resources Amazon AppFlow created or updated.</p>

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**message** | [**String**](String.md) |  |  [optional] |
|**result** | [**String**](String.md) |  |  [optional] |
|**status** | [**ExecutionStatus**](ExecutionStatus.md) |  |  [optional] |



