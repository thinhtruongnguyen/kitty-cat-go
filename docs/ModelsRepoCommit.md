# ModelsRepoCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**ModelsCommitUser**](ModelsCommitUser.md) |  | [optional] 
**Committer** | Pointer to [**ModelsCommitUser**](ModelsCommitUser.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Tree** | Pointer to [**ModelsCommitMeta**](ModelsCommitMeta.md) |  | [optional] 

## Methods

### NewModelsRepoCommit

`func NewModelsRepoCommit() *ModelsRepoCommit`

NewModelsRepoCommit instantiates a new ModelsRepoCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsRepoCommitWithDefaults

`func NewModelsRepoCommitWithDefaults() *ModelsRepoCommit`

NewModelsRepoCommitWithDefaults instantiates a new ModelsRepoCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *ModelsRepoCommit) GetAuthor() ModelsCommitUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ModelsRepoCommit) GetAuthorOk() (*ModelsCommitUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ModelsRepoCommit) SetAuthor(v ModelsCommitUser)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ModelsRepoCommit) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCommitter

`func (o *ModelsRepoCommit) GetCommitter() ModelsCommitUser`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *ModelsRepoCommit) GetCommitterOk() (*ModelsCommitUser, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *ModelsRepoCommit) SetCommitter(v ModelsCommitUser)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *ModelsRepoCommit) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.

### GetMessage

`func (o *ModelsRepoCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelsRepoCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelsRepoCommit) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ModelsRepoCommit) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTree

`func (o *ModelsRepoCommit) GetTree() ModelsCommitMeta`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *ModelsRepoCommit) GetTreeOk() (*ModelsCommitMeta, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *ModelsRepoCommit) SetTree(v ModelsCommitMeta)`

SetTree sets Tree field to given value.

### HasTree

`func (o *ModelsRepoCommit) HasTree() bool`

HasTree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


