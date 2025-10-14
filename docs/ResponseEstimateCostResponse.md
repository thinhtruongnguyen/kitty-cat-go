# ResponseEstimateCostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ResponseEstimateCostData**](ResponseEstimateCostData.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "success"]

## Methods

### NewResponseEstimateCostResponse

`func NewResponseEstimateCostResponse() *ResponseEstimateCostResponse`

NewResponseEstimateCostResponse instantiates a new ResponseEstimateCostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseEstimateCostResponseWithDefaults

`func NewResponseEstimateCostResponseWithDefaults() *ResponseEstimateCostResponse`

NewResponseEstimateCostResponseWithDefaults instantiates a new ResponseEstimateCostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseEstimateCostResponse) GetData() ResponseEstimateCostData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseEstimateCostResponse) GetDataOk() (*ResponseEstimateCostData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseEstimateCostResponse) SetData(v ResponseEstimateCostData)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseEstimateCostResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseEstimateCostResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseEstimateCostResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseEstimateCostResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseEstimateCostResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseEstimateCostResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseEstimateCostResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseEstimateCostResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseEstimateCostResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


