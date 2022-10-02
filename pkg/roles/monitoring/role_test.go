package monitoring_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"beryju.io/gravity/pkg/instance"
	"beryju.io/gravity/pkg/roles/monitoring"
	"beryju.io/gravity/pkg/tests"
	"github.com/stretchr/testify/assert"
)

func TestRoleStartNoConfig(t *testing.T) {
	rootInst := instance.New()
	inst := rootInst.ForRole("monitoring")
	role := monitoring.New(inst)
	assert.NotNil(t, role)
	ctx := tests.Context()
	assert.Nil(t, role.Start(ctx, []byte{}))
	defer role.Stop()
}

func TestRoleStartEmptyConfig(t *testing.T) {
	rootInst := instance.New()
	inst := rootInst.ForRole("monitoring")
	role := monitoring.New(inst)
	assert.NotNil(t, role)
	ctx := tests.Context()
	assert.Nil(t, role.Start(ctx, []byte("{}")))
	defer role.Stop()
}

func TestRoleHealth(t *testing.T) {
	rootInst := instance.New()
	inst := rootInst.ForRole("monitoring")
	role := monitoring.New(inst)
	assert.NotNil(t, role)
	ctx := tests.Context()
	assert.Nil(t, role.Start(ctx, []byte("{}")))
	defer role.Stop()
	rr := httptest.NewRecorder()
	role.HandleHealthLive(rr, httptest.NewRequest("GET", "/", nil))
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
}

func TestMetricsScrape(t *testing.T) {
	rootInst := instance.New()
	inst := rootInst.ForRole("monitoring")
	role := monitoring.New(inst)
	assert.NotNil(t, role)
	ctx := tests.Context()
	assert.Nil(t, role.Start(ctx, []byte("{}")))
	defer role.Stop()
	rr := httptest.NewRecorder()
	role.HandleMetrics(rr, httptest.NewRequest("GET", "/", nil))
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
}