package response

import (
	"encoding/json"
	"net/http"
	"time"
)

type (
	Error struct {
		statusCode int
		Errors     []string
		Input      interface{}
	}

	ErrorResponse struct {
		Content    interface{} `json:"content"`
		Details    []string    `json:"details"`
		StatusCode int         `json:"statusCode"`
		Timestamp  string      `json:"timestamp"`
	}
)

func NewError(err error, status int, input interface{}) *Error {
	return &Error{
		statusCode: status,
		Errors:     []string{err.Error()},
		Input:      input,
	}
}

func NewErrorMessage(messages []string, status int) *Error {
	return &Error{
		statusCode: status,
		Errors:     messages,
	}
}

func (e Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)

	response := ErrorResponse{
		StatusCode: e.statusCode,
		Timestamp:  time.Now().Format(time.RFC3339),
		Details:    e.Errors,
		Content:    e.Input,
	}

	return json.NewEncoder(w).Encode(response)
}
