/*
FunLess Platfom API

Testing ModulesAPIService

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

func Test_openapi_ModulesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ModulesAPIService CreateModule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ModulesAPI.CreateModule(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModulesAPIService DeleteModule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var moduleName string

		httpRes, err := apiClient.ModulesAPI.DeleteModule(context.Background(), moduleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModulesAPIService ListModules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ModulesAPI.ListModules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModulesAPIService ShowModuleByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var moduleName string

		resp, httpRes, err := apiClient.ModulesAPI.ShowModuleByName(context.Background(), moduleName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModulesAPIService UpdateModule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var moduleName string

		httpRes, err := apiClient.ModulesAPI.UpdateModule(context.Background(), moduleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
