/*
Daytona

Daytona AI platform API Docs

API version: 1.0
Contact: support@daytona.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the KeyboardTypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyboardTypeResponse{}

// KeyboardTypeResponse struct for KeyboardTypeResponse
type KeyboardTypeResponse struct {
	Success bool   `json:"success"`
	Typed   string `json:"typed"`
}

type _KeyboardTypeResponse KeyboardTypeResponse

// NewKeyboardTypeResponse instantiates a new KeyboardTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyboardTypeResponse(success bool, typed string) *KeyboardTypeResponse {
	this := KeyboardTypeResponse{}
	this.Success = success
	this.Typed = typed
	return &this
}

// NewKeyboardTypeResponseWithDefaults instantiates a new KeyboardTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyboardTypeResponseWithDefaults() *KeyboardTypeResponse {
	this := KeyboardTypeResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *KeyboardTypeResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *KeyboardTypeResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *KeyboardTypeResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetTyped returns the Typed field value
func (o *KeyboardTypeResponse) GetTyped() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Typed
}

// GetTypedOk returns a tuple with the Typed field value
// and a boolean to check if the value has been set.
func (o *KeyboardTypeResponse) GetTypedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typed, true
}

// SetTyped sets field value
func (o *KeyboardTypeResponse) SetTyped(v string) {
	o.Typed = v
}

func (o KeyboardTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyboardTypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["typed"] = o.Typed
	return toSerialize, nil
}

func (o *KeyboardTypeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"typed",
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

	varKeyboardTypeResponse := _KeyboardTypeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKeyboardTypeResponse)

	if err != nil {
		return err
	}

	*o = KeyboardTypeResponse(varKeyboardTypeResponse)

	return err
}

type NullableKeyboardTypeResponse struct {
	value *KeyboardTypeResponse
	isSet bool
}

func (v NullableKeyboardTypeResponse) Get() *KeyboardTypeResponse {
	return v.value
}

func (v *NullableKeyboardTypeResponse) Set(val *KeyboardTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyboardTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyboardTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyboardTypeResponse(val *KeyboardTypeResponse) *NullableKeyboardTypeResponse {
	return &NullableKeyboardTypeResponse{value: val, isSet: true}
}

func (v NullableKeyboardTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyboardTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
