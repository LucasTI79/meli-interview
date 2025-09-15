package router_test

import (
	"net/http/httptest"
	"testing"

	"github.com/lucasti79/meli-interview/cmd/http/router"
	"github.com/lucasti79/meli-interview/internal/factory"
	"github.com/stretchr/testify/assert"
)

func TestRouterMounts(t *testing.T) {
	factory := &factory.AppFactory{}

	r := router.NewRouter().MapRoutes(factory)

	req := httptest.NewRequest("GET", "/api/v1/products", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.NotNil(t, resp)
	assert.Contains(t, resp.Body.String(), "")
}
