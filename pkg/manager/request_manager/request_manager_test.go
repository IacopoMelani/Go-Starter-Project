package rmanager

import (
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/request"
)

// RemoteDataTest - Definisce una struct che implementa RemoteData
type RemoteDataTest struct{}

// EncodeQueryString - Si occupa di aggiungere i paratri dell'header alla request
func (u RemoteDataTest) EncodeQueryString(req *http.Request) {}

// GetBody - Restituisce il corpo della request
func (u RemoteDataTest) GetBody() io.Reader {
	return nil
}

// GetMethod - Restituisce il metodo della richiesta remota
func (u RemoteDataTest) GetMethod() string {
	return "GET"
}

// GetURL - Restituisce la url della richiesta remota
func (u RemoteDataTest) GetURL() string {
	return "https://randomuser.me/api/"
}

// HandlerData - Si occupa di eseguire la funzione di handler per ricevere i dati
func (u RemoteDataTest) HandlerData() (interface{}, error) {
	content, err := request.GetRemoteData(u)
	return content, err
}

func TestRequestManager(t *testing.T) {

	rm := GetRequestManager()

	for i := 0; i < 10; i++ {

		go func() {

			u := RemoteDataTest{}

			res, err := rm.AddRequest(u)

			select {
			case <-res:
				t.Log("response")
			case err := <-err:
				t.Log(err.Error())
			}

		}()
	}
	time.Sleep(4 * time.Second)
}
