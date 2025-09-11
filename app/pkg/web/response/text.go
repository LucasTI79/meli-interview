package response

import (
	"fmt"
	"net/http"
)

// Text writes text response
func Text(w http.ResponseWriter, code int, body string) {
	// set header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// set status code
	w.WriteHeader(code)

	// write body
	if _, err := w.Write([]byte(body)); err != nil {
		fmt.Println("error writing response:", err)
	}
}
