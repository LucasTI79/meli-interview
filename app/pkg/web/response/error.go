package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

func Error(w http.ResponseWriter, statusCode int, message string, code string) {
	// default status code
	defaultStatusCode := http.StatusInternalServerError
	// check if status code is valid
	if statusCode > 299 && statusCode < 600 {
		defaultStatusCode = statusCode
	}

	// response
	body := errorResponse{
		Status:  http.StatusText(defaultStatusCode),
		Message: message,
		Code:    code,
	}
	bytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write response
	// - set header: before code due to it sets by default "text/plain"
	w.Header().Set("Content-Type", "application/json")
	// - set status code
	w.WriteHeader(defaultStatusCode)
	// - write body
	if _, err := w.Write(bytes); err != nil {
		fmt.Println("error writing response:", err)
	}
}

func Errorf(w http.ResponseWriter, statusCode int, code, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	Error(w, statusCode, message, code)
}
