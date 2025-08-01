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

// checks if the ScreenshotResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScreenshotResponse{}

// ScreenshotResponse struct for ScreenshotResponse
type ScreenshotResponse struct {
	// Base64 encoded screenshot image data
	Screenshot string `json:"screenshot"`
	// The current cursor position when the screenshot was taken
	CursorPosition map[string]interface{} `json:"cursorPosition,omitempty"`
	// The size of the screenshot data in bytes
	SizeBytes *float32 `json:"sizeBytes,omitempty"`
}

type _ScreenshotResponse ScreenshotResponse

// NewScreenshotResponse instantiates a new ScreenshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreenshotResponse(screenshot string) *ScreenshotResponse {
	this := ScreenshotResponse{}
	this.Screenshot = screenshot
	return &this
}

// NewScreenshotResponseWithDefaults instantiates a new ScreenshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreenshotResponseWithDefaults() *ScreenshotResponse {
	this := ScreenshotResponse{}
	return &this
}

// GetScreenshot returns the Screenshot field value
func (o *ScreenshotResponse) GetScreenshot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Screenshot
}

// GetScreenshotOk returns a tuple with the Screenshot field value
// and a boolean to check if the value has been set.
func (o *ScreenshotResponse) GetScreenshotOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Screenshot, true
}

// SetScreenshot sets field value
func (o *ScreenshotResponse) SetScreenshot(v string) {
	o.Screenshot = v
}

// GetCursorPosition returns the CursorPosition field value if set, zero value otherwise.
func (o *ScreenshotResponse) GetCursorPosition() map[string]interface{} {
	if o == nil || IsNil(o.CursorPosition) {
		var ret map[string]interface{}
		return ret
	}
	return o.CursorPosition
}

// GetCursorPositionOk returns a tuple with the CursorPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenshotResponse) GetCursorPositionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CursorPosition) {
		return map[string]interface{}{}, false
	}
	return o.CursorPosition, true
}

// HasCursorPosition returns a boolean if a field has been set.
func (o *ScreenshotResponse) HasCursorPosition() bool {
	if o != nil && !IsNil(o.CursorPosition) {
		return true
	}

	return false
}

// SetCursorPosition gets a reference to the given map[string]interface{} and assigns it to the CursorPosition field.
func (o *ScreenshotResponse) SetCursorPosition(v map[string]interface{}) {
	o.CursorPosition = v
}

// GetSizeBytes returns the SizeBytes field value if set, zero value otherwise.
func (o *ScreenshotResponse) GetSizeBytes() float32 {
	if o == nil || IsNil(o.SizeBytes) {
		var ret float32
		return ret
	}
	return *o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenshotResponse) GetSizeBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.SizeBytes) {
		return nil, false
	}
	return o.SizeBytes, true
}

// HasSizeBytes returns a boolean if a field has been set.
func (o *ScreenshotResponse) HasSizeBytes() bool {
	if o != nil && !IsNil(o.SizeBytes) {
		return true
	}

	return false
}

// SetSizeBytes gets a reference to the given float32 and assigns it to the SizeBytes field.
func (o *ScreenshotResponse) SetSizeBytes(v float32) {
	o.SizeBytes = &v
}

func (o ScreenshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScreenshotResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["screenshot"] = o.Screenshot
	if !IsNil(o.CursorPosition) {
		toSerialize["cursorPosition"] = o.CursorPosition
	}
	if !IsNil(o.SizeBytes) {
		toSerialize["sizeBytes"] = o.SizeBytes
	}
	return toSerialize, nil
}

func (o *ScreenshotResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"screenshot",
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

	varScreenshotResponse := _ScreenshotResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScreenshotResponse)

	if err != nil {
		return err
	}

	*o = ScreenshotResponse(varScreenshotResponse)

	return err
}

type NullableScreenshotResponse struct {
	value *ScreenshotResponse
	isSet bool
}

func (v NullableScreenshotResponse) Get() *ScreenshotResponse {
	return v.value
}

func (v *NullableScreenshotResponse) Set(val *ScreenshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScreenshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScreenshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreenshotResponse(val *ScreenshotResponse) *NullableScreenshotResponse {
	return &NullableScreenshotResponse{value: val, isSet: true}
}

func (v NullableScreenshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreenshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
