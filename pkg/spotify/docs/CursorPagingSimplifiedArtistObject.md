# CursorPagingSimplifiedArtistObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | A link to the Web API endpoint returning the full result of the request. | [optional] 
**Limit** | Pointer to **int64** | The maximum number of items in the response (as set in the query or by default). | [optional] 
**Next** | Pointer to **string** | URL to the next page of items. ( &#x60;null&#x60; if none) | [optional] 
**Cursors** | Pointer to [**CursorObject**](CursorObject.md) | The cursors used to find the next set of items. | [optional] 
**Total** | Pointer to **int64** | The total number of items available to return. | [optional] 
**Items** | Pointer to [**[]ArtistObject**](ArtistObject.md) |  | [optional] 

## Methods

### NewCursorPagingSimplifiedArtistObject

`func NewCursorPagingSimplifiedArtistObject() *CursorPagingSimplifiedArtistObject`

NewCursorPagingSimplifiedArtistObject instantiates a new CursorPagingSimplifiedArtistObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPagingSimplifiedArtistObjectWithDefaults

`func NewCursorPagingSimplifiedArtistObjectWithDefaults() *CursorPagingSimplifiedArtistObject`

NewCursorPagingSimplifiedArtistObjectWithDefaults instantiates a new CursorPagingSimplifiedArtistObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CursorPagingSimplifiedArtistObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CursorPagingSimplifiedArtistObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CursorPagingSimplifiedArtistObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CursorPagingSimplifiedArtistObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLimit

`func (o *CursorPagingSimplifiedArtistObject) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CursorPagingSimplifiedArtistObject) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CursorPagingSimplifiedArtistObject) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CursorPagingSimplifiedArtistObject) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *CursorPagingSimplifiedArtistObject) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CursorPagingSimplifiedArtistObject) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CursorPagingSimplifiedArtistObject) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *CursorPagingSimplifiedArtistObject) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCursors

`func (o *CursorPagingSimplifiedArtistObject) GetCursors() CursorObject`

GetCursors returns the Cursors field if non-nil, zero value otherwise.

### GetCursorsOk

`func (o *CursorPagingSimplifiedArtistObject) GetCursorsOk() (*CursorObject, bool)`

GetCursorsOk returns a tuple with the Cursors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursors

`func (o *CursorPagingSimplifiedArtistObject) SetCursors(v CursorObject)`

SetCursors sets Cursors field to given value.

### HasCursors

`func (o *CursorPagingSimplifiedArtistObject) HasCursors() bool`

HasCursors returns a boolean if a field has been set.

### GetTotal

`func (o *CursorPagingSimplifiedArtistObject) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CursorPagingSimplifiedArtistObject) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CursorPagingSimplifiedArtistObject) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CursorPagingSimplifiedArtistObject) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *CursorPagingSimplifiedArtistObject) GetItems() []ArtistObject`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CursorPagingSimplifiedArtistObject) GetItemsOk() (*[]ArtistObject, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CursorPagingSimplifiedArtistObject) SetItems(v []ArtistObject)`

SetItems sets Items field to given value.

### HasItems

`func (o *CursorPagingSimplifiedArtistObject) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


