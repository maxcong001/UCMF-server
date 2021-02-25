# CreateSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfId** | Pointer to **string** |  | [optional] 
**UcmfNotificationUri** | **string** |  | 
**SuggestedExpires** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateSubscription

`func NewCreateSubscription(ucmfNotificationUri string, ) *CreateSubscription`

NewCreateSubscription instantiates a new CreateSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionWithDefaults

`func NewCreateSubscriptionWithDefaults() *CreateSubscription`

NewCreateSubscriptionWithDefaults instantiates a new CreateSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfId

`func (o *CreateSubscription) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *CreateSubscription) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *CreateSubscription) SetNfId(v string)`

SetNfId sets NfId field to given value.

### HasNfId

`func (o *CreateSubscription) HasNfId() bool`

HasNfId returns a boolean if a field has been set.

### GetUcmfNotificationUri

`func (o *CreateSubscription) GetUcmfNotificationUri() string`

GetUcmfNotificationUri returns the UcmfNotificationUri field if non-nil, zero value otherwise.

### GetUcmfNotificationUriOk

`func (o *CreateSubscription) GetUcmfNotificationUriOk() (*string, bool)`

GetUcmfNotificationUriOk returns a tuple with the UcmfNotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcmfNotificationUri

`func (o *CreateSubscription) SetUcmfNotificationUri(v string)`

SetUcmfNotificationUri sets UcmfNotificationUri field to given value.


### GetSuggestedExpires

`func (o *CreateSubscription) GetSuggestedExpires() time.Time`

GetSuggestedExpires returns the SuggestedExpires field if non-nil, zero value otherwise.

### GetSuggestedExpiresOk

`func (o *CreateSubscription) GetSuggestedExpiresOk() (*time.Time, bool)`

GetSuggestedExpiresOk returns a tuple with the SuggestedExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedExpires

`func (o *CreateSubscription) SetSuggestedExpires(v time.Time)`

SetSuggestedExpires sets SuggestedExpires field to given value.

### HasSuggestedExpires

`func (o *CreateSubscription) HasSuggestedExpires() bool`

HasSuggestedExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


