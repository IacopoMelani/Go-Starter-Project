package request

import (
	"io"
	"net/http"
	"testing"
)

// exampleRemoteData - An example of remote data with success response
type exampleRemoteData struct{}

// EncodeQueryString -
func (u exampleRemoteData) EncodeQueryString(req *http.Request) {}

// GetBody -
func (u exampleRemoteData) GetBody() io.Reader {
	return nil
}

// GetMethod -
func (u exampleRemoteData) GetMethod() string {
	return "GET"
}

// GetURL - Restituisce la url della richiesta remota
func (u exampleRemoteData) GetURL() string {
	return "https://randomuser.me/api/"
}

// exampleRemoteDataEmpty -An example of remote data with empty fields passed to pkg http
type exampleRemoteDataEmpty struct{}

// EncodeQueryString -
func (u exampleRemoteDataEmpty) EncodeQueryString(req *http.Request) {}

// GetBody -
func (u exampleRemoteDataEmpty) GetBody() io.Reader {
	return nil
}

// GetMethod -
func (u exampleRemoteDataEmpty) GetMethod() string {
	return "GET"
}

// GetURL -
func (u exampleRemoteDataEmpty) GetURL() string {
	return ""
}

// exampleRemoteDataErrorURL - An example of remote data with wrong method and url
type exampleRemoteDataErrorURL struct{}

// EncodeQueryString -
func (u exampleRemoteDataErrorURL) EncodeQueryString(req *http.Request) {}

// GetBody -
func (u exampleRemoteDataErrorURL) GetBody() io.Reader {
	return nil
}

// GetMethod -
func (u exampleRemoteDataErrorURL) GetMethod() string {
	return "435_34543"
}

// GetURL -
func (u exampleRemoteDataErrorURL) GetURL() string {
	return "--:233::"
}

// exampleRemoteDataErrorParseJSON - An example of remote data with error during parsing JSON
type exampleRemoteDataErrorParseJSON struct{}

// EncodeQueryString -
func (u exampleRemoteDataErrorParseJSON) EncodeQueryString(req *http.Request) {}

// GetBody -
func (u exampleRemoteDataErrorParseJSON) GetBody() io.Reader {
	return nil
}

// GetMethod -
func (u exampleRemoteDataErrorParseJSON) GetMethod() string {
	return "GET"
}

// GetURL -
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
