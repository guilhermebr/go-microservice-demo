package api

import (
	"encoding/json"
	"net/http"
)

var (
	AlertMissingData  = Alert{StatusCode: http.StatusBadRequest, Type: "missing_data", Message: "Missing required Data"}
	AlertExistingData = Alert{StatusCode: http.StatusOK, Type: "existing_data", Message: "Data already exists"}
)

// Alert
type Alert struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status,omitempty"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func (a Alert) Send(w http.ResponseWriter) error {
	statusCode := a.StatusCode

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(a)
}

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status,omitempty"`
	Type       string      `json:"type,omitempty"`
	Result     interface{} `json:"result,omitempty"`
}

func Success(result interface{}, status int) *Response {
	return &Response{
		Success:    true,
		StatusCode: status,
		Type:       "ok",
		Result:     result,
	}
}

func (r *Response) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	return json.NewEncoder(w).Encode(r)
}
