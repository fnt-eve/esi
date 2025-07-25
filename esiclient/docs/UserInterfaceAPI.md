# \UserInterfaceAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostUiAutopilotWaypoint**](UserInterfaceAPI.md#PostUiAutopilotWaypoint) | **Post** /ui/autopilot/waypoint | Set Autopilot Waypoint
[**PostUiOpenwindowContract**](UserInterfaceAPI.md#PostUiOpenwindowContract) | **Post** /ui/openwindow/contract | Open Contract Window
[**PostUiOpenwindowInformation**](UserInterfaceAPI.md#PostUiOpenwindowInformation) | **Post** /ui/openwindow/information | Open Information Window
[**PostUiOpenwindowMarketdetails**](UserInterfaceAPI.md#PostUiOpenwindowMarketdetails) | **Post** /ui/openwindow/marketdetails | Open Market Details
[**PostUiOpenwindowNewmail**](UserInterfaceAPI.md#PostUiOpenwindowNewmail) | **Post** /ui/openwindow/newmail | Open New Mail Window



## PostUiAutopilotWaypoint

> interface{} PostUiAutopilotWaypoint(ctx).AddToBeginning(addToBeginning).ClearOtherWaypoints(clearOtherWaypoints).DestinationId(destinationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Set Autopilot Waypoint



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
	addToBeginning := true // bool |  (default to false)
	clearOtherWaypoints := true // bool |  (default to false)
	destinationId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserInterfaceAPI.PostUiAutopilotWaypoint(context.Background()).AddToBeginning(addToBeginning).ClearOtherWaypoints(clearOtherWaypoints).DestinationId(destinationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInterfaceAPI.PostUiAutopilotWaypoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUiAutopilotWaypoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserInterfaceAPI.PostUiAutopilotWaypoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUiAutopilotWaypointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addToBeginning** | **bool** |  | [default to false]
 **clearOtherWaypoints** | **bool** |  | [default to false]
 **destinationId** | **int64** |  | 
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


## PostUiOpenwindowContract

> interface{} PostUiOpenwindowContract(ctx).ContractId(contractId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Open Contract Window



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
	contractId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserInterfaceAPI.PostUiOpenwindowContract(context.Background()).ContractId(contractId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInterfaceAPI.PostUiOpenwindowContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUiOpenwindowContract`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserInterfaceAPI.PostUiOpenwindowContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUiOpenwindowContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractId** | **int64** |  | 
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


## PostUiOpenwindowInformation

> interface{} PostUiOpenwindowInformation(ctx).TargetId(targetId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Open Information Window



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
	targetId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserInterfaceAPI.PostUiOpenwindowInformation(context.Background()).TargetId(targetId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInterfaceAPI.PostUiOpenwindowInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUiOpenwindowInformation`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserInterfaceAPI.PostUiOpenwindowInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUiOpenwindowInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetId** | **int64** |  | 
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


## PostUiOpenwindowMarketdetails

> interface{} PostUiOpenwindowMarketdetails(ctx).TypeId(typeId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Open Market Details



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
	typeId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserInterfaceAPI.PostUiOpenwindowMarketdetails(context.Background()).TypeId(typeId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInterfaceAPI.PostUiOpenwindowMarketdetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUiOpenwindowMarketdetails`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserInterfaceAPI.PostUiOpenwindowMarketdetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUiOpenwindowMarketdetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typeId** | **int64** |  | 
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


## PostUiOpenwindowNewmail

> interface{} PostUiOpenwindowNewmail(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).PostUiOpenwindowNewmailRequest(postUiOpenwindowNewmailRequest).Execute()

Open New Mail Window



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
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)
	postUiOpenwindowNewmailRequest := *openapiclient.NewPostUiOpenwindowNewmailRequest("Body_example", []int64{int64(123)}, "Subject_example") // PostUiOpenwindowNewmailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserInterfaceAPI.PostUiOpenwindowNewmail(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).PostUiOpenwindowNewmailRequest(postUiOpenwindowNewmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInterfaceAPI.PostUiOpenwindowNewmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUiOpenwindowNewmail`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserInterfaceAPI.PostUiOpenwindowNewmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUiOpenwindowNewmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 
 **postUiOpenwindowNewmailRequest** | [**PostUiOpenwindowNewmailRequest**](PostUiOpenwindowNewmailRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

