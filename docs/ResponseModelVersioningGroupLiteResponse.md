# ResponseModelVersioningGroupLiteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ModelsModelVersioningGroupLite**](ModelsModelVersioningGroupLite.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "success"]

## Methods

### NewResponseModelVersioningGroupLiteResponse

`func NewResponseModelVersioningGroupLiteResponse() *ResponseModelVersioningGroupLiteResponse`

NewResponseModelVersioningGroupLiteResponse instantiates a new ResponseModelVersioningGroupLiteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseModelVersioningGroupLiteResponseWithDefaults

`func NewResponseModelVersioningGroupLiteResponseWithDefaults() *ResponseModelVersioningGroupLiteResponse`

NewResponseModelVersioningGroupLiteResponseWithDefaults instantiates a new ResponseModelVersioningGroupLiteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseModelVersioningGroupLiteResponse) GetData() ModelsModelVersioningGroupLite`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseModelVersioningGroupLiteResponse) GetDataOk() (*ModelsModelVersioningGroupLite, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseModelVersioningGroupLiteResponse) SetData(v ModelsModelVersioningGroupLite)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseModelVersioningGroupLiteResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseModelVersioningGroupLiteResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseModelVersioningGroupLiteResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseModelVersioningGroupLiteResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseModelVersioningGroupLiteResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseModelVersioningGroupLiteResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseModelVersioningGroupLiteResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseModelVersioningGroupLiteResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseModelVersioningGroupLiteResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


