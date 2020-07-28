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

// DurationDataTest -
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

// EncodeQueryString -
func (u DurationDataTest) EncodeQueryString(req *http.Request) {}

// GetBody -
func (u DurationDataTest) GetBody() io.Reader {
	return nil
}

// GetMethod -
func (u DurationDataTest) GetMethod() string {
	return "GET"
}

// GetURL -
func (u DurationDataTest) GetURL() string {
	return "https://randomuser.me/api/"
}

// HandlerData -
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
	port := ":8889"

	bm := GetBootManager()

	bm.SetAppPort(port)
	bm.SetConnectionSting(conn)
	bm.SetDriverSQL(os.Getenv("SQL_DRIVER"))

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
