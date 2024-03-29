/*
FunLess Platfom API

Testing AppAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/funlessdev/fl-client-sdk-go"
)

func Test_openapi_AppAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppAPIService CreateApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AppAPI.CreateApp(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppAPIService DeleteApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var moduleName string

		httpRes, err := apiClient.AppAPI.DeleteApp(context.Background(), moduleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppAPIService ListApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppAPI.ListApp(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppAPIService ShowAppByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var moduleName string

		resp, httpRes, err := apiClient.AppAPI.ShowAppByName(context.Background(), moduleName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
