/*
 * Nucmf_UECapabilityManagement
 *
 * Nucmf_UECapabilityManagement Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"os"
)

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	JsonData *DicEntryData `json:"jsonData,omitempty"`
	BinaryDataUeRadioCapability5GS **os.File `json:"binaryDataUeRadioCapability5GS,omitempty"`
	BinaryDataUeRadioCapabilityEPS **os.File `json:"binaryDataUeRadioCapabilityEPS,omitempty"`
}

// NewInlineResponse200 instantiates a new InlineResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200WithDefaults() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *InlineResponse200) GetJsonData() DicEntryData {
	if o == nil || o.JsonData == nil {
		var ret DicEntryData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetJsonDataOk() (*DicEntryData, bool) {
	if o == nil || o.JsonData == nil {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *InlineResponse200) HasJsonData() bool {
	if o != nil && o.JsonData != nil {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given DicEntryData and assigns it to the JsonData field.
func (o *InlineResponse200) SetJsonData(v DicEntryData) {
	o.JsonData = &v
}

// GetBinaryDataUeRadioCapability5GS returns the BinaryDataUeRadioCapability5GS field value if set, zero value otherwise.
func (o *InlineResponse200) GetBinaryDataUeRadioCapability5GS() *os.File {
	if o == nil || o.BinaryDataUeRadioCapability5GS == nil {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataUeRadioCapability5GS
}

// GetBinaryDataUeRadioCapability5GSOk returns a tuple with the BinaryDataUeRadioCapability5GS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetBinaryDataUeRadioCapability5GSOk() (**os.File, bool) {
	if o == nil || o.BinaryDataUeRadioCapability5GS == nil {
		return nil, false
	}
	return o.BinaryDataUeRadioCapability5GS, true
}

// HasBinaryDataUeRadioCapability5GS returns a boolean if a field has been set.
func (o *InlineResponse200) HasBinaryDataUeRadioCapability5GS() bool {
	if o != nil && o.BinaryDataUeRadioCapability5GS != nil {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCapability5GS gets a reference to the given *os.File and assigns it to the BinaryDataUeRadioCapability5GS field.
func (o *InlineResponse200) SetBinaryDataUeRadioCapability5GS(v *os.File) {
	o.BinaryDataUeRadioCapability5GS = &v
}

// GetBinaryDataUeRadioCapabilityEPS returns the BinaryDataUeRadioCapabilityEPS field value if set, zero value otherwise.
func (o *InlineResponse200) GetBinaryDataUeRadioCapabilityEPS() *os.File {
	if o == nil || o.BinaryDataUeRadioCapabilityEPS == nil {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataUeRadioCapabilityEPS
}

// GetBinaryDataUeRadioCapabilityEPSOk returns a tuple with the BinaryDataUeRadioCapabilityEPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetBinaryDataUeRadioCapabilityEPSOk() (**os.File, bool) {
	if o == nil || o.BinaryDataUeRadioCapabilityEPS == nil {
		return nil, false
	}
	return o.BinaryDataUeRadioCapabilityEPS, true
}

// HasBinaryDataUeRadioCapabilityEPS returns a boolean if a field has been set.
func (o *InlineResponse200) HasBinaryDataUeRadioCapabilityEPS() bool {
	if o != nil && o.BinaryDataUeRadioCapabilityEPS != nil {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCapabilityEPS gets a reference to the given *os.File and assigns it to the BinaryDataUeRadioCapabilityEPS field.
func (o *InlineResponse200) SetBinaryDataUeRadioCapabilityEPS(v *os.File) {
	o.BinaryDataUeRadioCapabilityEPS = &v
}

func (o InlineResponse200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JsonData != nil {
		toSerialize["jsonData"] = o.JsonData
	}
	if o.BinaryDataUeRadioCapability5GS != nil {
		toSerialize["binaryDataUeRadioCapability5GS"] = o.BinaryDataUeRadioCapability5GS
	}
	if o.BinaryDataUeRadioCapabilityEPS != nil {
		toSerialize["binaryDataUeRadioCapabilityEPS"] = o.BinaryDataUeRadioCapabilityEPS
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200 struct {
	value *InlineResponse200
	isSet bool
}

func (v NullableInlineResponse200) Get() *InlineResponse200 {
	return v.value
}

func (v *NullableInlineResponse200) Set(val *InlineResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200(val *InlineResponse200) *NullableInlineResponse200 {
	return &NullableInlineResponse200{value: val, isSet: true}
}

func (v NullableInlineResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


