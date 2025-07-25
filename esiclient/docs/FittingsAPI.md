# \FittingsAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCharactersCharacterIdFittingsFittingId**](FittingsAPI.md#DeleteCharactersCharacterIdFittingsFittingId) | **Delete** /characters/{character_id}/fittings/{fitting_id} | Delete fitting
[**GetCharactersCharacterIdFittings**](FittingsAPI.md#GetCharactersCharacterIdFittings) | **Get** /characters/{character_id}/fittings | Get fittings
[**PostCharactersCharacterIdFittings**](FittingsAPI.md#PostCharactersCharacterIdFittings) | **Post** /characters/{character_id}/fittings | Create fitting



## DeleteCharactersCharacterIdFittingsFittingId

> interface{} DeleteCharactersCharacterIdFittingsFittingId(ctx, characterId, fittingId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Delete fitting



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
	fittingId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FittingsAPI.DeleteCharactersCharacterIdFittingsFittingId(context.Background(), characterId, fittingId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FittingsAPI.DeleteCharactersCharacterIdFittingsFittingId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCharactersCharacterIdFittingsFittingId`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FittingsAPI.DeleteCharactersCharacterIdFittingsFittingId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 
**fittingId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCharactersCharacterIdFittingsFittingIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCharactersCharacterIdFittings

> []CharactersCharacterIdFittingsGetInner GetCharactersCharacterIdFittings(ctx, characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get fittings



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FittingsAPI.GetCharactersCharacterIdFittings(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FittingsAPI.GetCharactersCharacterIdFittings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersCharacterIdFittings`: []CharactersCharacterIdFittingsGetInner
	fmt.Fprintf(os.Stdout, "Response from `FittingsAPI.GetCharactersCharacterIdFittings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersCharacterIdFittingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 

### Return type

[**[]CharactersCharacterIdFittingsGetInner**](CharactersCharacterIdFittingsGetInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCharactersCharacterIdFittings

> CharactersCharacterIdFittingsPost PostCharactersCharacterIdFittings(ctx, characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).PostCharactersCharacterIdFittingsRequest(postCharactersCharacterIdFittingsRequest).Execute()

Create fitting



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
	postCharactersCharacterIdFittingsRequest := *openapiclient.NewPostCharactersCharacterIdFittingsRequest("Description_example", []openapiclient.PostCharactersCharacterIdFittingsRequestItemsInner{*openapiclient.NewPostCharactersCharacterIdFittingsRequestItemsInner("Flag_example", int64(123), int64(123))}, "Name_example", int64(123)) // PostCharactersCharacterIdFittingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FittingsAPI.PostCharactersCharacterIdFittings(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).PostCharactersCharacterIdFittingsRequest(postCharactersCharacterIdFittingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FittingsAPI.PostCharactersCharacterIdFittings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCharactersCharacterIdFittings`: CharactersCharacterIdFittingsPost
	fmt.Fprintf(os.Stdout, "Response from `FittingsAPI.PostCharactersCharacterIdFittings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCharactersCharacterIdFittingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 
 **postCharactersCharacterIdFittingsRequest** | [**PostCharactersCharacterIdFittingsRequest**](PostCharactersCharacterIdFittingsRequest.md) |  | 

### Return type

[**CharactersCharacterIdFittingsPost**](CharactersCharacterIdFittingsPost.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

