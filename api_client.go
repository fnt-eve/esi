package esi

import (
	"net/http"
	"strings"
	"time"

	"github.com/fnt-eve/esi/esiclient"
)

// Handle our context keys.
type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

// ContextOAuth2 is the context for GoESI authentication. Pass a tokenSource with this key to a context for an ESI API Call
var (
	ContextOAuth2 = esiclient.ContextOAuth2
)

// APIClient manages communication with the EVE Swagger Interface API
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	ESI *esiclient.APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(httpClient *http.Client, userAgent string) *APIClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &APIClient{}

	conf := esiclient.NewConfiguration()
	conf.UserAgent = userAgent
	conf.HTTPClient = httpClient
	c.ESI = esiclient.NewAPIClient(conf)
	return c
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}
