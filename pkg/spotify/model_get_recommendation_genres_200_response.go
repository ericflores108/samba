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

// checks if the GetRecommendationGenres200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecommendationGenres200Response{}

// GetRecommendationGenres200Response struct for GetRecommendationGenres200Response
type GetRecommendationGenres200Response struct {
	Genres []string `json:"genres"`
}

type _GetRecommendationGenres200Response GetRecommendationGenres200Response

// NewGetRecommendationGenres200Response instantiates a new GetRecommendationGenres200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecommendationGenres200Response(genres []string) *GetRecommendationGenres200Response {
	this := GetRecommendationGenres200Response{}
	this.Genres = genres
	return &this
}

// NewGetRecommendationGenres200ResponseWithDefaults instantiates a new GetRecommendationGenres200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecommendationGenres200ResponseWithDefaults() *GetRecommendationGenres200Response {
	this := GetRecommendationGenres200Response{}
	return &this
}

// GetGenres returns the Genres field value
func (o *GetRecommendationGenres200Response) GetGenres() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value
// and a boolean to check if the value has been set.
func (o *GetRecommendationGenres200Response) GetGenresOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Genres, true
}

// SetGenres sets field value
func (o *GetRecommendationGenres200Response) SetGenres(v []string) {
	o.Genres = v
}

func (o GetRecommendationGenres200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecommendationGenres200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["genres"] = o.Genres
	return toSerialize, nil
}

func (o *GetRecommendationGenres200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"genres",
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

	varGetRecommendationGenres200Response := _GetRecommendationGenres200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRecommendationGenres200Response)

	if err != nil {
		return err
	}

	*o = GetRecommendationGenres200Response(varGetRecommendationGenres200Response)

	return err
}

type NullableGetRecommendationGenres200Response struct {
	value *GetRecommendationGenres200Response
	isSet bool
}

func (v NullableGetRecommendationGenres200Response) Get() *GetRecommendationGenres200Response {
	return v.value
}

func (v *NullableGetRecommendationGenres200Response) Set(val *GetRecommendationGenres200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecommendationGenres200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecommendationGenres200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecommendationGenres200Response(val *GetRecommendationGenres200Response) *NullableGetRecommendationGenres200Response {
	return &NullableGetRecommendationGenres200Response{value: val, isSet: true}
}

func (v NullableGetRecommendationGenres200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecommendationGenres200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


