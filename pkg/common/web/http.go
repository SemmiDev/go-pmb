package web

type Response struct {
	Code         int         `json:"code"`
	Status       string      `json:"status"`
	ErrorMessage interface{} `json:"error_message"`
	Data         interface{} `json:"data"`
}

func OkResponse(payload interface{}) *Response {
	return &Response{
		Code:         200,
		Status:       "OK",
		ErrorMessage: nil,
		Data:         payload,
	}
}

func CreatedResponse(payload interface{}) *Response {
	return &Response{
		Code:         201,
		Status:       "Created",
		ErrorMessage: nil,
		Data:         payload,
	}
}

func BadRequestResponse(err interface{}) *Response {
	return &Response{
		Code:         400,
		Status:       "Bad Request",
		ErrorMessage: err,
		Data:         nil,
	}
}

func NotFoundResponse(err interface{}) *Response {
	return &Response{
		Code:         404,
		Status:       "Not Found",
		ErrorMessage: err,
		Data:         nil,
	}
}

func UnprocessableEntityResponse(err interface{}) *Response {
	return &Response{
		Code:         422,
		Status:       "Unprocessable Entity",
		ErrorMessage: err,
		Data:         nil,
	}
}

func ResponseInternalServerError(err interface{}) *Response {
	return &Response{
		Code:         500,
		Status:       "Internal Server Error",
		ErrorMessage: err,
		Data:         nil,
	}
}
