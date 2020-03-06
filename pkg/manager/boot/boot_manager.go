package bootmanger

import (
	"sync"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"

	durationdata "github.com/IacopoMelani/Go-Starter-Project/pkg/models/duration_data"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
)

// BManager - Interfaccia che generalizza il boot manager
type BManager interface {
	StartApp()
	RegisterDDataProc(func() *durationdata.DurationData)
	RegisterProc(func())
	RegisterEchoRoutes(func(e *echo.Echo))
	SetAppPort(string)
	SetDriverSQL(string)
	SetConnectionSting(string)
	UseEchoLogger()
	UseEchoRecover()
}

// boot - Struct che si occupa di gestire il boot dell'app
type boot struct {
	appPort                string
	connection             string
	driverSQL              string
	durationDataRegistered []func() *durationdata.DurationData
	e                      *echo.Echo
	runnig                 bool
	mu                     sync.Mutex
	procRegistered         []func()
	wg                     sync.WaitGroup
}

const minWaitGroupCap = 2

var (
	bmanger *boot
	ok      chan bool
	onceB   = sync.Once{}
)

func (b *boot) initDbConnection() {
	db.InitConnection(b.driverSQL, b.connection)
}

func (b *boot) initDurationData() {
	for _, d := range b.durationDataRegistered {
		durationdata.RegisterInitDurationData(d)
	}
	durationdata.InitDurationData()
}

// GetBootManager - Restituisce il boot manager dell'applicazione
func GetBootManager() BManager {
	onceB.Do(func() {
		bmanger = new(boot)
		bmanger.e = echo.New()
	})
	return bmanger
}

func (b *boot) StartApp() {

	b.wg.Add(minWaitGroupCap + len(b.procRegistered))

	b.runnig = true

	go func() {
		defer b.wg.Done()
		b.initDbConnection()
	}()

	go func() {
		defer b.wg.Done()
		go b.initDurationData()
	}()

	b.StartProc()

	b.wg.Wait()

	// start echo

	b.e.Logger.Fatal(b.e.Start(b.appPort))

}

func (b *boot) RegisterDDataProc(f func() *durationdata.DurationData) {

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.runnig {
		return
	}

	b.durationDataRegistered = append(b.durationDataRegistered, f)
}

func (b *boot) RegisterProc(f func()) {

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.runnig {
		return
	}

	b.procRegistered = append(b.procRegistered, f)
}

func (b *boot) RegisterEchoRoutes(f func(e *echo.Echo)) {

	b.mu.Lock()
	defer b.mu.Unlock()

	f(b.e)
}

func (b *boot) SetAppPort(appPort string) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.appPort = appPort
}

func (b *boot) SetDriverSQL(driver string) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.driverSQL = driver
}

func (b *boot) SetConnectionSting(conn string) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.connection = conn
}

func (b *boot) StartProc() {
	for _, f := range b.procRegistered {
		go func(f func()) {
			defer b.wg.Done()
			f()
		}(f)
	}
}

func (b *boot) UseEchoLogger() {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.e.Use(middleware.Logger())
}

func (b *boot) UseEchoRecover() {

	b.mu.Lock()
	b.mu.Unlock()

	b.e.Use(middleware.Recover())
}
