/*
Daytona

Daytona AI platform API Docs

API version: 1.0
Contact: support@daytona.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daytonaapiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SetImageGeneralStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetImageGeneralStatus{}

// SetImageGeneralStatus struct for SetImageGeneralStatus
type SetImageGeneralStatus struct {
	// Whether the image is general
	General bool `json:"general"`
}

type _SetImageGeneralStatus SetImageGeneralStatus

// NewSetImageGeneralStatus instantiates a new SetImageGeneralStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetImageGeneralStatus(general bool) *SetImageGeneralStatus {
	this := SetImageGeneralStatus{}
	this.General = general
	return &this
}

// NewSetImageGeneralStatusWithDefaults instantiates a new SetImageGeneralStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetImageGeneralStatusWithDefaults() *SetImageGeneralStatus {
	this := SetImageGeneralStatus{}
	return &this
}

// GetGeneral returns the General field value
func (o *SetImageGeneralStatus) GetGeneral() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.General
}

// GetGeneralOk returns a tuple with the General field value
// and a boolean to check if the value has been set.
func (o *SetImageGeneralStatus) GetGeneralOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.General, true
}

// SetGeneral sets field value
func (o *SetImageGeneralStatus) SetGeneral(v bool) {
	o.General = v
}

func (o SetImageGeneralStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetImageGeneralStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["general"] = o.General
	return toSerialize, nil
}

func (o *SetImageGeneralStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"general",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSetImageGeneralStatus := _SetImageGeneralStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetImageGeneralStatus)

	if err != nil {
		return err
	}

	*o = SetImageGeneralStatus(varSetImageGeneralStatus)

	return err
}

type NullableSetImageGeneralStatus struct {
	value *SetImageGeneralStatus
	isSet bool
}

func (v NullableSetImageGeneralStatus) Get() *SetImageGeneralStatus {
	return v.value
}

func (v *NullableSetImageGeneralStatus) Set(val *SetImageGeneralStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSetImageGeneralStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSetImageGeneralStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetImageGeneralStatus(val *SetImageGeneralStatus) *NullableSetImageGeneralStatus {
	return &NullableSetImageGeneralStatus{value: val, isSet: true}
}

func (v NullableSetImageGeneralStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetImageGeneralStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
