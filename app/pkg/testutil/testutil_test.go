package testutil_test

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/meli-interview/pkg/testutil"
	"github.com/stretchr/testify/require"
)

func TestWithUrlParam(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	req = testutil.WithUrlParam(t, req, "id", "123")

	// Retrieve Chi route context
	chiCtx := chi.RouteContext(req.Context())
	require.NotNil(t, chiCtx, "Chi context should not be nil")

	value := chiCtx.URLParam("id")
	require.Equal(t, "123", value, "URL param 'id' should be '123'")
}

func TestWithUrlParamst(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	params := map[string]string{
		"id":    "123",
		"token": "abc",
	}
	req = testutil.WithUrlParamst(t, req, params)

	chiCtx := chi.RouteContext(req.Context())
	require.NotNil(t, chiCtx, "Chi context should not be nil")

	for k, expected := range params {
		got := chiCtx.URLParam(k)
		require.Equal(t, expected, got, "URL param %s mismatch", k)
	}
}
