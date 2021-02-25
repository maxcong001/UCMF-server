# CreatedSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DicEntryId** | **int32** |  | 
**ConfirmedExpires** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreatedSubscription

`func NewCreatedSubscription(dicEntryId int32, ) *CreatedSubscription`

NewCreatedSubscription instantiates a new CreatedSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedSubscriptionWithDefaults

`func NewCreatedSubscriptionWithDefaults() *CreatedSubscription`

NewCreatedSubscriptionWithDefaults instantiates a new CreatedSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDicEntryId

`func (o *CreatedSubscription) GetDicEntryId() int32`

GetDicEntryId returns the DicEntryId field if non-nil, zero value otherwise.

### GetDicEntryIdOk

`func (o *CreatedSubscription) GetDicEntryIdOk() (*int32, bool)`

GetDicEntryIdOk returns a tuple with the DicEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDicEntryId

`func (o *CreatedSubscription) SetDicEntryId(v int32)`

SetDicEntryId sets DicEntryId field to given value.


### GetConfirmedExpires

`func (o *CreatedSubscription) GetConfirmedExpires() time.Time`

GetConfirmedExpires returns the ConfirmedExpires field if non-nil, zero value otherwise.

### GetConfirmedExpiresOk

`func (o *CreatedSubscription) GetConfirmedExpiresOk() (*time.Time, bool)`

GetConfirmedExpiresOk returns a tuple with the ConfirmedExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedExpires

`func (o *CreatedSubscription) SetConfirmedExpires(v time.Time)`

SetConfirmedExpires sets ConfirmedExpires field to given value.

### HasConfirmedExpires

`func (o *CreatedSubscription) HasConfirmedExpires() bool`

HasConfirmedExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


