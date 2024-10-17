/*
gravity

Testing RolesDnsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	"testing"

	openapiclient "beryju.io/gravity/api"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_RolesDnsApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RolesDnsApiService DnsDeleteRecords", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.RolesDnsApi.DnsDeleteRecords(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RolesDnsApiService DnsDeleteZones", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.RolesDnsApi.DnsDeleteZones(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RolesDnsApiService DnsGetRecords", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RolesDnsApi.DnsGetRecords(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RolesDnsApiService DnsGetRoleConfig", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RolesDnsApi.DnsGetRoleConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RolesDnsApiService DnsGetZones", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RolesDnsApi.DnsGetZones(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RolesDnsApiService DnsPutRecords", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.RolesDnsApi.DnsPutRecords(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RolesDnsApiService DnsPutRoleConfig", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.RolesDnsApi.DnsPutRoleConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RolesDnsApiService DnsPutZones", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.RolesDnsApi.DnsPutZones(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
