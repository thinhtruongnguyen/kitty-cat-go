# ResponseCheckModelIsServingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ResponseCheckModelIsServingData**](ResponseCheckModelIsServingData.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "success"]

## Methods

### NewResponseCheckModelIsServingResponse

`func NewResponseCheckModelIsServingResponse() *ResponseCheckModelIsServingResponse`

NewResponseCheckModelIsServingResponse instantiates a new ResponseCheckModelIsServingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCheckModelIsServingResponseWithDefaults

`func NewResponseCheckModelIsServingResponseWithDefaults() *ResponseCheckModelIsServingResponse`

NewResponseCheckModelIsServingResponseWithDefaults instantiates a new ResponseCheckModelIsServingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseCheckModelIsServingResponse) GetData() ResponseCheckModelIsServingData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseCheckModelIsServingResponse) GetDataOk() (*ResponseCheckModelIsServingData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseCheckModelIsServingResponse) SetData(v ResponseCheckModelIsServingData)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseCheckModelIsServingResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseCheckModelIsServingResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseCheckModelIsServingResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseCheckModelIsServingResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseCheckModelIsServingResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseCheckModelIsServingResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseCheckModelIsServingResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseCheckModelIsServingResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseCheckModelIsServingResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


