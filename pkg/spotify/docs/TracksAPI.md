# \TracksAPI

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTracksToPlaylist**](TracksAPI.md#AddTracksToPlaylist) | **Post** /playlists/{playlist_id}/tracks | Add Items to Playlist 
[**CheckUsersSavedTracks**](TracksAPI.md#CheckUsersSavedTracks) | **Get** /me/tracks/contains | Check User&#39;s Saved Tracks 
[**GetAnAlbumsTracks**](TracksAPI.md#GetAnAlbumsTracks) | **Get** /albums/{id}/tracks | Get Album Tracks 
[**GetAnArtistsTopTracks**](TracksAPI.md#GetAnArtistsTopTracks) | **Get** /artists/{id}/top-tracks | Get Artist&#39;s Top Tracks 
[**GetAudioAnalysis**](TracksAPI.md#GetAudioAnalysis) | **Get** /audio-analysis/{id} | Get Track&#39;s Audio Analysis 
[**GetAudioFeatures**](TracksAPI.md#GetAudioFeatures) | **Get** /audio-features/{id} | Get Track&#39;s Audio Features 
[**GetPlaylistsTracks**](TracksAPI.md#GetPlaylistsTracks) | **Get** /playlists/{playlist_id}/tracks | Get Playlist Items 
[**GetRecommendations**](TracksAPI.md#GetRecommendations) | **Get** /recommendations | Get Recommendations 
[**GetSeveralAudioFeatures**](TracksAPI.md#GetSeveralAudioFeatures) | **Get** /audio-features | Get Several Tracks&#39; Audio Features 
[**GetSeveralTracks**](TracksAPI.md#GetSeveralTracks) | **Get** /tracks | Get Several Tracks 
[**GetTrack**](TracksAPI.md#GetTrack) | **Get** /tracks/{id} | Get Track 
[**GetUsersSavedTracks**](TracksAPI.md#GetUsersSavedTracks) | **Get** /me/tracks | Get User&#39;s Saved Tracks 
[**GetUsersTopArtistsAndTracks**](TracksAPI.md#GetUsersTopArtistsAndTracks) | **Get** /me/top/{type} | Get User&#39;s Top Items 
[**RemoveTracksPlaylist**](TracksAPI.md#RemoveTracksPlaylist) | **Delete** /playlists/{playlist_id}/tracks | Remove Playlist Items 
[**RemoveTracksUser**](TracksAPI.md#RemoveTracksUser) | **Delete** /me/tracks | Remove User&#39;s Saved Tracks 
[**ReorderOrReplacePlaylistsTracks**](TracksAPI.md#ReorderOrReplacePlaylistsTracks) | **Put** /playlists/{playlist_id}/tracks | Update Playlist Items 
[**SaveTracksUser**](TracksAPI.md#SaveTracksUser) | **Put** /me/tracks | Save Tracks for Current User 



## AddTracksToPlaylist

> ReorderOrReplacePlaylistsTracks200Response AddTracksToPlaylist(ctx, playlistId).Position(position).Uris(uris).AddTracksToPlaylistRequest(addTracksToPlaylistRequest).Execute()

Add Items to Playlist 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	playlistId := "3cEYpjA9oz9GiPac4AsH4n" // string | 
	position := int64(0) // int64 |  (optional)
	uris := "spotify:track:4iV5W9uYEdYUVa79Axb7Rh,spotify:track:1301WleyT98MSxVHPZCA6M" // string |  (optional)
	addTracksToPlaylistRequest := *openapiclient.NewAddTracksToPlaylistRequest() // AddTracksToPlaylistRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.AddTracksToPlaylist(context.Background(), playlistId).Position(position).Uris(uris).AddTracksToPlaylistRequest(addTracksToPlaylistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.AddTracksToPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTracksToPlaylist`: ReorderOrReplacePlaylistsTracks200Response
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.AddTracksToPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTracksToPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **position** | **int64** |  | 
 **uris** | **string** |  | 
 **addTracksToPlaylistRequest** | [**AddTracksToPlaylistRequest**](AddTracksToPlaylistRequest.md) |  | 

### Return type

[**ReorderOrReplacePlaylistsTracks200Response**](ReorderOrReplacePlaylistsTracks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckUsersSavedTracks

> []bool CheckUsersSavedTracks(ctx).Ids(ids).Execute()

Check User's Saved Tracks 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	ids := "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.CheckUsersSavedTracks(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.CheckUsersSavedTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUsersSavedTracks`: []bool
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.CheckUsersSavedTracks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUsersSavedTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 

### Return type

**[]bool**

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnAlbumsTracks

> PagingSimplifiedTrackObject GetAnAlbumsTracks(ctx, id).Market(market).Limit(limit).Offset(offset).Execute()

Get Album Tracks 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	id := "4aawyAB9vmqN3uQ7FjRGTy" // string | 
	market := "ES" // string |  (optional)
	limit := int64(10) // int64 |  (optional) (default to 20)
	offset := int64(5) // int64 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.GetAnAlbumsTracks(context.Background(), id).Market(market).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.GetAnAlbumsTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnAlbumsTracks`: PagingSimplifiedTrackObject
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.GetAnAlbumsTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnAlbumsTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 
 **limit** | **int64** |  | [default to 20]
 **offset** | **int64** |  | [default to 0]

### Return type

[**PagingSimplifiedTrackObject**](PagingSimplifiedTrackObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnArtistsTopTracks

> GetAnArtistsTopTracks200Response GetAnArtistsTopTracks(ctx, id).Market(market).Execute()

Get Artist's Top Tracks 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	id := "0TnOYISbd1XYRBk9myaseg" // string | 
	market := "ES" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.GetAnArtistsTopTracks(context.Background(), id).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.GetAnArtistsTopTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnArtistsTopTracks`: GetAnArtistsTopTracks200Response
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.GetAnArtistsTopTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnArtistsTopTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 

### Return type

[**GetAnArtistsTopTracks200Response**](GetAnArtistsTopTracks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudioAnalysis

> AudioAnalysisObject GetAudioAnalysis(ctx, id).Execute()

Get Track's Audio Analysis 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	id := "11dFghVXANMlKmJXsNCbNl" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.GetAudioAnalysis(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.GetAudioAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudioAnalysis`: AudioAnalysisObject
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.GetAudioAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AudioAnalysisObject**](AudioAnalysisObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudioFeatures

> AudioFeaturesObject GetAudioFeatures(ctx, id).Execute()

Get Track's Audio Features 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	id := "11dFghVXANMlKmJXsNCbNl" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.GetAudioFeatures(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.GetAudioFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudioFeatures`: AudioFeaturesObject
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.GetAudioFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AudioFeaturesObject**](AudioFeaturesObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylistsTracks

> PagingPlaylistTrackObject GetPlaylistsTracks(ctx, playlistId).Market(market).Fields(fields).Limit(limit).Offset(offset).AdditionalTypes(additionalTypes).Execute()

Get Playlist Items 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	playlistId := "3cEYpjA9oz9GiPac4AsH4n" // string | 
	market := "ES" // string |  (optional)
	fields := "items(added_by.id,track(name,href,album(name,href)))" // string |  (optional)
	limit := int64(10) // int64 |  (optional) (default to 20)
	offset := int64(5) // int64 |  (optional) (default to 0)
	additionalTypes := "additionalTypes_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.GetPlaylistsTracks(context.Background(), playlistId).Market(market).Fields(fields).Limit(limit).Offset(offset).AdditionalTypes(additionalTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.GetPlaylistsTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylistsTracks`: PagingPlaylistTrackObject
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.GetPlaylistsTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistsTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 
 **fields** | **string** |  | 
 **limit** | **int64** |  | [default to 20]
 **offset** | **int64** |  | [default to 0]
 **additionalTypes** | **string** |  | 

### Return type

[**PagingPlaylistTrackObject**](PagingPlaylistTrackObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendations

> RecommendationsObject GetRecommendations(ctx).SeedArtists(seedArtists).SeedGenres(seedGenres).SeedTracks(seedTracks).Limit(limit).Market(market).MinAcousticness(minAcousticness).MaxAcousticness(maxAcousticness).TargetAcousticness(targetAcousticness).MinDanceability(minDanceability).MaxDanceability(maxDanceability).TargetDanceability(targetDanceability).MinDurationMs(minDurationMs).MaxDurationMs(maxDurationMs).TargetDurationMs(targetDurationMs).MinEnergy(minEnergy).MaxEnergy(maxEnergy).TargetEnergy(targetEnergy).MinInstrumentalness(minInstrumentalness).MaxInstrumentalness(maxInstrumentalness).TargetInstrumentalness(targetInstrumentalness).MinKey(minKey).MaxKey(maxKey).TargetKey(targetKey).MinLiveness(minLiveness).MaxLiveness(maxLiveness).TargetLiveness(targetLiveness).MinLoudness(minLoudness).MaxLoudness(maxLoudness).TargetLoudness(targetLoudness).MinMode(minMode).MaxMode(maxMode).TargetMode(targetMode).MinPopularity(minPopularity).MaxPopularity(maxPopularity).TargetPopularity(targetPopularity).MinSpeechiness(minSpeechiness).MaxSpeechiness(maxSpeechiness).TargetSpeechiness(targetSpeechiness).MinTempo(minTempo).MaxTempo(maxTempo).TargetTempo(targetTempo).MinTimeSignature(minTimeSignature).MaxTimeSignature(maxTimeSignature).TargetTimeSignature(targetTimeSignature).MinValence(minValence).MaxValence(maxValence).TargetValence(targetValence).Execute()

Get Recommendations 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	seedArtists := "4NHQUGzhtTLFvgF5SZesLK" // string | 
	seedGenres := "classical,country" // string | 
	seedTracks := "0c6xIDDpzE81m2q797ordA" // string | 
	limit := int64(10) // int64 |  (optional) (default to 20)
	market := "ES" // string |  (optional)
	minAcousticness := float32(8.14) // float32 |  (optional)
	maxAcousticness := float32(8.14) // float32 |  (optional)
	targetAcousticness := float32(8.14) // float32 |  (optional)
	minDanceability := float32(8.14) // float32 |  (optional)
	maxDanceability := float32(8.14) // float32 |  (optional)
	targetDanceability := float32(8.14) // float32 |  (optional)
	minDurationMs := int64(56) // int64 |  (optional)
	maxDurationMs := int64(56) // int64 |  (optional)
	targetDurationMs := int64(56) // int64 |  (optional)
	minEnergy := float32(8.14) // float32 |  (optional)
	maxEnergy := float32(8.14) // float32 |  (optional)
	targetEnergy := float32(8.14) // float32 |  (optional)
	minInstrumentalness := float32(8.14) // float32 |  (optional)
	maxInstrumentalness := float32(8.14) // float32 |  (optional)
	targetInstrumentalness := float32(8.14) // float32 |  (optional)
	minKey := int64(56) // int64 |  (optional)
	maxKey := int64(56) // int64 |  (optional)
	targetKey := int64(56) // int64 |  (optional)
	minLiveness := float32(8.14) // float32 |  (optional)
	maxLiveness := float32(8.14) // float32 |  (optional)
	targetLiveness := float32(8.14) // float32 |  (optional)
	minLoudness := float32(8.14) // float32 |  (optional)
	maxLoudness := float32(8.14) // float32 |  (optional)
	targetLoudness := float32(8.14) // float32 |  (optional)
	minMode := int64(56) // int64 |  (optional)
	maxMode := int64(56) // int64 |  (optional)
	targetMode := int64(56) // int64 |  (optional)
	minPopularity := int64(56) // int64 |  (optional)
	maxPopularity := int64(56) // int64 |  (optional)
	targetPopularity := int64(56) // int64 |  (optional)
	minSpeechiness := float32(8.14) // float32 |  (optional)
	maxSpeechiness := float32(8.14) // float32 |  (optional)
	targetSpeechiness := float32(8.14) // float32 |  (optional)
	minTempo := float32(8.14) // float32 |  (optional)
	maxTempo := float32(8.14) // float32 |  (optional)
	targetTempo := float32(8.14) // float32 |  (optional)
	minTimeSignature := int64(56) // int64 |  (optional)
	maxTimeSignature := int64(56) // int64 |  (optional)
	targetTimeSignature := int64(56) // int64 |  (optional)
	minValence := float32(8.14) // float32 |  (optional)
	maxValence := float32(8.14) // float32 |  (optional)
	targetValence := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.GetRecommendations(context.Background()).SeedArtists(seedArtists).SeedGenres(seedGenres).SeedTracks(seedTracks).Limit(limit).Market(market).MinAcousticness(minAcousticness).MaxAcousticness(maxAcousticness).TargetAcousticness(targetAcousticness).MinDanceability(minDanceability).MaxDanceability(maxDanceability).TargetDanceability(targetDanceability).MinDurationMs(minDurationMs).MaxDurationMs(maxDurationMs).TargetDurationMs(targetDurationMs).MinEnergy(minEnergy).MaxEnergy(maxEnergy).TargetEnergy(targetEnergy).MinInstrumentalness(minInstrumentalness).MaxInstrumentalness(maxInstrumentalness).TargetInstrumentalness(targetInstrumentalness).MinKey(minKey).MaxKey(maxKey).TargetKey(targetKey).MinLiveness(minLiveness).MaxLiveness(maxLiveness).TargetLiveness(targetLiveness).MinLoudness(minLoudness).MaxLoudness(maxLoudness).TargetLoudness(targetLoudness).MinMode(minMode).MaxMode(maxMode).TargetMode(targetMode).MinPopularity(minPopularity).MaxPopularity(maxPopularity).TargetPopularity(targetPopularity).MinSpeechiness(minSpeechiness).MaxSpeechiness(maxSpeechiness).TargetSpeechiness(targetSpeechiness).MinTempo(minTempo).MaxTempo(maxTempo).TargetTempo(targetTempo).MinTimeSignature(minTimeSignature).MaxTimeSignature(maxTimeSignature).TargetTimeSignature(targetTimeSignature).MinValence(minValence).MaxValence(maxValence).TargetValence(targetValence).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.GetRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecommendations`: RecommendationsObject
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.GetRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seedArtists** | **string** |  | 
 **seedGenres** | **string** |  | 
 **seedTracks** | **string** |  | 
 **limit** | **int64** |  | [default to 20]
 **market** | **string** |  | 
 **minAcousticness** | **float32** |  | 
 **maxAcousticness** | **float32** |  | 
 **targetAcousticness** | **float32** |  | 
 **minDanceability** | **float32** |  | 
 **maxDanceability** | **float32** |  | 
 **targetDanceability** | **float32** |  | 
 **minDurationMs** | **int64** |  | 
 **maxDurationMs** | **int64** |  | 
 **targetDurationMs** | **int64** |  | 
 **minEnergy** | **float32** |  | 
 **maxEnergy** | **float32** |  | 
 **targetEnergy** | **float32** |  | 
 **minInstrumentalness** | **float32** |  | 
 **maxInstrumentalness** | **float32** |  | 
 **targetInstrumentalness** | **float32** |  | 
 **minKey** | **int64** |  | 
 **maxKey** | **int64** |  | 
 **targetKey** | **int64** |  | 
 **minLiveness** | **float32** |  | 
 **maxLiveness** | **float32** |  | 
 **targetLiveness** | **float32** |  | 
 **minLoudness** | **float32** |  | 
 **maxLoudness** | **float32** |  | 
 **targetLoudness** | **float32** |  | 
 **minMode** | **int64** |  | 
 **maxMode** | **int64** |  | 
 **targetMode** | **int64** |  | 
 **minPopularity** | **int64** |  | 
 **maxPopularity** | **int64** |  | 
 **targetPopularity** | **int64** |  | 
 **minSpeechiness** | **float32** |  | 
 **maxSpeechiness** | **float32** |  | 
 **targetSpeechiness** | **float32** |  | 
 **minTempo** | **float32** |  | 
 **maxTempo** | **float32** |  | 
 **targetTempo** | **float32** |  | 
 **minTimeSignature** | **int64** |  | 
 **maxTimeSignature** | **int64** |  | 
 **targetTimeSignature** | **int64** |  | 
 **minValence** | **float32** |  | 
 **maxValence** | **float32** |  | 
 **targetValence** | **float32** |  | 

### Return type

[**RecommendationsObject**](RecommendationsObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeveralAudioFeatures

> GetSeveralAudioFeatures200Response GetSeveralAudioFeatures(ctx).Ids(ids).Execute()

Get Several Tracks' Audio Features 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	ids := "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.GetSeveralAudioFeatures(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.GetSeveralAudioFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeveralAudioFeatures`: GetSeveralAudioFeatures200Response
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.GetSeveralAudioFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeveralAudioFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 

### Return type

[**GetSeveralAudioFeatures200Response**](GetSeveralAudioFeatures200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeveralTracks

> GetAnArtistsTopTracks200Response GetSeveralTracks(ctx).Ids(ids).Market(market).Execute()

Get Several Tracks 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	ids := "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B" // string | 
	market := "ES" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.GetSeveralTracks(context.Background()).Ids(ids).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.GetSeveralTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeveralTracks`: GetAnArtistsTopTracks200Response
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.GetSeveralTracks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeveralTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **market** | **string** |  | 

### Return type

[**GetAnArtistsTopTracks200Response**](GetAnArtistsTopTracks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrack

> TrackObject GetTrack(ctx, id).Market(market).Execute()

Get Track 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	id := "11dFghVXANMlKmJXsNCbNl" // string | 
	market := "ES" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.GetTrack(context.Background(), id).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.GetTrack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrack`: TrackObject
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.GetTrack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 

### Return type

[**TrackObject**](TrackObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersSavedTracks

> PagingSavedTrackObject GetUsersSavedTracks(ctx).Market(market).Limit(limit).Offset(offset).Execute()

Get User's Saved Tracks 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	market := "ES" // string |  (optional)
	limit := int64(10) // int64 |  (optional) (default to 20)
	offset := int64(5) // int64 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.GetUsersSavedTracks(context.Background()).Market(market).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.GetUsersSavedTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersSavedTracks`: PagingSavedTrackObject
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.GetUsersSavedTracks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersSavedTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** |  | 
 **limit** | **int64** |  | [default to 20]
 **offset** | **int64** |  | [default to 0]

### Return type

[**PagingSavedTrackObject**](PagingSavedTrackObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersTopArtistsAndTracks

> GetUsersTopArtistsAndTracks200Response GetUsersTopArtistsAndTracks(ctx, type_).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()

Get User's Top Items 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	type_ := "type__example" // string | 
	timeRange := "medium_term" // string |  (optional) (default to "medium_term")
	limit := int64(10) // int64 |  (optional) (default to 20)
	offset := int64(5) // int64 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.GetUsersTopArtistsAndTracks(context.Background(), type_).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.GetUsersTopArtistsAndTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersTopArtistsAndTracks`: GetUsersTopArtistsAndTracks200Response
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.GetUsersTopArtistsAndTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersTopArtistsAndTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeRange** | **string** |  | [default to &quot;medium_term&quot;]
 **limit** | **int64** |  | [default to 20]
 **offset** | **int64** |  | [default to 0]

### Return type

[**GetUsersTopArtistsAndTracks200Response**](GetUsersTopArtistsAndTracks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTracksPlaylist

> ReorderOrReplacePlaylistsTracks200Response RemoveTracksPlaylist(ctx, playlistId).RemoveTracksPlaylistRequest(removeTracksPlaylistRequest).Execute()

Remove Playlist Items 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	playlistId := "3cEYpjA9oz9GiPac4AsH4n" // string | 
	removeTracksPlaylistRequest := *openapiclient.NewRemoveTracksPlaylistRequest([]openapiclient.RemoveTracksPlaylistRequestTracksInner{*openapiclient.NewRemoveTracksPlaylistRequestTracksInner()}) // RemoveTracksPlaylistRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.RemoveTracksPlaylist(context.Background(), playlistId).RemoveTracksPlaylistRequest(removeTracksPlaylistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.RemoveTracksPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveTracksPlaylist`: ReorderOrReplacePlaylistsTracks200Response
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.RemoveTracksPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTracksPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeTracksPlaylistRequest** | [**RemoveTracksPlaylistRequest**](RemoveTracksPlaylistRequest.md) |  | 

### Return type

[**ReorderOrReplacePlaylistsTracks200Response**](ReorderOrReplacePlaylistsTracks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTracksUser

> RemoveTracksUser(ctx).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()

Remove User's Saved Tracks 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	ids := "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B" // string | 
	saveAlbumsUserRequest := *openapiclient.NewSaveAlbumsUserRequest() // SaveAlbumsUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TracksAPI.RemoveTracksUser(context.Background()).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.RemoveTracksUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTracksUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **saveAlbumsUserRequest** | [**SaveAlbumsUserRequest**](SaveAlbumsUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReorderOrReplacePlaylistsTracks

> ReorderOrReplacePlaylistsTracks200Response ReorderOrReplacePlaylistsTracks(ctx, playlistId).Uris(uris).ReorderOrReplacePlaylistsTracksRequest(reorderOrReplacePlaylistsTracksRequest).Execute()

Update Playlist Items 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	playlistId := "3cEYpjA9oz9GiPac4AsH4n" // string | 
	uris := "uris_example" // string |  (optional)
	reorderOrReplacePlaylistsTracksRequest := *openapiclient.NewReorderOrReplacePlaylistsTracksRequest() // ReorderOrReplacePlaylistsTracksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracksAPI.ReorderOrReplacePlaylistsTracks(context.Background(), playlistId).Uris(uris).ReorderOrReplacePlaylistsTracksRequest(reorderOrReplacePlaylistsTracksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.ReorderOrReplacePlaylistsTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderOrReplacePlaylistsTracks`: ReorderOrReplacePlaylistsTracks200Response
	fmt.Fprintf(os.Stdout, "Response from `TracksAPI.ReorderOrReplacePlaylistsTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReorderOrReplacePlaylistsTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uris** | **string** |  | 
 **reorderOrReplacePlaylistsTracksRequest** | [**ReorderOrReplacePlaylistsTracksRequest**](ReorderOrReplacePlaylistsTracksRequest.md) |  | 

### Return type

[**ReorderOrReplacePlaylistsTracks200Response**](ReorderOrReplacePlaylistsTracks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveTracksUser

> SaveTracksUser(ctx).Ids(ids).SaveTracksUserRequest(saveTracksUserRequest).Execute()

Save Tracks for Current User 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ericflores108/samba"
)

func main() {
	ids := "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B" // string | 
	saveTracksUserRequest := *openapiclient.NewSaveTracksUserRequest() // SaveTracksUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TracksAPI.SaveTracksUser(context.Background()).Ids(ids).SaveTracksUserRequest(saveTracksUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracksAPI.SaveTracksUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveTracksUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **saveTracksUserRequest** | [**SaveTracksUserRequest**](SaveTracksUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

