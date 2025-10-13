# ResponseGetApiKeyPermissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ResponseGetApiKeyPermissionData**](ResponseGetApiKeyPermissionData.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "success"]

## Methods

### NewResponseGetApiKeyPermissionResponse

`func NewResponseGetApiKeyPermissionResponse() *ResponseGetApiKeyPermissionResponse`

NewResponseGetApiKeyPermissionResponse instantiates a new ResponseGetApiKeyPermissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseGetApiKeyPermissionResponseWithDefaults

`func NewResponseGetApiKeyPermissionResponseWithDefaults() *ResponseGetApiKeyPermissionResponse`

NewResponseGetApiKeyPermissionResponseWithDefaults instantiates a new ResponseGetApiKeyPermissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseGetApiKeyPermissionResponse) GetData() ResponseGetApiKeyPermissionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseGetApiKeyPermissionResponse) GetDataOk() (*ResponseGetApiKeyPermissionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseGetApiKeyPermissionResponse) SetData(v ResponseGetApiKeyPermissionData)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseGetApiKeyPermissionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseGetApiKeyPermissionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseGetApiKeyPermissionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseGetApiKeyPermissionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseGetApiKeyPermissionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseGetApiKeyPermissionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseGetApiKeyPermissionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseGetApiKeyPermissionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseGetApiKeyPermissionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


