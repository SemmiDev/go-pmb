package model

type Meta struct {
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
}

type Response struct {
	Meta *Meta       `json:"meta"`
	Data interface{} `json:"data"`
}

func APIResponse(message interface{}, code int, status string, data interface{}) *Response {
	return &Response{
		Meta: &Meta{
			Message: message,
			Code:    code,
			Status:  status,
		},
		Data: data,
	}
}
