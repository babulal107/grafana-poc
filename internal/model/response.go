package model

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse forms API response with success result
func SuccessResponse(feed interface{}) Response {
	return Response{Code: http.StatusOK, Message: http.StatusText(http.StatusOK), Data: feed}
}

// ErrorResponse forms API response with error
func ErrorResponse(code int, err error) Response {
	return Response{Code: code, Message: err.Error()}
}
