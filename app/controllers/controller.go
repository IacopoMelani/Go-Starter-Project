package controllers

import (
	"github.com/labstack/echo/v4"
)

// MARK: Consts

// constants
const (
	ResponseMessageOk = "ok!"
)

// MARK: Response interface

// Response - Generalizes a resposne for web/api
type Response interface {
	GetCode() int
	GetSuccess() bool
	GetMessage() string
	GetContent() echo.Map
}

// NewResponse - Returns a new resposne based on the echo context
func NewResponse(c echo.Context, success bool, code int, message string, content echo.Map) Response {
	// no custom context defined, returns basic response api
	return NewResponseAPI(success, code, message, content)
}

// MARK: Response API, constructor and implementation

// ResponseData - Defines response data
type ResponseData struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ResponseAPI - Definea a standard struct response
type ResponseAPI struct {
	Data    ResponseData `json:"data"`
	Content echo.Map     `json:"content,omitempty"`
}

// NewResponseAPI - Returns a new response api
func NewResponseAPI(success bool, code int, message string, content echo.Map) ResponseAPI {
	return ResponseAPI{
		Data: ResponseData{
			Code:    code,
			Success: success,
			Message: message,
		},
		Content: content,
	}
}

// GetCode - Returns the api code response
func (r ResponseAPI) GetCode() int {
	return r.Data.Code
}

// GetSuccess - Returns the outcome of the request
func (r ResponseAPI) GetSuccess() bool {
	return r.Data.Success
}

// GetMessage - Returns the response message
func (r ResponseAPI) GetMessage() string {
	return r.Data.Message
}

// GetContent - Returns the content of the response
func (r ResponseAPI) GetContent() echo.Map {
	return r.Content
}

// MARK Exported funcs

// FailedResponse - Returns a failed Response
func FailedResponse(c echo.Context, code int, message string, content echo.Map) Response {
	return NewResponse(c, false, code, message, content)
}

// SuccessResponse - Returns a success Response
func SuccessResponse(c echo.Context, content echo.Map) Response {
	return NewResponse(c, true, 0, ResponseMessageOk, content)
}
