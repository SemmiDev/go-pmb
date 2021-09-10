package web

import (
	"fmt"
	"net/http"
)

type Meta struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type HttpResponse struct {
	Meta *Meta       `json:"meta"`
	Data interface{} `json:"data"`
}

func NewHttpResponse(message interface{}, code int, status string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Meta: &Meta{
			Code:    code,
			Status:  status,
			Message: message,
		},
		Data: data,
	}
}

func formatError(message interface{}) map[string]string {
	mapError := make(map[string]string)
	mapError["error"] = fmt.Sprintf("%v", message)
	return mapError
}

func Ok(data interface{}) *HttpResponse {
	return NewHttpResponse(nil, http.StatusOK, http.StatusText(http.StatusOK), data)
}

func NoContent() *HttpResponse {
	return NewHttpResponse(nil, http.StatusOK, http.StatusText(http.StatusOK), nil)
}

func Created(data interface{}) *HttpResponse {
	return NewHttpResponse(nil, http.StatusCreated, http.StatusText(http.StatusCreated), data)
}

func BadRequest(err interface{}) *HttpResponse {
	return NewHttpResponse(formatError(err), http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
}

func NotFound(err interface{}) *HttpResponse {
	return NewHttpResponse(formatError(err), http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
}

func UnprocessableEntity(err interface{}) *HttpResponse {
	return NewHttpResponse(formatError(err), http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), nil)
}

func InternalServerError(err interface{}) *HttpResponse {
	return NewHttpResponse(formatError(err), http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
}

func Unauthorized(err interface{}) *HttpResponse {
	return NewHttpResponse(formatError(err), http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
}

func Forbidden(err interface{}) *HttpResponse {
	return NewHttpResponse(formatError(err), http.StatusForbidden, http.StatusText(http.StatusForbidden), nil)
}
