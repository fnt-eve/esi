# \AssetsAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdAssets**](AssetsAPI.md#GetCharactersCharacterIdAssets) | **Get** /characters/{character_id}/assets | Get character assets
[**GetCorporationsCorporationIdAssets**](AssetsAPI.md#GetCorporationsCorporationIdAssets) | **Get** /corporations/{corporation_id}/assets | Get corporation assets
[**PostCharactersCharacterIdAssetsLocations**](AssetsAPI.md#PostCharactersCharacterIdAssetsLocations) | **Post** /characters/{character_id}/assets/locations | Get character asset locations
[**PostCharactersCharacterIdAssetsNames**](AssetsAPI.md#PostCharactersCharacterIdAssetsNames) | **Post** /characters/{character_id}/assets/names | Get character asset names
[**PostCorporationsCorporationIdAssetsLocations**](AssetsAPI.md#PostCorporationsCorporationIdAssetsLocations) | **Post** /corporations/{corporation_id}/assets/locations | Get corporation asset locations
[**PostCorporationsCorporationIdAssetsNames**](AssetsAPI.md#PostCorporationsCorporationIdAssetsNames) | **Post** /corporations/{corporation_id}/assets/names | Get corporation asset names



## GetCharactersCharacterIdAssets

> []CharactersCharacterIdAssetsGetInner GetCharactersCharacterIdAssets(ctx, characterId).XCompatibilityDate(xCompatibilityDate).Page(page).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get character assets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fnt-eve/esi/esiclient"
)

func main() {
	characterId := int64(789) // int64 | The ID of the character
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	page := int32(56) // int32 |  (optional)
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.GetCharactersCharacterIdAssets(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).Page(page).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetCharactersCharacterIdAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersCharacterIdAssets`: []CharactersCharacterIdAssetsGetInner
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetCharactersCharacterIdAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersCharacterIdAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **page** | **int32** |  | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 

### Return type

[**[]CharactersCharacterIdAssetsGetInner**](CharactersCharacterIdAssetsGetInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorporationsCorporationIdAssets

> []CorporationsCorporationIdAssetsGetInner GetCorporationsCorporationIdAssets(ctx, corporationId).XCompatibilityDate(xCompatibilityDate).Page(page).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get corporation assets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fnt-eve/esi/esiclient"
)

func main() {
	corporationId := int64(789) // int64 | The ID of the corporation
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	page := int32(56) // int32 |  (optional)
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.GetCorporationsCorporationIdAssets(context.Background(), corporationId).XCompatibilityDate(xCompatibilityDate).Page(page).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetCorporationsCorporationIdAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCorporationsCorporationIdAssets`: []CorporationsCorporationIdAssetsGetInner
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetCorporationsCorporationIdAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int64** | The ID of the corporation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorporationsCorporationIdAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **page** | **int32** |  | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 

### Return type

[**[]CorporationsCorporationIdAssetsGetInner**](CorporationsCorporationIdAssetsGetInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCharactersCharacterIdAssetsLocations

> []CharactersCharacterIdAssetsLocationsPostInner PostCharactersCharacterIdAssetsLocations(ctx, characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).RequestBody(requestBody).Execute()

Get character asset locations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fnt-eve/esi/esiclient"
)

func main() {
	characterId := int64(789) // int64 | The ID of the character
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)
	requestBody := []int64{int64(123)} // []int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.PostCharactersCharacterIdAssetsLocations(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.PostCharactersCharacterIdAssetsLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCharactersCharacterIdAssetsLocations`: []CharactersCharacterIdAssetsLocationsPostInner
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.PostCharactersCharacterIdAssetsLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCharactersCharacterIdAssetsLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 
 **requestBody** | **[]int64** |  | 

### Return type

[**[]CharactersCharacterIdAssetsLocationsPostInner**](CharactersCharacterIdAssetsLocationsPostInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCharactersCharacterIdAssetsNames

> []CharactersCharacterIdAssetsNamesPostInner PostCharactersCharacterIdAssetsNames(ctx, characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).RequestBody(requestBody).Execute()

Get character asset names



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fnt-eve/esi/esiclient"
)

func main() {
	characterId := int64(789) // int64 | The ID of the character
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)
	requestBody := []int64{int64(123)} // []int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.PostCharactersCharacterIdAssetsNames(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.PostCharactersCharacterIdAssetsNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCharactersCharacterIdAssetsNames`: []CharactersCharacterIdAssetsNamesPostInner
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.PostCharactersCharacterIdAssetsNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCharactersCharacterIdAssetsNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 
 **requestBody** | **[]int64** |  | 

### Return type

[**[]CharactersCharacterIdAssetsNamesPostInner**](CharactersCharacterIdAssetsNamesPostInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCorporationsCorporationIdAssetsLocations

> []CharactersCharacterIdAssetsLocationsPostInner PostCorporationsCorporationIdAssetsLocations(ctx, corporationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).RequestBody(requestBody).Execute()

Get corporation asset locations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fnt-eve/esi/esiclient"
)

func main() {
	corporationId := int64(789) // int64 | The ID of the corporation
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)
	requestBody := []int64{int64(123)} // []int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.PostCorporationsCorporationIdAssetsLocations(context.Background(), corporationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.PostCorporationsCorporationIdAssetsLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCorporationsCorporationIdAssetsLocations`: []CharactersCharacterIdAssetsLocationsPostInner
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.PostCorporationsCorporationIdAssetsLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int64** | The ID of the corporation | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCorporationsCorporationIdAssetsLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 
 **requestBody** | **[]int64** |  | 

### Return type

[**[]CharactersCharacterIdAssetsLocationsPostInner**](CharactersCharacterIdAssetsLocationsPostInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCorporationsCorporationIdAssetsNames

> []CharactersCharacterIdAssetsNamesPostInner PostCorporationsCorporationIdAssetsNames(ctx, corporationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).RequestBody(requestBody).Execute()

Get corporation asset names



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fnt-eve/esi/esiclient"
)

func main() {
	corporationId := int64(789) // int64 | The ID of the corporation
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)
	requestBody := []int64{int64(123)} // []int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.PostCorporationsCorporationIdAssetsNames(context.Background(), corporationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.PostCorporationsCorporationIdAssetsNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCorporationsCorporationIdAssetsNames`: []CharactersCharacterIdAssetsNamesPostInner
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.PostCorporationsCorporationIdAssetsNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int64** | The ID of the corporation | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCorporationsCorporationIdAssetsNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 
 **requestBody** | **[]int64** |  | 

### Return type

[**[]CharactersCharacterIdAssetsNamesPostInner**](CharactersCharacterIdAssetsNamesPostInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

