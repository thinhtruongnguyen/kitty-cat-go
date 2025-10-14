# RequestVerifyAiModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitHash** | Pointer to **string** |  | [optional] 
**Platforms** | **[]string** |  | 

## Methods

### NewRequestVerifyAiModelRequest

`func NewRequestVerifyAiModelRequest(platforms []string, ) *RequestVerifyAiModelRequest`

NewRequestVerifyAiModelRequest instantiates a new RequestVerifyAiModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestVerifyAiModelRequestWithDefaults

`func NewRequestVerifyAiModelRequestWithDefaults() *RequestVerifyAiModelRequest`

NewRequestVerifyAiModelRequestWithDefaults instantiates a new RequestVerifyAiModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitHash

`func (o *RequestVerifyAiModelRequest) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *RequestVerifyAiModelRequest) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *RequestVerifyAiModelRequest) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *RequestVerifyAiModelRequest) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetPlatforms

`func (o *RequestVerifyAiModelRequest) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *RequestVerifyAiModelRequest) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *RequestVerifyAiModelRequest) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


