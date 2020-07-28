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

// DurationDataTest - Example
type DurationDataTest struct{}

var ddt *DurationData
var onceDD sync.Once

// GeDurationDataTest - Example
func GeDurationDataTest() *DurationData {
	onceDD.Do(func() {
		ddt = new(DurationData)
		ddt.SetDurationDataInterface(DurationDataTest{})
		ddt.SetTimeToRefresh(1)
		ddt.Daemon()
	})
	return ddt
}

// EncodeQueryString - Example
func (u DurationDataTest) EncodeQueryString(req *http.Request) {}

// GetBody - Example
func (u DurationDataTest) GetBody() io.Reader {
	return nil
}

// GetMethod - Example
func (u DurationDataTest) GetMethod() string {
	return "GET"
}

// GetURL - Example
func (u DurationDataTest) GetURL() string {
	return "https://randomuser.me/api/"
}

// HandlerData - Example
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
