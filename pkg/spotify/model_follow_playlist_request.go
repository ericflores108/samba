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

// checks if the FollowPlaylistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowPlaylistRequest{}

// FollowPlaylistRequest struct for FollowPlaylistRequest
type FollowPlaylistRequest struct {
	// Defaults to `true`. If `true` the playlist will be included in user's public playlists (added to profile), if `false` it will remain private. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists) 
	Public *bool `json:"public,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FollowPlaylistRequest FollowPlaylistRequest

// NewFollowPlaylistRequest instantiates a new FollowPlaylistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowPlaylistRequest() *FollowPlaylistRequest {
	this := FollowPlaylistRequest{}
	return &this
}

// NewFollowPlaylistRequestWithDefaults instantiates a new FollowPlaylistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowPlaylistRequestWithDefaults() *FollowPlaylistRequest {
	this := FollowPlaylistRequest{}
	return &this
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *FollowPlaylistRequest) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowPlaylistRequest) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *FollowPlaylistRequest) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *FollowPlaylistRequest) SetPublic(v bool) {
	o.Public = &v
}

func (o FollowPlaylistRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowPlaylistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FollowPlaylistRequest) UnmarshalJSON(data []byte) (err error) {
	varFollowPlaylistRequest := _FollowPlaylistRequest{}

	err = json.Unmarshal(data, &varFollowPlaylistRequest)

	if err != nil {
		return err
	}

	*o = FollowPlaylistRequest(varFollowPlaylistRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "public")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFollowPlaylistRequest struct {
	value *FollowPlaylistRequest
	isSet bool
}

func (v NullableFollowPlaylistRequest) Get() *FollowPlaylistRequest {
	return v.value
}

func (v *NullableFollowPlaylistRequest) Set(val *FollowPlaylistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowPlaylistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowPlaylistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowPlaylistRequest(val *FollowPlaylistRequest) *NullableFollowPlaylistRequest {
	return &NullableFollowPlaylistRequest{value: val, isSet: true}
}

func (v NullableFollowPlaylistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowPlaylistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


