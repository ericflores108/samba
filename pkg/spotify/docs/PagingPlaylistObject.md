# PagingPlaylistObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | A link to the Web API endpoint returning the full result of the request  | 
**Limit** | **int64** | The maximum number of items in the response (as set in the query or by default).  | 
**Next** | **NullableString** | URL to the next page of items. ( &#x60;null&#x60; if none)  | 
**Offset** | **int64** | The offset of the items returned (as set in the query or by default)  | 
**Previous** | **NullableString** | URL to the previous page of items. ( &#x60;null&#x60; if none)  | 
**Total** | **int64** | The total number of items available to return.  | 
**Items** | [**[]SimplifiedPlaylistObject**](SimplifiedPlaylistObject.md) |  | 

## Methods

### NewPagingPlaylistObject

`func NewPagingPlaylistObject(href string, limit int64, next NullableString, offset int64, previous NullableString, total int64, items []SimplifiedPlaylistObject, ) *PagingPlaylistObject`

NewPagingPlaylistObject instantiates a new PagingPlaylistObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingPlaylistObjectWithDefaults

`func NewPagingPlaylistObjectWithDefaults() *PagingPlaylistObject`

NewPagingPlaylistObjectWithDefaults instantiates a new PagingPlaylistObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PagingPlaylistObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PagingPlaylistObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PagingPlaylistObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetLimit

`func (o *PagingPlaylistObject) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagingPlaylistObject) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagingPlaylistObject) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *PagingPlaylistObject) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PagingPlaylistObject) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PagingPlaylistObject) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *PagingPlaylistObject) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PagingPlaylistObject) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetOffset

`func (o *PagingPlaylistObject) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PagingPlaylistObject) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PagingPlaylistObject) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetPrevious

`func (o *PagingPlaylistObject) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PagingPlaylistObject) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PagingPlaylistObject) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *PagingPlaylistObject) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PagingPlaylistObject) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetTotal

`func (o *PagingPlaylistObject) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PagingPlaylistObject) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PagingPlaylistObject) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetItems

`func (o *PagingPlaylistObject) GetItems() []SimplifiedPlaylistObject`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PagingPlaylistObject) GetItemsOk() (*[]SimplifiedPlaylistObject, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PagingPlaylistObject) SetItems(v []SimplifiedPlaylistObject)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


