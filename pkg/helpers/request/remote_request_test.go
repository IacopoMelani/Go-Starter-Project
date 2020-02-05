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

type exampleRemoteDataEmpty struct{}

// EncodeQueryString - Si occupa di aggiungere i paratri dell'header alla request
func (u exampleRemoteDataEmpty) EncodeQueryString(req *http.Request) {}

// GetBody - Restituisce il corpo della request
func (u exampleRemoteDataEmpty) GetBody() io.Reader {
	return nil
}

// GetMethod - Restituisce il metodo della richiesta remota
func (u exampleRemoteDataEmpty) GetMethod() string {
	return "GET"
}

// GetURL - Restituisce la url della richiesta remota
func (u exampleRemoteDataEmpty) GetURL() string {
	return ""
}

type exampleRemoteDataErrorURL struct{}

// EncodeQueryString - Si occupa di aggiungere i paratri dell'header alla request
func (u exampleRemoteDataErrorURL) EncodeQueryString(req *http.Request) {}

// GetBody - Restituisce il corpo della request
func (u exampleRemoteDataErrorURL) GetBody() io.Reader {
	return nil
}

// GetMethod - Restituisce il metodo della richiesta remota
func (u exampleRemoteDataErrorURL) GetMethod() string {
	return "435_34543"
}

// GetURL - Restituisce la url della richiesta remota
func (u exampleRemoteDataErrorURL) GetURL() string {
	return "--:233::"
}

type exampleRemoteDataErrorParseJSON struct{}

// EncodeQueryString - Si occupa di aggiungere i paratri dell'header alla request
func (u exampleRemoteDataErrorParseJSON) EncodeQueryString(req *http.Request) {}

// GetBody - Restituisce il corpo della request
func (u exampleRemoteDataErrorParseJSON) GetBody() io.Reader {
	return nil
}

// GetMethod - Restituisce il metodo della richiesta remota
func (u exampleRemoteDataErrorParseJSON) GetMethod() string {
	return "GET"
}

// GetURL - Restituisce la url della richiesta remota
func (u exampleRemoteDataErrorParseJSON) GetURL() string {
	return "http://www.mocky.io/v2/5e3b4cd32f00006be356c9f7"
}

func TestRemoteData(t *testing.T) {

	r := exampleRemoteData{}

	content, err := GetRemoteData(r)
	if err != nil {
		t.Fatal(err.Error())
	}

	if content == nil {
		t.Fatal("Risposta vuota")
	}

	r2 := exampleRemoteDataErrorURL{}
	_, err = GetRemoteData(r2)
	if err == nil {
		t.Fatal("Dovrebbe essere Error")
	}

	r3 := exampleRemoteDataEmpty{}
	_, err = GetRemoteData(r3)
	if err == nil {
		t.Fatal(err.Error())
	}

	r4 := exampleRemoteDataErrorParseJSON{}
	_, err = GetRemoteData(r4)
	if err == nil {
		t.Fatal(err.Error())
	}
}
