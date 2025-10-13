# ModelsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRequestToJoin** | Pointer to **bool** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Bio** | Pointer to **string** |  | [optional] 
**Blocked** | Pointer to **bool** |  | [optional] 
**BlockedAt** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Followers** | Pointer to [**[]ModelsFollow**](ModelsFollow.md) |  | [optional] 
**FollowersCount** | Pointer to **int32** |  | [optional] 
**Followings** | Pointer to [**[]ModelsFollow**](ModelsFollow.md) |  | [optional] 
**FollowingsCount** | Pointer to **int32** |  | [optional] 
**GithubLink** | Pointer to **string** |  | [optional] 
**GithubName** | Pointer to **string** |  | [optional] 
**HomePageName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Interests** | Pointer to **string** |  | [optional] 
**InviteOffers** | Pointer to [**[]ModelsOffer**](ModelsOffer.md) |  | [optional] 
**InviteOffersCount** | Pointer to **int32** |  | [optional] 
**IsFollowing** | Pointer to **bool** |  | [optional] 
**JoinId** | Pointer to **string** |  | [optional] 
**JoinOffers** | Pointer to [**[]ModelsOffer**](ModelsOffer.md) |  | [optional] 
**JoinOffersCount** | Pointer to **int32** |  | [optional] 
**Members** | Pointer to [**[]ModelsMember**](ModelsMember.md) |  | [optional] 
**MembersCount** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TwitterLink** | Pointer to **string** |  | [optional] 
**TwitterName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Wallet** | Pointer to [**ModelsWallet**](ModelsWallet.md) |  | [optional] 
**WalletAddress** | Pointer to **string** |  | [optional] 
**WalletConnection** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsUser

`func NewModelsUser() *ModelsUser`

NewModelsUser instantiates a new ModelsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsUserWithDefaults

`func NewModelsUserWithDefaults() *ModelsUser`

NewModelsUserWithDefaults instantiates a new ModelsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowRequestToJoin

`func (o *ModelsUser) GetAllowRequestToJoin() bool`

GetAllowRequestToJoin returns the AllowRequestToJoin field if non-nil, zero value otherwise.

### GetAllowRequestToJoinOk

`func (o *ModelsUser) GetAllowRequestToJoinOk() (*bool, bool)`

GetAllowRequestToJoinOk returns a tuple with the AllowRequestToJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRequestToJoin

`func (o *ModelsUser) SetAllowRequestToJoin(v bool)`

SetAllowRequestToJoin sets AllowRequestToJoin field to given value.

### HasAllowRequestToJoin

`func (o *ModelsUser) HasAllowRequestToJoin() bool`

HasAllowRequestToJoin returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *ModelsUser) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ModelsUser) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ModelsUser) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *ModelsUser) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetBio

`func (o *ModelsUser) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *ModelsUser) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *ModelsUser) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *ModelsUser) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetBlocked

`func (o *ModelsUser) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *ModelsUser) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *ModelsUser) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *ModelsUser) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBlockedAt

`func (o *ModelsUser) GetBlockedAt() string`

GetBlockedAt returns the BlockedAt field if non-nil, zero value otherwise.

### GetBlockedAtOk

`func (o *ModelsUser) GetBlockedAtOk() (*string, bool)`

GetBlockedAtOk returns a tuple with the BlockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedAt

`func (o *ModelsUser) SetBlockedAt(v string)`

SetBlockedAt sets BlockedAt field to given value.

### HasBlockedAt

`func (o *ModelsUser) HasBlockedAt() bool`

HasBlockedAt returns a boolean if a field has been set.

### GetEmail

`func (o *ModelsUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModelsUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFollowers

`func (o *ModelsUser) GetFollowers() []ModelsFollow`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *ModelsUser) GetFollowersOk() (*[]ModelsFollow, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *ModelsUser) SetFollowers(v []ModelsFollow)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *ModelsUser) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetFollowersCount

`func (o *ModelsUser) GetFollowersCount() int32`

GetFollowersCount returns the FollowersCount field if non-nil, zero value otherwise.

### GetFollowersCountOk

`func (o *ModelsUser) GetFollowersCountOk() (*int32, bool)`

GetFollowersCountOk returns a tuple with the FollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersCount

`func (o *ModelsUser) SetFollowersCount(v int32)`

SetFollowersCount sets FollowersCount field to given value.

### HasFollowersCount

`func (o *ModelsUser) HasFollowersCount() bool`

HasFollowersCount returns a boolean if a field has been set.

### GetFollowings

`func (o *ModelsUser) GetFollowings() []ModelsFollow`

GetFollowings returns the Followings field if non-nil, zero value otherwise.

### GetFollowingsOk

`func (o *ModelsUser) GetFollowingsOk() (*[]ModelsFollow, bool)`

GetFollowingsOk returns a tuple with the Followings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowings

`func (o *ModelsUser) SetFollowings(v []ModelsFollow)`

SetFollowings sets Followings field to given value.

### HasFollowings

`func (o *ModelsUser) HasFollowings() bool`

HasFollowings returns a boolean if a field has been set.

### GetFollowingsCount

`func (o *ModelsUser) GetFollowingsCount() int32`

GetFollowingsCount returns the FollowingsCount field if non-nil, zero value otherwise.

### GetFollowingsCountOk

`func (o *ModelsUser) GetFollowingsCountOk() (*int32, bool)`

GetFollowingsCountOk returns a tuple with the FollowingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingsCount

`func (o *ModelsUser) SetFollowingsCount(v int32)`

SetFollowingsCount sets FollowingsCount field to given value.

### HasFollowingsCount

`func (o *ModelsUser) HasFollowingsCount() bool`

HasFollowingsCount returns a boolean if a field has been set.

### GetGithubLink

`func (o *ModelsUser) GetGithubLink() string`

GetGithubLink returns the GithubLink field if non-nil, zero value otherwise.

### GetGithubLinkOk

`func (o *ModelsUser) GetGithubLinkOk() (*string, bool)`

GetGithubLinkOk returns a tuple with the GithubLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLink

`func (o *ModelsUser) SetGithubLink(v string)`

SetGithubLink sets GithubLink field to given value.

### HasGithubLink

`func (o *ModelsUser) HasGithubLink() bool`

HasGithubLink returns a boolean if a field has been set.

### GetGithubName

`func (o *ModelsUser) GetGithubName() string`

GetGithubName returns the GithubName field if non-nil, zero value otherwise.

### GetGithubNameOk

`func (o *ModelsUser) GetGithubNameOk() (*string, bool)`

GetGithubNameOk returns a tuple with the GithubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubName

`func (o *ModelsUser) SetGithubName(v string)`

SetGithubName sets GithubName field to given value.

### HasGithubName

`func (o *ModelsUser) HasGithubName() bool`

HasGithubName returns a boolean if a field has been set.

### GetHomePageName

`func (o *ModelsUser) GetHomePageName() string`

GetHomePageName returns the HomePageName field if non-nil, zero value otherwise.

### GetHomePageNameOk

`func (o *ModelsUser) GetHomePageNameOk() (*string, bool)`

GetHomePageNameOk returns a tuple with the HomePageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageName

`func (o *ModelsUser) SetHomePageName(v string)`

SetHomePageName sets HomePageName field to given value.

### HasHomePageName

`func (o *ModelsUser) HasHomePageName() bool`

HasHomePageName returns a boolean if a field has been set.

### GetId

`func (o *ModelsUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterests

`func (o *ModelsUser) GetInterests() string`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *ModelsUser) GetInterestsOk() (*string, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterests

`func (o *ModelsUser) SetInterests(v string)`

SetInterests sets Interests field to given value.

### HasInterests

`func (o *ModelsUser) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### GetInviteOffers

`func (o *ModelsUser) GetInviteOffers() []ModelsOffer`

GetInviteOffers returns the InviteOffers field if non-nil, zero value otherwise.

### GetInviteOffersOk

`func (o *ModelsUser) GetInviteOffersOk() (*[]ModelsOffer, bool)`

GetInviteOffersOk returns a tuple with the InviteOffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteOffers

`func (o *ModelsUser) SetInviteOffers(v []ModelsOffer)`

SetInviteOffers sets InviteOffers field to given value.

### HasInviteOffers

`func (o *ModelsUser) HasInviteOffers() bool`

HasInviteOffers returns a boolean if a field has been set.

### GetInviteOffersCount

`func (o *ModelsUser) GetInviteOffersCount() int32`

GetInviteOffersCount returns the InviteOffersCount field if non-nil, zero value otherwise.

### GetInviteOffersCountOk

`func (o *ModelsUser) GetInviteOffersCountOk() (*int32, bool)`

GetInviteOffersCountOk returns a tuple with the InviteOffersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteOffersCount

`func (o *ModelsUser) SetInviteOffersCount(v int32)`

SetInviteOffersCount sets InviteOffersCount field to given value.

### HasInviteOffersCount

`func (o *ModelsUser) HasInviteOffersCount() bool`

HasInviteOffersCount returns a boolean if a field has been set.

### GetIsFollowing

`func (o *ModelsUser) GetIsFollowing() bool`

GetIsFollowing returns the IsFollowing field if non-nil, zero value otherwise.

### GetIsFollowingOk

`func (o *ModelsUser) GetIsFollowingOk() (*bool, bool)`

GetIsFollowingOk returns a tuple with the IsFollowing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFollowing

`func (o *ModelsUser) SetIsFollowing(v bool)`

SetIsFollowing sets IsFollowing field to given value.

### HasIsFollowing

`func (o *ModelsUser) HasIsFollowing() bool`

HasIsFollowing returns a boolean if a field has been set.

### GetJoinId

`func (o *ModelsUser) GetJoinId() string`

GetJoinId returns the JoinId field if non-nil, zero value otherwise.

### GetJoinIdOk

`func (o *ModelsUser) GetJoinIdOk() (*string, bool)`

GetJoinIdOk returns a tuple with the JoinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinId

`func (o *ModelsUser) SetJoinId(v string)`

SetJoinId sets JoinId field to given value.

### HasJoinId

`func (o *ModelsUser) HasJoinId() bool`

HasJoinId returns a boolean if a field has been set.

### GetJoinOffers

`func (o *ModelsUser) GetJoinOffers() []ModelsOffer`

GetJoinOffers returns the JoinOffers field if non-nil, zero value otherwise.

### GetJoinOffersOk

`func (o *ModelsUser) GetJoinOffersOk() (*[]ModelsOffer, bool)`

GetJoinOffersOk returns a tuple with the JoinOffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinOffers

`func (o *ModelsUser) SetJoinOffers(v []ModelsOffer)`

SetJoinOffers sets JoinOffers field to given value.

### HasJoinOffers

`func (o *ModelsUser) HasJoinOffers() bool`

HasJoinOffers returns a boolean if a field has been set.

### GetJoinOffersCount

`func (o *ModelsUser) GetJoinOffersCount() int32`

GetJoinOffersCount returns the JoinOffersCount field if non-nil, zero value otherwise.

### GetJoinOffersCountOk

`func (o *ModelsUser) GetJoinOffersCountOk() (*int32, bool)`

GetJoinOffersCountOk returns a tuple with the JoinOffersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinOffersCount

`func (o *ModelsUser) SetJoinOffersCount(v int32)`

SetJoinOffersCount sets JoinOffersCount field to given value.

### HasJoinOffersCount

`func (o *ModelsUser) HasJoinOffersCount() bool`

HasJoinOffersCount returns a boolean if a field has been set.

### GetMembers

`func (o *ModelsUser) GetMembers() []ModelsMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ModelsUser) GetMembersOk() (*[]ModelsMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ModelsUser) SetMembers(v []ModelsMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ModelsUser) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMembersCount

`func (o *ModelsUser) GetMembersCount() int32`

GetMembersCount returns the MembersCount field if non-nil, zero value otherwise.

### GetMembersCountOk

`func (o *ModelsUser) GetMembersCountOk() (*int32, bool)`

GetMembersCountOk returns a tuple with the MembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount

`func (o *ModelsUser) SetMembersCount(v int32)`

SetMembersCount sets MembersCount field to given value.

### HasMembersCount

`func (o *ModelsUser) HasMembersCount() bool`

HasMembersCount returns a boolean if a field has been set.

### GetName

`func (o *ModelsUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *ModelsUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ModelsUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ModelsUser) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ModelsUser) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetToken

`func (o *ModelsUser) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ModelsUser) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ModelsUser) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ModelsUser) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTwitterLink

`func (o *ModelsUser) GetTwitterLink() string`

GetTwitterLink returns the TwitterLink field if non-nil, zero value otherwise.

### GetTwitterLinkOk

`func (o *ModelsUser) GetTwitterLinkOk() (*string, bool)`

GetTwitterLinkOk returns a tuple with the TwitterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterLink

`func (o *ModelsUser) SetTwitterLink(v string)`

SetTwitterLink sets TwitterLink field to given value.

### HasTwitterLink

`func (o *ModelsUser) HasTwitterLink() bool`

HasTwitterLink returns a boolean if a field has been set.

### GetTwitterName

`func (o *ModelsUser) GetTwitterName() string`

GetTwitterName returns the TwitterName field if non-nil, zero value otherwise.

### GetTwitterNameOk

`func (o *ModelsUser) GetTwitterNameOk() (*string, bool)`

GetTwitterNameOk returns a tuple with the TwitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterName

`func (o *ModelsUser) SetTwitterName(v string)`

SetTwitterName sets TwitterName field to given value.

### HasTwitterName

`func (o *ModelsUser) HasTwitterName() bool`

HasTwitterName returns a boolean if a field has been set.

### GetType

`func (o *ModelsUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelsUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelsUser) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelsUser) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVerified

`func (o *ModelsUser) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ModelsUser) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ModelsUser) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ModelsUser) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetVisibility

`func (o *ModelsUser) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ModelsUser) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ModelsUser) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ModelsUser) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetWallet

`func (o *ModelsUser) GetWallet() ModelsWallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *ModelsUser) GetWalletOk() (*ModelsWallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *ModelsUser) SetWallet(v ModelsWallet)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *ModelsUser) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetWalletAddress

`func (o *ModelsUser) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *ModelsUser) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *ModelsUser) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *ModelsUser) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.

### GetWalletConnection

`func (o *ModelsUser) GetWalletConnection() string`

GetWalletConnection returns the WalletConnection field if non-nil, zero value otherwise.

### GetWalletConnectionOk

`func (o *ModelsUser) GetWalletConnectionOk() (*string, bool)`

GetWalletConnectionOk returns a tuple with the WalletConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletConnection

`func (o *ModelsUser) SetWalletConnection(v string)`

SetWalletConnection sets WalletConnection field to given value.

### HasWalletConnection

`func (o *ModelsUser) HasWalletConnection() bool`

HasWalletConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


