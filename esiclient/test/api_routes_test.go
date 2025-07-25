/*
EVE Stellar Information (ESI) - tranquility

Testing RoutesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package esiclient

import (
	"context"
	"testing"

	openapiclient "github.com/fnt-eve/esi/esiclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_esiclient_RoutesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoutesAPIService GetRouteOriginDestination", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var destination int64
		var origin int64

		resp, httpRes, err := apiClient.RoutesAPI.GetRouteOriginDestination(context.Background(), destination, origin).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
