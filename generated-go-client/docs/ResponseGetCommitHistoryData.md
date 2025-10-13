# ResponseGetCommitHistoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to [**[]ModelsCommit**](ModelsCommit.md) |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**LastPage** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseGetCommitHistoryData

`func NewResponseGetCommitHistoryData() *ResponseGetCommitHistoryData`

NewResponseGetCommitHistoryData instantiates a new ResponseGetCommitHistoryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseGetCommitHistoryDataWithDefaults

`func NewResponseGetCommitHistoryDataWithDefaults() *ResponseGetCommitHistoryData`

NewResponseGetCommitHistoryDataWithDefaults instantiates a new ResponseGetCommitHistoryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *ResponseGetCommitHistoryData) GetCommit() []ModelsCommit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ResponseGetCommitHistoryData) GetCommitOk() (*[]ModelsCommit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ResponseGetCommitHistoryData) SetCommit(v []ModelsCommit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ResponseGetCommitHistoryData) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetHasMore

`func (o *ResponseGetCommitHistoryData) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ResponseGetCommitHistoryData) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ResponseGetCommitHistoryData) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ResponseGetCommitHistoryData) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetLastPage

`func (o *ResponseGetCommitHistoryData) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *ResponseGetCommitHistoryData) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *ResponseGetCommitHistoryData) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *ResponseGetCommitHistoryData) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


