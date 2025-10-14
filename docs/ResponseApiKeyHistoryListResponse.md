# ResponseApiKeyHistoryListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ResponseApiKeyHistoryListData**](ResponseApiKeyHistoryListData.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "success"]

## Methods

### NewResponseApiKeyHistoryListResponse

`func NewResponseApiKeyHistoryListResponse() *ResponseApiKeyHistoryListResponse`

NewResponseApiKeyHistoryListResponse instantiates a new ResponseApiKeyHistoryListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseApiKeyHistoryListResponseWithDefaults

`func NewResponseApiKeyHistoryListResponseWithDefaults() *ResponseApiKeyHistoryListResponse`

NewResponseApiKeyHistoryListResponseWithDefaults instantiates a new ResponseApiKeyHistoryListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseApiKeyHistoryListResponse) GetData() ResponseApiKeyHistoryListData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseApiKeyHistoryListResponse) GetDataOk() (*ResponseApiKeyHistoryListData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseApiKeyHistoryListResponse) SetData(v ResponseApiKeyHistoryListData)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseApiKeyHistoryListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseApiKeyHistoryListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseApiKeyHistoryListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseApiKeyHistoryListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseApiKeyHistoryListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseApiKeyHistoryListResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseApiKeyHistoryListResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseApiKeyHistoryListResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseApiKeyHistoryListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


