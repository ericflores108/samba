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

// checks if the PagingFeaturedPlaylistObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagingFeaturedPlaylistObject{}

// PagingFeaturedPlaylistObject struct for PagingFeaturedPlaylistObject
type PagingFeaturedPlaylistObject struct {
	// The localized message of a playlist. 
	Message *string `json:"message,omitempty"`
	Playlists *PagingPlaylistObject `json:"playlists,omitempty"`
}

// NewPagingFeaturedPlaylistObject instantiates a new PagingFeaturedPlaylistObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingFeaturedPlaylistObject() *PagingFeaturedPlaylistObject {
	this := PagingFeaturedPlaylistObject{}
	return &this
}

// NewPagingFeaturedPlaylistObjectWithDefaults instantiates a new PagingFeaturedPlaylistObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingFeaturedPlaylistObjectWithDefaults() *PagingFeaturedPlaylistObject {
	this := PagingFeaturedPlaylistObject{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PagingFeaturedPlaylistObject) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagingFeaturedPlaylistObject) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PagingFeaturedPlaylistObject) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PagingFeaturedPlaylistObject) SetMessage(v string) {
	o.Message = &v
}

// GetPlaylists returns the Playlists field value if set, zero value otherwise.
func (o *PagingFeaturedPlaylistObject) GetPlaylists() PagingPlaylistObject {
	if o == nil || IsNil(o.Playlists) {
		var ret PagingPlaylistObject
		return ret
	}
	return *o.Playlists
}

// GetPlaylistsOk returns a tuple with the Playlists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagingFeaturedPlaylistObject) GetPlaylistsOk() (*PagingPlaylistObject, bool) {
	if o == nil || IsNil(o.Playlists) {
		return nil, false
	}
	return o.Playlists, true
}

// HasPlaylists returns a boolean if a field has been set.
func (o *PagingFeaturedPlaylistObject) HasPlaylists() bool {
	if o != nil && !IsNil(o.Playlists) {
		return true
	}

	return false
}

// SetPlaylists gets a reference to the given PagingPlaylistObject and assigns it to the Playlists field.
func (o *PagingFeaturedPlaylistObject) SetPlaylists(v PagingPlaylistObject) {
	o.Playlists = &v
}

func (o PagingFeaturedPlaylistObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagingFeaturedPlaylistObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Playlists) {
		toSerialize["playlists"] = o.Playlists
	}
	return toSerialize, nil
}

type NullablePagingFeaturedPlaylistObject struct {
	value *PagingFeaturedPlaylistObject
	isSet bool
}

func (v NullablePagingFeaturedPlaylistObject) Get() *PagingFeaturedPlaylistObject {
	return v.value
}

func (v *NullablePagingFeaturedPlaylistObject) Set(val *PagingFeaturedPlaylistObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingFeaturedPlaylistObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingFeaturedPlaylistObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingFeaturedPlaylistObject(val *PagingFeaturedPlaylistObject) *NullablePagingFeaturedPlaylistObject {
	return &NullablePagingFeaturedPlaylistObject{value: val, isSet: true}
}

func (v NullablePagingFeaturedPlaylistObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingFeaturedPlaylistObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


