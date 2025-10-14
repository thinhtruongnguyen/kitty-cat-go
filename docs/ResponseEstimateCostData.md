# ResponseEstimateCostData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **float32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] [default to "$"]
**Unit** | Pointer to **string** |  | [optional] [default to "USD"]

## Methods

### NewResponseEstimateCostData

`func NewResponseEstimateCostData() *ResponseEstimateCostData`

NewResponseEstimateCostData instantiates a new ResponseEstimateCostData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseEstimateCostDataWithDefaults

`func NewResponseEstimateCostDataWithDefaults() *ResponseEstimateCostData`

NewResponseEstimateCostDataWithDefaults instantiates a new ResponseEstimateCostData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *ResponseEstimateCostData) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ResponseEstimateCostData) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ResponseEstimateCostData) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ResponseEstimateCostData) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetSymbol

`func (o *ResponseEstimateCostData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ResponseEstimateCostData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ResponseEstimateCostData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ResponseEstimateCostData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetUnit

`func (o *ResponseEstimateCostData) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ResponseEstimateCostData) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ResponseEstimateCostData) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ResponseEstimateCostData) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


