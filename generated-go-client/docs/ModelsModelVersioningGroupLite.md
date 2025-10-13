# ModelsModelVersioningGroupLite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitHash** | Pointer to **string** | Dependency        map[string]interface{} &#x60;json:\&quot;dependency\&quot;&#x60; | [optional] 
**CommitMessage** | Pointer to **string** | TestResult        map[string]interface{} &#x60;json:\&quot;test_result\&quot;&#x60; InputFormat       map[string]interface{} &#x60;json:\&quot;input_format\&quot;&#x60; OutputFormat      map[string]interface{} &#x60;json:\&quot;output_format\&quot;&#x60; SysRequired       map[string]interface{} &#x60;json:\&quot;sys_require\&quot;&#x60; | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**LastCheckedAt** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**PendingPlatforms** | Pointer to [**[]ModelsPlatformTask**](ModelsPlatformTask.md) |  | [optional] 
**RejectedPlatforms** | Pointer to [**[]ModelsPlatformTask**](ModelsPlatformTask.md) |  | [optional] 
**VerifiedPlatforms** | Pointer to [**[]ModelsPlatformTask**](ModelsPlatformTask.md) |  | [optional] 
**VerifyStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsModelVersioningGroupLite

`func NewModelsModelVersioningGroupLite() *ModelsModelVersioningGroupLite`

NewModelsModelVersioningGroupLite instantiates a new ModelsModelVersioningGroupLite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsModelVersioningGroupLiteWithDefaults

`func NewModelsModelVersioningGroupLiteWithDefaults() *ModelsModelVersioningGroupLite`

NewModelsModelVersioningGroupLiteWithDefaults instantiates a new ModelsModelVersioningGroupLite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitHash

`func (o *ModelsModelVersioningGroupLite) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ModelsModelVersioningGroupLite) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ModelsModelVersioningGroupLite) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *ModelsModelVersioningGroupLite) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitMessage

`func (o *ModelsModelVersioningGroupLite) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *ModelsModelVersioningGroupLite) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *ModelsModelVersioningGroupLite) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *ModelsModelVersioningGroupLite) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetIsActive

`func (o *ModelsModelVersioningGroupLite) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ModelsModelVersioningGroupLite) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ModelsModelVersioningGroupLite) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ModelsModelVersioningGroupLite) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastCheckedAt

`func (o *ModelsModelVersioningGroupLite) GetLastCheckedAt() string`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *ModelsModelVersioningGroupLite) GetLastCheckedAtOk() (*string, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *ModelsModelVersioningGroupLite) SetLastCheckedAt(v string)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *ModelsModelVersioningGroupLite) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### GetModelId

`func (o *ModelsModelVersioningGroupLite) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ModelsModelVersioningGroupLite) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ModelsModelVersioningGroupLite) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *ModelsModelVersioningGroupLite) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetPendingPlatforms

`func (o *ModelsModelVersioningGroupLite) GetPendingPlatforms() []ModelsPlatformTask`

GetPendingPlatforms returns the PendingPlatforms field if non-nil, zero value otherwise.

### GetPendingPlatformsOk

`func (o *ModelsModelVersioningGroupLite) GetPendingPlatformsOk() (*[]ModelsPlatformTask, bool)`

GetPendingPlatformsOk returns a tuple with the PendingPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPlatforms

`func (o *ModelsModelVersioningGroupLite) SetPendingPlatforms(v []ModelsPlatformTask)`

SetPendingPlatforms sets PendingPlatforms field to given value.

### HasPendingPlatforms

`func (o *ModelsModelVersioningGroupLite) HasPendingPlatforms() bool`

HasPendingPlatforms returns a boolean if a field has been set.

### GetRejectedPlatforms

`func (o *ModelsModelVersioningGroupLite) GetRejectedPlatforms() []ModelsPlatformTask`

GetRejectedPlatforms returns the RejectedPlatforms field if non-nil, zero value otherwise.

### GetRejectedPlatformsOk

`func (o *ModelsModelVersioningGroupLite) GetRejectedPlatformsOk() (*[]ModelsPlatformTask, bool)`

GetRejectedPlatformsOk returns a tuple with the RejectedPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedPlatforms

`func (o *ModelsModelVersioningGroupLite) SetRejectedPlatforms(v []ModelsPlatformTask)`

SetRejectedPlatforms sets RejectedPlatforms field to given value.

### HasRejectedPlatforms

`func (o *ModelsModelVersioningGroupLite) HasRejectedPlatforms() bool`

HasRejectedPlatforms returns a boolean if a field has been set.

### GetVerifiedPlatforms

`func (o *ModelsModelVersioningGroupLite) GetVerifiedPlatforms() []ModelsPlatformTask`

GetVerifiedPlatforms returns the VerifiedPlatforms field if non-nil, zero value otherwise.

### GetVerifiedPlatformsOk

`func (o *ModelsModelVersioningGroupLite) GetVerifiedPlatformsOk() (*[]ModelsPlatformTask, bool)`

GetVerifiedPlatformsOk returns a tuple with the VerifiedPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPlatforms

`func (o *ModelsModelVersioningGroupLite) SetVerifiedPlatforms(v []ModelsPlatformTask)`

SetVerifiedPlatforms sets VerifiedPlatforms field to given value.

### HasVerifiedPlatforms

`func (o *ModelsModelVersioningGroupLite) HasVerifiedPlatforms() bool`

HasVerifiedPlatforms returns a boolean if a field has been set.

### GetVerifyStatus

`func (o *ModelsModelVersioningGroupLite) GetVerifyStatus() string`

GetVerifyStatus returns the VerifyStatus field if non-nil, zero value otherwise.

### GetVerifyStatusOk

`func (o *ModelsModelVersioningGroupLite) GetVerifyStatusOk() (*string, bool)`

GetVerifyStatusOk returns a tuple with the VerifyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyStatus

`func (o *ModelsModelVersioningGroupLite) SetVerifyStatus(v string)`

SetVerifyStatus sets VerifyStatus field to given value.

### HasVerifyStatus

`func (o *ModelsModelVersioningGroupLite) HasVerifyStatus() bool`

HasVerifyStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


