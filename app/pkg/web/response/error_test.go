package response_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lucasti79/meli-interview/pkg/web/response"
	"github.com/stretchr/testify/require"
)

// Tests for Error
func TestError(t *testing.T) {
	t.Run("should return status code 500 - invalid code", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		statusCode := 0
		message := "error message"
		code := "error_code"
		response.Error(rr, statusCode, message, code)

		// assert
		expectedCode := http.StatusInternalServerError
		expectedBody := `{"status":"Internal Server Error","message":"error message","code":"error_code"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
		require.Equal(t, expectedHeaders, rr.Header())
	})

	t.Run("should return status code 400", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		statusCode := http.StatusBadRequest
		message := "error message"
		code := "error_code"
		response.Error(rr, statusCode, message, code)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := `{"status":"Bad Request","message":"error message","code":"error_code"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
		require.Equal(t, expectedHeaders, rr.Header())
	})
}

// Tests for Errorf
func TestErrorf(t *testing.T) {
	t.Run("should return status code 500 - invalid code", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := "error_code"
		format := "error message %s"
		args := []interface{}{"arg"}
		response.Errorf(rr, http.StatusInternalServerError, code, format, args...)

		// assert
		expectedCode := http.StatusInternalServerError
		expectedBody := `{"status":"Internal Server Error","message":"error message arg","code":"error_code"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
		require.Equal(t, expectedHeaders, rr.Header())
	})

	t.Run("should return status code 400", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := "error_code"
		format := "error message %s"
		args := []interface{}{"arg"}
		response.Errorf(rr, http.StatusBadRequest, code, format, args...)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := `{"status":"Bad Request","message":"error message arg","code":"error_code"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
		require.Equal(t, expectedHeaders, rr.Header())
	})
}
