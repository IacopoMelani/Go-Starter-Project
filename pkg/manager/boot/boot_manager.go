package bootmanager

import (
	"sync"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"

	durationdata "github.com/IacopoMelani/Go-Starter-Project/pkg/models/duration_data"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
)

// BManager - Generalize a boot manager
type BManager interface {
	RegisterDDataProc(func() *durationdata.DurationData)
	RegisterProc(func())
	RegisterEchoRoutes(func(e *echo.Echo))
	SetAdditionalConnection(name string, driver string, conn string)
	SetAppPort(port string)
	SetConnectionSting(conn string)
	SetDriverSQL(driver string)
	StartApp()
	UseEchoLogger()
	UseEchoRecover()
}

// boot - Manage the app boot,
// implements BManager to impone the user to use the method GetBootManager to retrive the instance of boot
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

// define the min cap for WaitGroup in the boot phase
const minWaitGroupCap = 2

var (
	bmanger *boot
	onceB   = sync.Once{}
)

// initDbConnection - Initializes connection to database
func (b *boot) initDbConnection() {
	db.InitConnection(b.driverSQL, b.connection)
}

// initDurationData - Initializes all DurationData registered before the the boot phase, to register a component use RegisterDDataProc
func (b *boot) initDurationData() {
	for _, d := range b.durationDataRegistered {
		durationdata.RegisterInitDurationData(d)
	}
	durationdata.InitDurationData()
}

// GetBootManager - Returns the one BManager
func GetBootManager() BManager {
	onceB.Do(func() {
		bmanger = new(boot)
		bmanger.e = echo.New()
	})
	return bmanger
}

// startProc - Exec all generics procedure define before the boot phase, to register a proc use RegisterProc
func (b *boot) startProc() {
	for _, f := range b.procRegistered {
		go func(f func()) {
			defer b.wg.Done()
			f()
		}(f)
	}
}

// RegisterDDataProc - Register a DurationData component
func (b *boot) RegisterDDataProc(f func() *durationdata.DurationData) {

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.runnig {
		return
	}

	b.durationDataRegistered = append(b.durationDataRegistered, f)
}

// RegisterProc - Register a generics procedure, you can define whatever func you want, this procedure will start in the boot phase
func (b *boot) RegisterProc(f func()) {

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.runnig {
		return
	}

	b.procRegistered = append(b.procRegistered, f)
}

// RegisterEchoRoutes - Register a func that register an Echo routes, this func will called in the boot phase
func (b *boot) RegisterEchoRoutes(f func(e *echo.Echo)) {

	b.mu.Lock()
	defer b.mu.Unlock()

	f(b.e)
}

// SetAdditionalConnection -
func (b *boot) SetAdditionalConnection(name string, driver string, conn string) {}

// SetAppPort - Set the listening port
func (b *boot) SetAppPort(appPort string) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.appPort = appPort
}

// SetConnectionSting - Set the connection string
func (b *boot) SetConnectionSting(conn string) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.connection = conn
}

// SetDriverSQL - Set the SQL driver
func (b *boot) SetDriverSQL(driver string) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.driverSQL = driver
}

// StartApp - Try to boot the app, exec all registered procedures and finaly start listeing on the designed port
func (b *boot) StartApp() {

	b.wg.Add(minWaitGroupCap + len(b.procRegistered))

	b.mu.Lock()
	b.runnig = true
	b.mu.Unlock()

	if b.connection != "" && b.driverSQL != "" {

		go func() {
			defer b.wg.Done()
			b.initDbConnection()
		}()
	} else {
		b.wg.Done()
	}

	go func() {
		defer b.wg.Done()
		go b.initDurationData()
	}()

	b.startProc()

	b.wg.Wait()

	// start echo

	b.e.Logger.Fatal(b.e.Start(b.appPort))
}

// UseEchoLogger - Declare that Echo will uses internal logger
func (b *boot) UseEchoLogger() {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.e.Use(middleware.Logger())
}

// UseEchoRecover - Declare that Echo uses the internal recover
func (b *boot) UseEchoRecover() {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.e.Use(middleware.Recover())
}
