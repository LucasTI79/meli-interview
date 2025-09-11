package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type dataResponseType struct {
	Data any `json:"data"`
}

// JSON writes json response
func JSON(w http.ResponseWriter, code int, body any) {
	// check body
	if body == nil {
		w.WriteHeader(code)
		return
	}

	// marshal body
	bytes, err := json.Marshal(body)
	if err != nil {
		// default error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// set header (before code due to it sets by default "text/plain")
	w.Header().Set("Content-Type", "application/json")

	// set status code
	w.WriteHeader(code)

	// write body
	if _, err := w.Write(bytes); err != nil {
		fmt.Println("error writing response:", err)
	}
} // JSON writes json response

func DataContentJson(w http.ResponseWriter, code int, body any) {

	response := dataResponseType{
		Data: body,
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	if _, err := w.Write(bytes); err != nil {
		fmt.Println("error writing response:", err)
	}
}
