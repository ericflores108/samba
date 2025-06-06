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

// checks if the ExternalUrlObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalUrlObject{}

// ExternalUrlObject struct for ExternalUrlObject
type ExternalUrlObject struct {
	// The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object. 
	Spotify *string `json:"spotify,omitempty"`
}

// NewExternalUrlObject instantiates a new ExternalUrlObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalUrlObject() *ExternalUrlObject {
	this := ExternalUrlObject{}
	return &this
}

// NewExternalUrlObjectWithDefaults instantiates a new ExternalUrlObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalUrlObjectWithDefaults() *ExternalUrlObject {
	this := ExternalUrlObject{}
	return &this
}

// GetSpotify returns the Spotify field value if set, zero value otherwise.
func (o *ExternalUrlObject) GetSpotify() string {
	if o == nil || IsNil(o.Spotify) {
		var ret string
		return ret
	}
	return *o.Spotify
}

// GetSpotifyOk returns a tuple with the Spotify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalUrlObject) GetSpotifyOk() (*string, bool) {
	if o == nil || IsNil(o.Spotify) {
		return nil, false
	}
	return o.Spotify, true
}

// HasSpotify returns a boolean if a field has been set.
func (o *ExternalUrlObject) HasSpotify() bool {
	if o != nil && !IsNil(o.Spotify) {
		return true
	}

	return false
}

// SetSpotify gets a reference to the given string and assigns it to the Spotify field.
func (o *ExternalUrlObject) SetSpotify(v string) {
	o.Spotify = &v
}

func (o ExternalUrlObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalUrlObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Spotify) {
		toSerialize["spotify"] = o.Spotify
	}
	return toSerialize, nil
}

type NullableExternalUrlObject struct {
	value *ExternalUrlObject
	isSet bool
}

func (v NullableExternalUrlObject) Get() *ExternalUrlObject {
	return v.value
}

func (v *NullableExternalUrlObject) Set(val *ExternalUrlObject) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalUrlObject) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalUrlObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalUrlObject(val *ExternalUrlObject) *NullableExternalUrlObject {
	return &NullableExternalUrlObject{value: val, isSet: true}
}

func (v NullableExternalUrlObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalUrlObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


