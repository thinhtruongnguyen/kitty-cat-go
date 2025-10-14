# ResponseModelVersioningGroupLiteListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ResponseModelVersioningGroupLiteListData**](ResponseModelVersioningGroupLiteListData.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "success"]

## Methods

### NewResponseModelVersioningGroupLiteListResponse

`func NewResponseModelVersioningGroupLiteListResponse() *ResponseModelVersioningGroupLiteListResponse`

NewResponseModelVersioningGroupLiteListResponse instantiates a new ResponseModelVersioningGroupLiteListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseModelVersioningGroupLiteListResponseWithDefaults

`func NewResponseModelVersioningGroupLiteListResponseWithDefaults() *ResponseModelVersioningGroupLiteListResponse`

NewResponseModelVersioningGroupLiteListResponseWithDefaults instantiates a new ResponseModelVersioningGroupLiteListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseModelVersioningGroupLiteListResponse) GetData() ResponseModelVersioningGroupLiteListData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseModelVersioningGroupLiteListResponse) GetDataOk() (*ResponseModelVersioningGroupLiteListData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseModelVersioningGroupLiteListResponse) SetData(v ResponseModelVersioningGroupLiteListData)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseModelVersioningGroupLiteListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseModelVersioningGroupLiteListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseModelVersioningGroupLiteListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseModelVersioningGroupLiteListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseModelVersioningGroupLiteListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseModelVersioningGroupLiteListResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseModelVersioningGroupLiteListResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseModelVersioningGroupLiteListResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseModelVersioningGroupLiteListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


