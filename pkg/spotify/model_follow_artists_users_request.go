/*
Spotify Web API

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spotify

import (
	"encoding/json"
	"fmt"
)

// checks if the FollowArtistsUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowArtistsUsersRequest{}

// FollowArtistsUsersRequest struct for FollowArtistsUsersRequest
type FollowArtistsUsersRequest struct {
	// A JSON array of the artist or user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `{ids:[\"74ASZWbe4lXaubB36ztrGX\", \"08td7MxkoHQkXnWAYD8d6Q\"]}`. A maximum of 50 IDs can be sent in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._ 
	Ids []string `json:"ids"`
	AdditionalProperties map[string]interface{}
}

type _FollowArtistsUsersRequest FollowArtistsUsersRequest

// NewFollowArtistsUsersRequest instantiates a new FollowArtistsUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowArtistsUsersRequest(ids []string) *FollowArtistsUsersRequest {
	this := FollowArtistsUsersRequest{}
	this.Ids = ids
	return &this
}

// NewFollowArtistsUsersRequestWithDefaults instantiates a new FollowArtistsUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowArtistsUsersRequestWithDefaults() *FollowArtistsUsersRequest {
	this := FollowArtistsUsersRequest{}
	return &this
}

// GetIds returns the Ids field value
func (o *FollowArtistsUsersRequest) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *FollowArtistsUsersRequest) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *FollowArtistsUsersRequest) SetIds(v []string) {
	o.Ids = v
}

func (o FollowArtistsUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowArtistsUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FollowArtistsUsersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ids",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFollowArtistsUsersRequest := _FollowArtistsUsersRequest{}

	err = json.Unmarshal(data, &varFollowArtistsUsersRequest)

	if err != nil {
		return err
	}

	*o = FollowArtistsUsersRequest(varFollowArtistsUsersRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFollowArtistsUsersRequest struct {
	value *FollowArtistsUsersRequest
	isSet bool
}

func (v NullableFollowArtistsUsersRequest) Get() *FollowArtistsUsersRequest {
	return v.value
}

func (v *NullableFollowArtistsUsersRequest) Set(val *FollowArtistsUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowArtistsUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowArtistsUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowArtistsUsersRequest(val *FollowArtistsUsersRequest) *NullableFollowArtistsUsersRequest {
	return &NullableFollowArtistsUsersRequest{value: val, isSet: true}
}

func (v NullableFollowArtistsUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowArtistsUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


