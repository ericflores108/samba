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

// checks if the GetAnAlbum401Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAnAlbum401Response{}

// GetAnAlbum401Response struct for GetAnAlbum401Response
type GetAnAlbum401Response struct {
	Error ErrorObject `json:"error"`
}

type _GetAnAlbum401Response GetAnAlbum401Response

// NewGetAnAlbum401Response instantiates a new GetAnAlbum401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAnAlbum401Response(error_ ErrorObject) *GetAnAlbum401Response {
	this := GetAnAlbum401Response{}
	this.Error = error_
	return &this
}

// NewGetAnAlbum401ResponseWithDefaults instantiates a new GetAnAlbum401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAnAlbum401ResponseWithDefaults() *GetAnAlbum401Response {
	this := GetAnAlbum401Response{}
	return &this
}

// GetError returns the Error field value
func (o *GetAnAlbum401Response) GetError() ErrorObject {
	if o == nil {
		var ret ErrorObject
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *GetAnAlbum401Response) GetErrorOk() (*ErrorObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *GetAnAlbum401Response) SetError(v ErrorObject) {
	o.Error = v
}

func (o GetAnAlbum401Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAnAlbum401Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *GetAnAlbum401Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
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

	varGetAnAlbum401Response := _GetAnAlbum401Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAnAlbum401Response)

	if err != nil {
		return err
	}

	*o = GetAnAlbum401Response(varGetAnAlbum401Response)

	return err
}

type NullableGetAnAlbum401Response struct {
	value *GetAnAlbum401Response
	isSet bool
}

func (v NullableGetAnAlbum401Response) Get() *GetAnAlbum401Response {
	return v.value
}

func (v *NullableGetAnAlbum401Response) Set(val *GetAnAlbum401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAnAlbum401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAnAlbum401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAnAlbum401Response(val *GetAnAlbum401Response) *NullableGetAnAlbum401Response {
	return &NullableGetAnAlbum401Response{value: val, isSet: true}
}

func (v NullableGetAnAlbum401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAnAlbum401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


