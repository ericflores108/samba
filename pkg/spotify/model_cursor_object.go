/*
Spotify Web API

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spotify

import (
	"encoding/json"
)

// checks if the CursorObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CursorObject{}

// CursorObject struct for CursorObject
type CursorObject struct {
	// The cursor to use as key to find the next page of items.
	After *string `json:"after,omitempty"`
	// The cursor to use as key to find the previous page of items.
	Before *string `json:"before,omitempty"`
}

// NewCursorObject instantiates a new CursorObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCursorObject() *CursorObject {
	this := CursorObject{}
	return &this
}

// NewCursorObjectWithDefaults instantiates a new CursorObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCursorObjectWithDefaults() *CursorObject {
	this := CursorObject{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *CursorObject) GetAfter() string {
	if o == nil || IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CursorObject) GetAfterOk() (*string, bool) {
	if o == nil || IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *CursorObject) HasAfter() bool {
	if o != nil && !IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *CursorObject) SetAfter(v string) {
	o.After = &v
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *CursorObject) GetBefore() string {
	if o == nil || IsNil(o.Before) {
		var ret string
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CursorObject) GetBeforeOk() (*string, bool) {
	if o == nil || IsNil(o.Before) {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *CursorObject) HasBefore() bool {
	if o != nil && !IsNil(o.Before) {
		return true
	}

	return false
}

// SetBefore gets a reference to the given string and assigns it to the Before field.
func (o *CursorObject) SetBefore(v string) {
	o.Before = &v
}

func (o CursorObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CursorObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.After) {
		toSerialize["after"] = o.After
	}
	if !IsNil(o.Before) {
		toSerialize["before"] = o.Before
	}
	return toSerialize, nil
}

type NullableCursorObject struct {
	value *CursorObject
	isSet bool
}

func (v NullableCursorObject) Get() *CursorObject {
	return v.value
}

func (v *NullableCursorObject) Set(val *CursorObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCursorObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCursorObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCursorObject(val *CursorObject) *NullableCursorObject {
	return &NullableCursorObject{value: val, isSet: true}
}

func (v NullableCursorObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCursorObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


