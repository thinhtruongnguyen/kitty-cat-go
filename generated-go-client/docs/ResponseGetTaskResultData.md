# ResponseGetTaskResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ResponseResult**](ResponseResult.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Traceback** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseGetTaskResultData

`func NewResponseGetTaskResultData() *ResponseGetTaskResultData`

NewResponseGetTaskResultData instantiates a new ResponseGetTaskResultData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseGetTaskResultDataWithDefaults

`func NewResponseGetTaskResultDataWithDefaults() *ResponseGetTaskResultData`

NewResponseGetTaskResultDataWithDefaults instantiates a new ResponseGetTaskResultData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ResponseGetTaskResultData) GetResult() ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ResponseGetTaskResultData) GetResultOk() (*ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ResponseGetTaskResultData) SetResult(v ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ResponseGetTaskResultData) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseGetTaskResultData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseGetTaskResultData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseGetTaskResultData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseGetTaskResultData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccess

`func (o *ResponseGetTaskResultData) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseGetTaskResultData) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseGetTaskResultData) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResponseGetTaskResultData) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTraceback

`func (o *ResponseGetTaskResultData) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *ResponseGetTaskResultData) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *ResponseGetTaskResultData) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *ResponseGetTaskResultData) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


