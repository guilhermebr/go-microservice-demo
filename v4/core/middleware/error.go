package middleware

import (
	"encoding/json"
	"net/http"
)

var (
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
	statusCode := e.StatusCode

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(e)
}
