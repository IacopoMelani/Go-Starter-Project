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

// Response - Interface per generalizzare una response api/web
type Response interface {
	GetCode() int
	GetSuccess() bool
	GetMessage() string
	GetContent() echo.Map
}

// NewResponse - Restituisce una nuova response generica in base al context passato
func NewResponse(c echo.Context, success bool, code int, message string, content echo.Map) Response {
	return NewResponseAPI(success, code, message, content)
}

// MARK: Response API, constructor and implementation

// ResponseData - Definisce la struct della response.data
type ResponseData struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ResponseAPI - Define a standard struct response
type ResponseAPI struct {
	Data    ResponseData `json:"data"`
	Content echo.Map     `json:"content,omitempty"`
}

// NewResponseAPI - Restituisce una response
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

// GetCode - Restituisce il codice della response api
func (r ResponseAPI) GetCode() int {
	return r.Data.Code
}

// GetSuccess - Restituisce l'esito della response api
func (r ResponseAPI) GetSuccess() bool {
	return r.Data.Success
}

// GetMessage - Restituisce il message della response api
func (r ResponseAPI) GetMessage() string {
	return r.Data.Message
}

// GetContent - Restituisce l'eventuale content della response api
func (r ResponseAPI) GetContent() echo.Map {
	return r.Content
}

// MARK Exported funcs

// FailedResponse - Restitusce una risposta failed
func FailedResponse(c echo.Context, code int, message string, content echo.Map) Response {
	return NewResponse(c, false, code, message, content)
}

// SuccessResponse - Restituisce una risposta success
func SuccessResponse(c echo.Context, content echo.Map) Response {
	return NewResponse(c, true, 0, ResponseMessageOk, content)
}
