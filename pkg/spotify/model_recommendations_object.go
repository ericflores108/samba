/*
Spotify Web API

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spotify

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RecommendationsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendationsObject{}

// RecommendationsObject struct for RecommendationsObject
type RecommendationsObject struct {
	// An array of recommendation seed objects. 
	Seeds []RecommendationSeedObject `json:"seeds"`
	// An array of track object (simplified) ordered according to the parameters supplied. 
	Tracks []TrackObject `json:"tracks"`
}

type _RecommendationsObject RecommendationsObject

// NewRecommendationsObject instantiates a new RecommendationsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationsObject(seeds []RecommendationSeedObject, tracks []TrackObject) *RecommendationsObject {
	this := RecommendationsObject{}
	this.Seeds = seeds
	this.Tracks = tracks
	return &this
}

// NewRecommendationsObjectWithDefaults instantiates a new RecommendationsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationsObjectWithDefaults() *RecommendationsObject {
	this := RecommendationsObject{}
	return &this
}

// GetSeeds returns the Seeds field value
func (o *RecommendationsObject) GetSeeds() []RecommendationSeedObject {
	if o == nil {
		var ret []RecommendationSeedObject
		return ret
	}

	return o.Seeds
}

// GetSeedsOk returns a tuple with the Seeds field value
// and a boolean to check if the value has been set.
func (o *RecommendationsObject) GetSeedsOk() ([]RecommendationSeedObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Seeds, true
}

// SetSeeds sets field value
func (o *RecommendationsObject) SetSeeds(v []RecommendationSeedObject) {
	o.Seeds = v
}

// GetTracks returns the Tracks field value
func (o *RecommendationsObject) GetTracks() []TrackObject {
	if o == nil {
		var ret []TrackObject
		return ret
	}

	return o.Tracks
}

// GetTracksOk returns a tuple with the Tracks field value
// and a boolean to check if the value has been set.
func (o *RecommendationsObject) GetTracksOk() ([]TrackObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tracks, true
}

// SetTracks sets field value
func (o *RecommendationsObject) SetTracks(v []TrackObject) {
	o.Tracks = v
}

func (o RecommendationsObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["seeds"] = o.Seeds
	toSerialize["tracks"] = o.Tracks
	return toSerialize, nil
}

func (o *RecommendationsObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"seeds",
		"tracks",
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

	varRecommendationsObject := _RecommendationsObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecommendationsObject)

	if err != nil {
		return err
	}

	*o = RecommendationsObject(varRecommendationsObject)

	return err
}

type NullableRecommendationsObject struct {
	value *RecommendationsObject
	isSet bool
}

func (v NullableRecommendationsObject) Get() *RecommendationsObject {
	return v.value
}

func (v *NullableRecommendationsObject) Set(val *RecommendationsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationsObject(val *RecommendationsObject) *NullableRecommendationsObject {
	return &NullableRecommendationsObject{value: val, isSet: true}
}

func (v NullableRecommendationsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


