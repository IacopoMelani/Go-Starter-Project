package controllers

import (
	"testing"
)

func TestStandardResponse(t *testing.T) {

	response := new(Response)

	response.SetContent(1)
	response.SetMessage("ok")
	response.SetStatus(0)
	response.SetSuccess(true)

	if response.Content != 1 || response.Message != "ok" || response.Status != 0 || response.Success != true {
		t.Fatal("Errrore nella response standard")
	}

	InitCustomHandler()
}
