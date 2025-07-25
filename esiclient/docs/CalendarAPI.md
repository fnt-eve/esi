# \CalendarAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdCalendar**](CalendarAPI.md#GetCharactersCharacterIdCalendar) | **Get** /characters/{character_id}/calendar | List calendar event summaries
[**GetCharactersCharacterIdCalendarEventId**](CalendarAPI.md#GetCharactersCharacterIdCalendarEventId) | **Get** /characters/{character_id}/calendar/{event_id} | Get an event
[**GetCharactersCharacterIdCalendarEventIdAttendees**](CalendarAPI.md#GetCharactersCharacterIdCalendarEventIdAttendees) | **Get** /characters/{character_id}/calendar/{event_id}/attendees | Get attendees
[**PutCharactersCharacterIdCalendarEventId**](CalendarAPI.md#PutCharactersCharacterIdCalendarEventId) | **Put** /characters/{character_id}/calendar/{event_id} | Respond to an event



## GetCharactersCharacterIdCalendar

> []CharactersCharacterIdCalendarGetInner GetCharactersCharacterIdCalendar(ctx, characterId).XCompatibilityDate(xCompatibilityDate).FromEvent(fromEvent).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

List calendar event summaries



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
	fromEvent := int64(789) // int64 |  (optional)
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.GetCharactersCharacterIdCalendar(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).FromEvent(fromEvent).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.GetCharactersCharacterIdCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersCharacterIdCalendar`: []CharactersCharacterIdCalendarGetInner
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.GetCharactersCharacterIdCalendar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersCharacterIdCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **fromEvent** | **int64** |  | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 

### Return type

[**[]CharactersCharacterIdCalendarGetInner**](CharactersCharacterIdCalendarGetInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCharactersCharacterIdCalendarEventId

> CharactersCharacterIdCalendarEventIdGet GetCharactersCharacterIdCalendarEventId(ctx, characterId, eventId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get an event



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
	eventId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.GetCharactersCharacterIdCalendarEventId(context.Background(), characterId, eventId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.GetCharactersCharacterIdCalendarEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersCharacterIdCalendarEventId`: CharactersCharacterIdCalendarEventIdGet
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.GetCharactersCharacterIdCalendarEventId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 
**eventId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersCharacterIdCalendarEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 

### Return type

[**CharactersCharacterIdCalendarEventIdGet**](CharactersCharacterIdCalendarEventIdGet.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCharactersCharacterIdCalendarEventIdAttendees

> []CharactersCharacterIdCalendarEventIdAttendeesGetInner GetCharactersCharacterIdCalendarEventIdAttendees(ctx, characterId, eventId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get attendees



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
	eventId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.GetCharactersCharacterIdCalendarEventIdAttendees(context.Background(), characterId, eventId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.GetCharactersCharacterIdCalendarEventIdAttendees``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersCharacterIdCalendarEventIdAttendees`: []CharactersCharacterIdCalendarEventIdAttendeesGetInner
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.GetCharactersCharacterIdCalendarEventIdAttendees`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 
**eventId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersCharacterIdCalendarEventIdAttendeesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 

### Return type

[**[]CharactersCharacterIdCalendarEventIdAttendeesGetInner**](CharactersCharacterIdCalendarEventIdAttendeesGetInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCharactersCharacterIdCalendarEventId

> interface{} PutCharactersCharacterIdCalendarEventId(ctx, characterId, eventId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).PutCharactersCharacterIdCalendarEventIdRequest(putCharactersCharacterIdCalendarEventIdRequest).Execute()

Respond to an event



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
	eventId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "en" // string | The language to use for the response. Defaults to 'en'. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "tranquility" // string | The tenant ID for the request. Defaults to 'tranquility'. (optional)
	putCharactersCharacterIdCalendarEventIdRequest := *openapiclient.NewPutCharactersCharacterIdCalendarEventIdRequest("Response_example") // PutCharactersCharacterIdCalendarEventIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.PutCharactersCharacterIdCalendarEventId(context.Background(), characterId, eventId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).PutCharactersCharacterIdCalendarEventIdRequest(putCharactersCharacterIdCalendarEventIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.PutCharactersCharacterIdCalendarEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCharactersCharacterIdCalendarEventId`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.PutCharactersCharacterIdCalendarEventId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 
**eventId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCharactersCharacterIdCalendarEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. Defaults to &#39;en&#39;. | 
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. Defaults to &#39;tranquility&#39;. | 
 **putCharactersCharacterIdCalendarEventIdRequest** | [**PutCharactersCharacterIdCalendarEventIdRequest**](PutCharactersCharacterIdCalendarEventIdRequest.md) |  | 

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

