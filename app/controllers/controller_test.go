package controllers

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestStandardResponse(t *testing.T) {

	e := echo.New()
	c := e.NewContext(nil, nil)

	response := FailedResponse(c, 1, "Error", echo.Map{"error": "this is an error"})

	if response.GetCode() != 1 || response.GetMessage() != "Error" || response.GetSuccess() || response.GetContent() == nil {
		t.Fatal("Error failed response")
	}

	response = SuccessResponse(c, nil)
	if response.GetCode() != 0 || response.GetMessage() != ResponseMessageOk || !response.GetSuccess() || response.GetContent() != nil {
		t.Fatal("Error failed response")
	}
}
