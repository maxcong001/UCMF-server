# DicEntryCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeAllocationCode** | **string** |  | 
**UeRadioCapability5GS** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UeRadioCapabilityEPS** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 

## Methods

### NewDicEntryCreateData

`func NewDicEntryCreateData(typeAllocationCode string, ) *DicEntryCreateData`

NewDicEntryCreateData instantiates a new DicEntryCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDicEntryCreateDataWithDefaults

`func NewDicEntryCreateDataWithDefaults() *DicEntryCreateData`

NewDicEntryCreateDataWithDefaults instantiates a new DicEntryCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeAllocationCode

`func (o *DicEntryCreateData) GetTypeAllocationCode() string`

GetTypeAllocationCode returns the TypeAllocationCode field if non-nil, zero value otherwise.

### GetTypeAllocationCodeOk

`func (o *DicEntryCreateData) GetTypeAllocationCodeOk() (*string, bool)`

GetTypeAllocationCodeOk returns a tuple with the TypeAllocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeAllocationCode

`func (o *DicEntryCreateData) SetTypeAllocationCode(v string)`

SetTypeAllocationCode sets TypeAllocationCode field to given value.


### GetUeRadioCapability5GS

`func (o *DicEntryCreateData) GetUeRadioCapability5GS() RefToBinaryData`

GetUeRadioCapability5GS returns the UeRadioCapability5GS field if non-nil, zero value otherwise.

### GetUeRadioCapability5GSOk

`func (o *DicEntryCreateData) GetUeRadioCapability5GSOk() (*RefToBinaryData, bool)`

GetUeRadioCapability5GSOk returns a tuple with the UeRadioCapability5GS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCapability5GS

`func (o *DicEntryCreateData) SetUeRadioCapability5GS(v RefToBinaryData)`

SetUeRadioCapability5GS sets UeRadioCapability5GS field to given value.

### HasUeRadioCapability5GS

`func (o *DicEntryCreateData) HasUeRadioCapability5GS() bool`

HasUeRadioCapability5GS returns a boolean if a field has been set.

### GetUeRadioCapabilityEPS

`func (o *DicEntryCreateData) GetUeRadioCapabilityEPS() RefToBinaryData`

GetUeRadioCapabilityEPS returns the UeRadioCapabilityEPS field if non-nil, zero value otherwise.

### GetUeRadioCapabilityEPSOk

`func (o *DicEntryCreateData) GetUeRadioCapabilityEPSOk() (*RefToBinaryData, bool)`

GetUeRadioCapabilityEPSOk returns a tuple with the UeRadioCapabilityEPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCapabilityEPS

`func (o *DicEntryCreateData) SetUeRadioCapabilityEPS(v RefToBinaryData)`

SetUeRadioCapabilityEPS sets UeRadioCapabilityEPS field to given value.

### HasUeRadioCapabilityEPS

`func (o *DicEntryCreateData) HasUeRadioCapabilityEPS() bool`

HasUeRadioCapabilityEPS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


