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

// checks if the PlaylistTracksRefObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaylistTracksRefObject{}

// PlaylistTracksRefObject struct for PlaylistTracksRefObject
type PlaylistTracksRefObject struct {
	// A link to the Web API endpoint where full details of the playlist's tracks can be retrieved. 
	Href *string `json:"href,omitempty"`
	// Number of tracks in the playlist. 
	Total *int64 `json:"total,omitempty"`
}

// NewPlaylistTracksRefObject instantiates a new PlaylistTracksRefObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaylistTracksRefObject() *PlaylistTracksRefObject {
	this := PlaylistTracksRefObject{}
	return &this
}

// NewPlaylistTracksRefObjectWithDefaults instantiates a new PlaylistTracksRefObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaylistTracksRefObjectWithDefaults() *PlaylistTracksRefObject {
	this := PlaylistTracksRefObject{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PlaylistTracksRefObject) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistTracksRefObject) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PlaylistTracksRefObject) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PlaylistTracksRefObject) SetHref(v string) {
	o.Href = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PlaylistTracksRefObject) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistTracksRefObject) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PlaylistTracksRefObject) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *PlaylistTracksRefObject) SetTotal(v int64) {
	o.Total = &v
}

func (o PlaylistTracksRefObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaylistTracksRefObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullablePlaylistTracksRefObject struct {
	value *PlaylistTracksRefObject
	isSet bool
}

func (v NullablePlaylistTracksRefObject) Get() *PlaylistTracksRefObject {
	return v.value
}

func (v *NullablePlaylistTracksRefObject) Set(val *PlaylistTracksRefObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaylistTracksRefObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaylistTracksRefObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaylistTracksRefObject(val *PlaylistTracksRefObject) *NullablePlaylistTracksRefObject {
	return &NullablePlaylistTracksRefObject{value: val, isSet: true}
}

func (v NullablePlaylistTracksRefObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaylistTracksRefObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


