# ResponseGetApiKeyPermissionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyPackages** | Pointer to [**[]ModelsApiKeyPackage**](ModelsApiKeyPackage.md) |  | [optional] 
**LimitModels** | Pointer to **bool** |  | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResponseGetApiKeyPermissionData

`func NewResponseGetApiKeyPermissionData() *ResponseGetApiKeyPermissionData`

NewResponseGetApiKeyPermissionData instantiates a new ResponseGetApiKeyPermissionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseGetApiKeyPermissionDataWithDefaults

`func NewResponseGetApiKeyPermissionDataWithDefaults() *ResponseGetApiKeyPermissionData`

NewResponseGetApiKeyPermissionDataWithDefaults instantiates a new ResponseGetApiKeyPermissionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyPackages

`func (o *ResponseGetApiKeyPermissionData) GetApiKeyPackages() []ModelsApiKeyPackage`

GetApiKeyPackages returns the ApiKeyPackages field if non-nil, zero value otherwise.

### GetApiKeyPackagesOk

`func (o *ResponseGetApiKeyPermissionData) GetApiKeyPackagesOk() (*[]ModelsApiKeyPackage, bool)`

GetApiKeyPackagesOk returns a tuple with the ApiKeyPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyPackages

`func (o *ResponseGetApiKeyPermissionData) SetApiKeyPackages(v []ModelsApiKeyPackage)`

SetApiKeyPackages sets ApiKeyPackages field to given value.

### HasApiKeyPackages

`func (o *ResponseGetApiKeyPermissionData) HasApiKeyPackages() bool`

HasApiKeyPackages returns a boolean if a field has been set.

### GetLimitModels

`func (o *ResponseGetApiKeyPermissionData) GetLimitModels() bool`

GetLimitModels returns the LimitModels field if non-nil, zero value otherwise.

### GetLimitModelsOk

`func (o *ResponseGetApiKeyPermissionData) GetLimitModelsOk() (*bool, bool)`

GetLimitModelsOk returns a tuple with the LimitModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitModels

`func (o *ResponseGetApiKeyPermissionData) SetLimitModels(v bool)`

SetLimitModels sets LimitModels field to given value.

### HasLimitModels

`func (o *ResponseGetApiKeyPermissionData) HasLimitModels() bool`

HasLimitModels returns a boolean if a field has been set.

### GetModels

`func (o *ResponseGetApiKeyPermissionData) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ResponseGetApiKeyPermissionData) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ResponseGetApiKeyPermissionData) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *ResponseGetApiKeyPermissionData) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


