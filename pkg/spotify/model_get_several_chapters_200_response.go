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

// checks if the GetSeveralChapters200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSeveralChapters200Response{}

// GetSeveralChapters200Response struct for GetSeveralChapters200Response
type GetSeveralChapters200Response struct {
	Chapters []ChapterObject `json:"chapters"`
}

type _GetSeveralChapters200Response GetSeveralChapters200Response

// NewGetSeveralChapters200Response instantiates a new GetSeveralChapters200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSeveralChapters200Response(chapters []ChapterObject) *GetSeveralChapters200Response {
	this := GetSeveralChapters200Response{}
	this.Chapters = chapters
	return &this
}

// NewGetSeveralChapters200ResponseWithDefaults instantiates a new GetSeveralChapters200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSeveralChapters200ResponseWithDefaults() *GetSeveralChapters200Response {
	this := GetSeveralChapters200Response{}
	return &this
}

// GetChapters returns the Chapters field value
func (o *GetSeveralChapters200Response) GetChapters() []ChapterObject {
	if o == nil {
		var ret []ChapterObject
		return ret
	}

	return o.Chapters
}

// GetChaptersOk returns a tuple with the Chapters field value
// and a boolean to check if the value has been set.
func (o *GetSeveralChapters200Response) GetChaptersOk() ([]ChapterObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Chapters, true
}

// SetChapters sets field value
func (o *GetSeveralChapters200Response) SetChapters(v []ChapterObject) {
	o.Chapters = v
}

func (o GetSeveralChapters200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSeveralChapters200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chapters"] = o.Chapters
	return toSerialize, nil
}

func (o *GetSeveralChapters200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chapters",
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

	varGetSeveralChapters200Response := _GetSeveralChapters200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSeveralChapters200Response)

	if err != nil {
		return err
	}

	*o = GetSeveralChapters200Response(varGetSeveralChapters200Response)

	return err
}

type NullableGetSeveralChapters200Response struct {
	value *GetSeveralChapters200Response
	isSet bool
}

func (v NullableGetSeveralChapters200Response) Get() *GetSeveralChapters200Response {
	return v.value
}

func (v *NullableGetSeveralChapters200Response) Set(val *GetSeveralChapters200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSeveralChapters200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSeveralChapters200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSeveralChapters200Response(val *GetSeveralChapters200Response) *NullableGetSeveralChapters200Response {
	return &NullableGetSeveralChapters200Response{value: val, isSet: true}
}

func (v NullableGetSeveralChapters200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSeveralChapters200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


