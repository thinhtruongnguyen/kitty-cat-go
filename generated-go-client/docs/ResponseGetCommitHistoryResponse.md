# ResponseGetCommitHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ResponseGetCommitHistoryData**](ResponseGetCommitHistoryData.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "success"]

## Methods

### NewResponseGetCommitHistoryResponse

`func NewResponseGetCommitHistoryResponse() *ResponseGetCommitHistoryResponse`

NewResponseGetCommitHistoryResponse instantiates a new ResponseGetCommitHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseGetCommitHistoryResponseWithDefaults

`func NewResponseGetCommitHistoryResponseWithDefaults() *ResponseGetCommitHistoryResponse`

NewResponseGetCommitHistoryResponseWithDefaults instantiates a new ResponseGetCommitHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseGetCommitHistoryResponse) GetData() ResponseGetCommitHistoryData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseGetCommitHistoryResponse) GetDataOk() (*ResponseGetCommitHistoryData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseGetCommitHistoryResponse) SetData(v ResponseGetCommitHistoryData)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseGetCommitHistoryResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseGetCommitHistoryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseGetCommitHistoryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseGetCommitHistoryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseGetCommitHistoryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseGetCommitHistoryResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseGetCommitHistoryResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseGetCommitHistoryResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseGetCommitHistoryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


