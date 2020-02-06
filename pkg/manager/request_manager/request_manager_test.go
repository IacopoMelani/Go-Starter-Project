package rmanager

import (
	"io"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/request"
	durationdata "github.com/IacopoMelani/Go-Starter-Project/pkg/models/duration_data"
)

// DurationDataTest - Definisce una struct che implementa RemoteData
type DurationDataTest struct{}

var ddt *durationdata.DurationData
var onceUser sync.Once

// GeDurationDataTest - Restituisce l'istanza di DurantionData relativo agli utenti
func GeDurationDataTest() *durationdata.DurationData {
	onceUser.Do(func() {
		ddt = new(durationdata.DurationData)
		ddt.SetDurationDataInterface(DurationDataTest{})
		ddt.SetTimeToRefresh(1)
		ddt.Daemon()
	})
	return ddt
}

// EncodeQueryString - Si occupa di aggiungere i paratri dell'header alla request
func (u DurationDataTest) EncodeQueryString(req *http.Request) {}

// GetBody - Restituisce il corpo della request
func (u DurationDataTest) GetBody() io.Reader {
	return nil
}

// GetMethod - Restituisce il metodo della richiesta remota
func (u DurationDataTest) GetMethod() string {
	return "GET"
}

// GetURL - Restituisce la url della richiesta remota
func (u DurationDataTest) GetURL() string {
	return "https://randomuser.me/api/"
}

// HandlerData - Si occupa di eseguire la funzione di handler per ricevere i dati
func (u DurationDataTest) HandlerData() (interface{}, error) {
	content, err := request.GetRemoteData(u)
	return content, err
}

func TestRequestManager(t *testing.T) {

	rm := GetRequestManager()

	for i := 0; i < 10; i++ {

		go func() {

			u := DurationDataTest{}

			res, err := rm.AddRequest(u)

			select {
			case <-res:
				t.Log("response")
			case err := <-err:
				t.Log(err.Error())
			}

		}()
	}

	time.Sleep(5 * time.Second)
}
