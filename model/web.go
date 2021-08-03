package model

type WebResponse struct {
	Code         int         `json:"code"`
	Status       string      `json:"status"`
	Error        bool        `json:"error"`
	ErrorMessage interface{} `json:"error_message"`
	Data         interface{} `json:"data"`
}
