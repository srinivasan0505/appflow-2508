

# ScheduledTriggerProperties

 Specifies the configuration details of a schedule-triggered flow as defined by the user. Currently, these settings only apply to the <code>Scheduled</code> trigger type. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**scheduleExpression** | [**String**](String.md) |  |  |
|**dataPullMode** | [**DataPullMode**](DataPullMode.md) |  |  [optional] |
|**scheduleStartTime** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**scheduleEndTime** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**timezone** | [**String**](String.md) |  |  [optional] |
|**scheduleOffset** | [**Integer**](Integer.md) |  |  [optional] |
|**firstExecutionFrom** | [**OffsetDateTime**](OffsetDateTime.md) |  |  [optional] |
|**flowErrorDeactivationThreshold** | [**Integer**](Integer.md) |  |  [optional] |



