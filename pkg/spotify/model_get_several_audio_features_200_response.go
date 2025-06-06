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

// checks if the GetSeveralAudioFeatures200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSeveralAudioFeatures200Response{}

// GetSeveralAudioFeatures200Response struct for GetSeveralAudioFeatures200Response
type GetSeveralAudioFeatures200Response struct {
	AudioFeatures []AudioFeaturesObject `json:"audio_features"`
}

type _GetSeveralAudioFeatures200Response GetSeveralAudioFeatures200Response

// NewGetSeveralAudioFeatures200Response instantiates a new GetSeveralAudioFeatures200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSeveralAudioFeatures200Response(audioFeatures []AudioFeaturesObject) *GetSeveralAudioFeatures200Response {
	this := GetSeveralAudioFeatures200Response{}
	this.AudioFeatures = audioFeatures
	return &this
}

// NewGetSeveralAudioFeatures200ResponseWithDefaults instantiates a new GetSeveralAudioFeatures200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSeveralAudioFeatures200ResponseWithDefaults() *GetSeveralAudioFeatures200Response {
	this := GetSeveralAudioFeatures200Response{}
	return &this
}

// GetAudioFeatures returns the AudioFeatures field value
func (o *GetSeveralAudioFeatures200Response) GetAudioFeatures() []AudioFeaturesObject {
	if o == nil {
		var ret []AudioFeaturesObject
		return ret
	}

	return o.AudioFeatures
}

// GetAudioFeaturesOk returns a tuple with the AudioFeatures field value
// and a boolean to check if the value has been set.
func (o *GetSeveralAudioFeatures200Response) GetAudioFeaturesOk() ([]AudioFeaturesObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioFeatures, true
}

// SetAudioFeatures sets field value
func (o *GetSeveralAudioFeatures200Response) SetAudioFeatures(v []AudioFeaturesObject) {
	o.AudioFeatures = v
}

func (o GetSeveralAudioFeatures200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSeveralAudioFeatures200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audio_features"] = o.AudioFeatures
	return toSerialize, nil
}

func (o *GetSeveralAudioFeatures200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"audio_features",
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

	varGetSeveralAudioFeatures200Response := _GetSeveralAudioFeatures200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSeveralAudioFeatures200Response)

	if err != nil {
		return err
	}

	*o = GetSeveralAudioFeatures200Response(varGetSeveralAudioFeatures200Response)

	return err
}

type NullableGetSeveralAudioFeatures200Response struct {
	value *GetSeveralAudioFeatures200Response
	isSet bool
}

func (v NullableGetSeveralAudioFeatures200Response) Get() *GetSeveralAudioFeatures200Response {
	return v.value
}

func (v *NullableGetSeveralAudioFeatures200Response) Set(val *GetSeveralAudioFeatures200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSeveralAudioFeatures200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSeveralAudioFeatures200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSeveralAudioFeatures200Response(val *GetSeveralAudioFeatures200Response) *NullableGetSeveralAudioFeatures200Response {
	return &NullableGetSeveralAudioFeatures200Response{value: val, isSet: true}
}

func (v NullableGetSeveralAudioFeatures200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSeveralAudioFeatures200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


