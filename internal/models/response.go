package models

type Meta struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type Response struct {
	Meta *Meta       `json:"meta"`
	Data interface{} `json:"data"`
}

func APIResponse(message interface{}, code int, status string, data interface{}) *Response {
	return &Response{
		Meta: &Meta{
			Code:    code,
			Status:  status,
			Message: message,
		},
		Data: data,
	}
}
