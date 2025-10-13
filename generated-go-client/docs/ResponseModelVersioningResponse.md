# ResponseModelVersioningResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ModelsModelVersioning**](ModelsModelVersioning.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "success"]

## Methods

### NewResponseModelVersioningResponse

`func NewResponseModelVersioningResponse() *ResponseModelVersioningResponse`

NewResponseModelVersioningResponse instantiates a new ResponseModelVersioningResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseModelVersioningResponseWithDefaults

`func NewResponseModelVersioningResponseWithDefaults() *ResponseModelVersioningResponse`

NewResponseModelVersioningResponseWithDefaults instantiates a new ResponseModelVersioningResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseModelVersioningResponse) GetData() ModelsModelVersioning`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseModelVersioningResponse) GetDataOk() (*ModelsModelVersioning, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseModelVersioningResponse) SetData(v ModelsModelVersioning)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseModelVersioningResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseModelVersioningResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseModelVersioningResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseModelVersioningResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseModelVersioningResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseModelVersioningResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseModelVersioningResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseModelVersioningResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseModelVersioningResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


