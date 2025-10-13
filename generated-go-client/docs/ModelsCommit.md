# ModelsCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**ModelsUser**](ModelsUser.md) |  | [optional] 
**CommitAffectedFiles** | Pointer to **[]string** |  | [optional] 
**CommitData** | Pointer to [**ModelsUser**](ModelsUser.md) |  | [optional] 
**CommitMeta** | Pointer to [**ModelsCommitMeta**](ModelsCommitMeta.md) |  | [optional] 
**Parents** | Pointer to [**[]ModelsCommitMeta**](ModelsCommitMeta.md) |  | [optional] 
**RepoCommit** | Pointer to [**ModelsRepoCommit**](ModelsRepoCommit.md) |  | [optional] 
**Stats** | Pointer to [**ModelsCommitStats**](ModelsCommitStats.md) |  | [optional] 

## Methods

### NewModelsCommit

`func NewModelsCommit() *ModelsCommit`

NewModelsCommit instantiates a new ModelsCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCommitWithDefaults

`func NewModelsCommitWithDefaults() *ModelsCommit`

NewModelsCommitWithDefaults instantiates a new ModelsCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *ModelsCommit) GetAuthor() ModelsUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ModelsCommit) GetAuthorOk() (*ModelsUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ModelsCommit) SetAuthor(v ModelsUser)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ModelsCommit) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCommitAffectedFiles

`func (o *ModelsCommit) GetCommitAffectedFiles() []string`

GetCommitAffectedFiles returns the CommitAffectedFiles field if non-nil, zero value otherwise.

### GetCommitAffectedFilesOk

`func (o *ModelsCommit) GetCommitAffectedFilesOk() (*[]string, bool)`

GetCommitAffectedFilesOk returns a tuple with the CommitAffectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitAffectedFiles

`func (o *ModelsCommit) SetCommitAffectedFiles(v []string)`

SetCommitAffectedFiles sets CommitAffectedFiles field to given value.

### HasCommitAffectedFiles

`func (o *ModelsCommit) HasCommitAffectedFiles() bool`

HasCommitAffectedFiles returns a boolean if a field has been set.

### GetCommitData

`func (o *ModelsCommit) GetCommitData() ModelsUser`

GetCommitData returns the CommitData field if non-nil, zero value otherwise.

### GetCommitDataOk

`func (o *ModelsCommit) GetCommitDataOk() (*ModelsUser, bool)`

GetCommitDataOk returns a tuple with the CommitData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitData

`func (o *ModelsCommit) SetCommitData(v ModelsUser)`

SetCommitData sets CommitData field to given value.

### HasCommitData

`func (o *ModelsCommit) HasCommitData() bool`

HasCommitData returns a boolean if a field has been set.

### GetCommitMeta

`func (o *ModelsCommit) GetCommitMeta() ModelsCommitMeta`

GetCommitMeta returns the CommitMeta field if non-nil, zero value otherwise.

### GetCommitMetaOk

`func (o *ModelsCommit) GetCommitMetaOk() (*ModelsCommitMeta, bool)`

GetCommitMetaOk returns a tuple with the CommitMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMeta

`func (o *ModelsCommit) SetCommitMeta(v ModelsCommitMeta)`

SetCommitMeta sets CommitMeta field to given value.

### HasCommitMeta

`func (o *ModelsCommit) HasCommitMeta() bool`

HasCommitMeta returns a boolean if a field has been set.

### GetParents

`func (o *ModelsCommit) GetParents() []ModelsCommitMeta`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *ModelsCommit) GetParentsOk() (*[]ModelsCommitMeta, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *ModelsCommit) SetParents(v []ModelsCommitMeta)`

SetParents sets Parents field to given value.

### HasParents

`func (o *ModelsCommit) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetRepoCommit

`func (o *ModelsCommit) GetRepoCommit() ModelsRepoCommit`

GetRepoCommit returns the RepoCommit field if non-nil, zero value otherwise.

### GetRepoCommitOk

`func (o *ModelsCommit) GetRepoCommitOk() (*ModelsRepoCommit, bool)`

GetRepoCommitOk returns a tuple with the RepoCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoCommit

`func (o *ModelsCommit) SetRepoCommit(v ModelsRepoCommit)`

SetRepoCommit sets RepoCommit field to given value.

### HasRepoCommit

`func (o *ModelsCommit) HasRepoCommit() bool`

HasRepoCommit returns a boolean if a field has been set.

### GetStats

`func (o *ModelsCommit) GetStats() ModelsCommitStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ModelsCommit) GetStatsOk() (*ModelsCommitStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ModelsCommit) SetStats(v ModelsCommitStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ModelsCommit) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


