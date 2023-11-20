package web

import (
	"encoding/json"
	"net/http"
)

type suucessResponse struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Message string `json:"message"`
}

// handles success responses
func RespondSuccess(w http.ResponseWriter, data interface{}) {
	suucessResponse := suucessResponse{
		Data: data,
	}
	s, err := json.Marshal(suucessResponse)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(s)
}

// handles error responses
func RespondError(w http.ResponseWriter, code int, err error) {
	errorResponse := errorResponse{
		Message: err.Error(),
	}

	b, err := json.Marshal(errorResponse) // Marshal errorResponse to JSON
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "we encountered an error processing your request"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b) // Write the JSON byte slice to the response writer
}
