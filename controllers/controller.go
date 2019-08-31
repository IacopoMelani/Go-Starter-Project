package controllers

import (
	"github.com/labstack/echo/middleware"
)

// Response - struttura di una response standard
type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

// InitCustomHandler - inizializza i codici di errore e messaggi per custom response
func InitCustomHandler() {

	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "token missing or expired"
}

// SetContent - Imposta content della risposta
func (r *Response) SetContent(content interface{}) {
	r.Content = content
}

// SetMessage - imposta il messaggio della response
func (r *Response) SetMessage(message string) {
	r.Message = message
}

// SetStatus - imposta lo status della response
func (r *Response) SetStatus(status int) {
	r.Status = status
}

// SetSuccess - imposta il success della response
func (r *Response) SetSuccess(success bool) {
	r.Success = success
}
