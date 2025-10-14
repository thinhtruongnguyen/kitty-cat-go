# ModelsModelVersioning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitHash** | Pointer to **string** |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Dependency** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InputFormat** | Pointer to **map[string]interface{}** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**NodeReward** | Pointer to **float32** |  | [optional] 
**OutputFormat** | Pointer to **map[string]interface{}** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**SysRequire** | Pointer to **map[string]interface{}** |  | [optional] 
**TestResult** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VerifyStatus** | Pointer to **string** |  | [optional] 
**VerifyTaskId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsModelVersioning

`func NewModelsModelVersioning() *ModelsModelVersioning`

NewModelsModelVersioning instantiates a new ModelsModelVersioning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsModelVersioningWithDefaults

`func NewModelsModelVersioningWithDefaults() *ModelsModelVersioning`

NewModelsModelVersioningWithDefaults instantiates a new ModelsModelVersioning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitHash

`func (o *ModelsModelVersioning) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ModelsModelVersioning) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ModelsModelVersioning) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *ModelsModelVersioning) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitMessage

`func (o *ModelsModelVersioning) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *ModelsModelVersioning) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *ModelsModelVersioning) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *ModelsModelVersioning) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsModelVersioning) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsModelVersioning) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsModelVersioning) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsModelVersioning) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDependency

`func (o *ModelsModelVersioning) GetDependency() map[string]interface{}`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *ModelsModelVersioning) GetDependencyOk() (*map[string]interface{}, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *ModelsModelVersioning) SetDependency(v map[string]interface{})`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *ModelsModelVersioning) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetId

`func (o *ModelsModelVersioning) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsModelVersioning) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsModelVersioning) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsModelVersioning) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputFormat

`func (o *ModelsModelVersioning) GetInputFormat() map[string]interface{}`

GetInputFormat returns the InputFormat field if non-nil, zero value otherwise.

### GetInputFormatOk

`func (o *ModelsModelVersioning) GetInputFormatOk() (*map[string]interface{}, bool)`

GetInputFormatOk returns a tuple with the InputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFormat

`func (o *ModelsModelVersioning) SetInputFormat(v map[string]interface{})`

SetInputFormat sets InputFormat field to given value.

### HasInputFormat

`func (o *ModelsModelVersioning) HasInputFormat() bool`

HasInputFormat returns a boolean if a field has been set.

### GetModelId

`func (o *ModelsModelVersioning) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ModelsModelVersioning) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ModelsModelVersioning) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *ModelsModelVersioning) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetNodeReward

`func (o *ModelsModelVersioning) GetNodeReward() float32`

GetNodeReward returns the NodeReward field if non-nil, zero value otherwise.

### GetNodeRewardOk

`func (o *ModelsModelVersioning) GetNodeRewardOk() (*float32, bool)`

GetNodeRewardOk returns a tuple with the NodeReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeReward

`func (o *ModelsModelVersioning) SetNodeReward(v float32)`

SetNodeReward sets NodeReward field to given value.

### HasNodeReward

`func (o *ModelsModelVersioning) HasNodeReward() bool`

HasNodeReward returns a boolean if a field has been set.

### GetOutputFormat

`func (o *ModelsModelVersioning) GetOutputFormat() map[string]interface{}`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *ModelsModelVersioning) GetOutputFormatOk() (*map[string]interface{}, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *ModelsModelVersioning) SetOutputFormat(v map[string]interface{})`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *ModelsModelVersioning) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPlatform

`func (o *ModelsModelVersioning) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ModelsModelVersioning) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ModelsModelVersioning) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ModelsModelVersioning) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSysRequire

`func (o *ModelsModelVersioning) GetSysRequire() map[string]interface{}`

GetSysRequire returns the SysRequire field if non-nil, zero value otherwise.

### GetSysRequireOk

`func (o *ModelsModelVersioning) GetSysRequireOk() (*map[string]interface{}, bool)`

GetSysRequireOk returns a tuple with the SysRequire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysRequire

`func (o *ModelsModelVersioning) SetSysRequire(v map[string]interface{})`

SetSysRequire sets SysRequire field to given value.

### HasSysRequire

`func (o *ModelsModelVersioning) HasSysRequire() bool`

HasSysRequire returns a boolean if a field has been set.

### GetTestResult

`func (o *ModelsModelVersioning) GetTestResult() map[string]interface{}`

GetTestResult returns the TestResult field if non-nil, zero value otherwise.

### GetTestResultOk

`func (o *ModelsModelVersioning) GetTestResultOk() (*map[string]interface{}, bool)`

GetTestResultOk returns a tuple with the TestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResult

`func (o *ModelsModelVersioning) SetTestResult(v map[string]interface{})`

SetTestResult sets TestResult field to given value.

### HasTestResult

`func (o *ModelsModelVersioning) HasTestResult() bool`

HasTestResult returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsModelVersioning) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsModelVersioning) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsModelVersioning) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsModelVersioning) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVerifyStatus

`func (o *ModelsModelVersioning) GetVerifyStatus() string`

GetVerifyStatus returns the VerifyStatus field if non-nil, zero value otherwise.

### GetVerifyStatusOk

`func (o *ModelsModelVersioning) GetVerifyStatusOk() (*string, bool)`

GetVerifyStatusOk returns a tuple with the VerifyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyStatus

`func (o *ModelsModelVersioning) SetVerifyStatus(v string)`

SetVerifyStatus sets VerifyStatus field to given value.

### HasVerifyStatus

`func (o *ModelsModelVersioning) HasVerifyStatus() bool`

HasVerifyStatus returns a boolean if a field has been set.

### GetVerifyTaskId

`func (o *ModelsModelVersioning) GetVerifyTaskId() string`

GetVerifyTaskId returns the VerifyTaskId field if non-nil, zero value otherwise.

### GetVerifyTaskIdOk

`func (o *ModelsModelVersioning) GetVerifyTaskIdOk() (*string, bool)`

GetVerifyTaskIdOk returns a tuple with the VerifyTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTaskId

`func (o *ModelsModelVersioning) SetVerifyTaskId(v string)`

SetVerifyTaskId sets VerifyTaskId field to given value.

### HasVerifyTaskId

`func (o *ModelsModelVersioning) HasVerifyTaskId() bool`

HasVerifyTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


