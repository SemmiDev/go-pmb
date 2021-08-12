package responses

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
	return NewHttpResponse(nil, http.StatusOK, "Ok", data)
}

func Created(data interface{}) *HttpResponse {
	return NewHttpResponse(nil, http.StatusCreated, "Created", data)
}

func BadRequest(err interface{}) *HttpResponse {
	return NewHttpResponse(formatError(err), http.StatusBadRequest, "Bad Request", nil)
}

func UnprocessableEntity(err interface{}) *HttpResponse {
	return NewHttpResponse(formatError(err), http.StatusUnprocessableEntity, "Unprocessable Entity", nil)
}

func InternalServerError(err interface{}) *HttpResponse {
	return NewHttpResponse(formatError(err), http.StatusInternalServerError, "Internal Server Error", nil)
}

func Unauthorized(err interface{}) *HttpResponse {
	return NewHttpResponse(formatError(err), http.StatusUnauthorized, "Unauthorized", nil)
}

func Forbidden(err interface{}) *HttpResponse {
	return NewHttpResponse(formatError(err), http.StatusForbidden, "Forbidded", nil)
}
