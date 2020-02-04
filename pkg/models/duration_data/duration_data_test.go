package durationdata

import (
	"io"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/request"
	"github.com/subosito/gotenv"
)

// DurationDataTest - Definisce una struct che implementa RemoteData
type DurationDataTest struct{}

var ddt *DurationData
var onceUser sync.Once

// GeDurationDataTest - Restituisce l'istanza di DurantionData relativo agli utenti
func GeDurationDataTest() *DurationData {
	onceUser.Do(func() {
		ddt = new(DurationData)
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

func TestDurationData(t *testing.T) {

	if err := gotenv.Load("./../../../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}

	RegisterInitDurationData(GeDurationDataTest)

	InitDurationData()
	time.Sleep(2 * time.Second)

	d := GeDurationDataTest()

	d.StopDaemon()

	d.GetSafeContent()
}
