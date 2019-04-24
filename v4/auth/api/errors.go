package api

import (
	"encoding/json"
	"net/http"
)

var (
	ErrInvalidJson    = Error{StatusCode: http.StatusBadRequest, Type: "invalid_json", Message: "Invalid or malformed JSON"}
	ErrMissingData    = Error{StatusCode: http.StatusBadRequest, Type: "missing_data", Message: "Missing required Data"}
	ErrInternalServer = Error{StatusCode: http.StatusInternalServerError, Type: "server_error", Message: "Internal server Error"}
	ErrUnauthorized   = Error{StatusCode: http.StatusUnauthorized, Type: "unauthorized", Message: "Unauthorized"}
	ErrForbidden      = Error{StatusCode: http.StatusForbidden, Type: "forbidden", Message: "Forbidden"}
)

// Error
type Error struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status,omitempty"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func (e Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)
	return json.NewEncoder(w).Encode(e)
}
