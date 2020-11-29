package controllers

import (
	"github.com/labstack/echo/middleware"
)

// Response - Define a standard struct response
type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

// InitCustomHandler - initialize custom error code and message
func InitCustomHandler() {

	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "token missing or expired"
}

// SetContent - Set content of response
func (r *Response) SetContent(content interface{}) {
	r.Content = content
}

// SetMessage - Set message of response
func (r *Response) SetMessage(message string) {
	r.Message = message
}

// SetStatus - Set the status code of response
func (r *Response) SetStatus(status int) {
	r.Status = status
}

// SetSuccess - Set the result of response
func (r *Response) SetSuccess(success bool) {
	r.Success = success
}
