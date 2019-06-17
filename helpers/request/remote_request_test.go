package request

import (
	"io"
	"net/http"
	"testing"
)

type exampleRemoteData struct{}

// EncodeQueryString - Si occupa di aggiungere i paratri dell'header alla request
func (u exampleRemoteData) EncodeQueryString(req *http.Request) {}

// GetBody - Restituisce il corpo della request
func (u exampleRemoteData) GetBody() io.Reader {
	return nil
}

// GetMethod - Restituisce il metodo della richiesta remota
func (u exampleRemoteData) GetMethod() string {
	return "GET"
}

// GetURL - Restituisce la url della richiesta remota
func (u exampleRemoteData) GetURL() string {
	return "https://randomuser.me/api/"
}

func TestRemoteData(t *testing.T) {

	tr := exampleRemoteData{}

	content, err := GetRemoteData(tr)
	if err != nil {
		t.Fatal(err.Error())
	}

	if content == nil {
		t.Fatal("Risposta vuota")
	}

}
