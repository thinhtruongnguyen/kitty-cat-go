# ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TLogs** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseResult

`func NewResponseResult() *ResponseResult`

NewResponseResult instantiates a new ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseResultWithDefaults

`func NewResponseResultWithDefaults() *ResponseResult`

NewResponseResultWithDefaults instantiates a new ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTLogs

`func (o *ResponseResult) GetTLogs() string`

GetTLogs returns the TLogs field if non-nil, zero value otherwise.

### GetTLogsOk

`func (o *ResponseResult) GetTLogsOk() (*string, bool)`

GetTLogsOk returns a tuple with the TLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLogs

`func (o *ResponseResult) SetTLogs(v string)`

SetTLogs sets TLogs field to given value.

### HasTLogs

`func (o *ResponseResult) HasTLogs() bool`

HasTLogs returns a boolean if a field has been set.

### GetOutput

`func (o *ResponseResult) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ResponseResult) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ResponseResult) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ResponseResult) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


