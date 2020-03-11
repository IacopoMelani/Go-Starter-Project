package bootmanager

import (
	"io"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/labstack/echo"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/request"
	durationdata "github.com/IacopoMelani/Go-Starter-Project/pkg/models/duration_data"
	"github.com/subosito/gotenv"
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

func laodEnv() error {

	if err := gotenv.Load("./../../../.env"); err != nil {
		return err
	}

	return nil
}

func TestBootManager(t *testing.T) {

	if err := laodEnv(); err != nil {
		t.Fatal(err.Error())
	}

	conn := os.Getenv("STRING_CONNECTION")
	port := os.Getenv("APP_PORT")

	bm := GetBootManager()

	bm.SetAppPort(port)
	bm.SetConnectionSting(conn)
	bm.SetDriverSQL("mysql")

	bm.RegisterDDataProc(GeDurationDataTest)

	bm.RegisterProc(func() {
		println("hello world!")
	})

	bm.UseEchoLogger()
	bm.UseEchoRecover()
	bm.RegisterEchoRoutes(func(e *echo.Echo) {
		e.GET("/", func(c echo.Context) error { return nil })
	})

	go bm.StartApp()

	time.Sleep(1 * time.Second)

	bm.RegisterDDataProc(GeDurationDataTest)

	bm.RegisterProc(func() {
		println("hello world!")
	})
}
