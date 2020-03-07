package bootmanager

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

// initDbConnection - Si occupa di inizializzare la connessione
func (b *boot) initDbConnection() {
	db.InitConnection(b.driverSQL, b.connection)
}

// initDurationData - Si occupa di inizializzare tutti i componenti che utilizzano DurationData
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

// startProc - Si occupa di avviare tutte le procedure generiche durante l'avvio dell'applicazione
func (b *boot) startProc() {
	for _, f := range b.procRegistered {
		go func(f func()) {
			defer b.wg.Done()
			f()
		}(f)
	}
}

// StartApp - Si occupa di effettuare il boot dell'applicazione e avviare tutte le procedure registrate, al termine viene avviata l'applicazione sulla porta designata
func (b *boot) StartApp() {

	b.wg.Add(minWaitGroupCap + len(b.procRegistered))

	b.mu.Lock()
	b.runnig = true
	b.mu.Unlock()

	go func() {
		defer b.wg.Done()
		b.initDbConnection()
	}()

	go func() {
		defer b.wg.Done()
		go b.initDurationData()
	}()

	b.startProc()

	b.wg.Wait()

	// start echo

	b.e.Logger.Fatal(b.e.Start(b.appPort))
}

// RegisterDDataProc - Si occupa di registrare tutti i componenti che utilizzano duration Data
func (b *boot) RegisterDDataProc(f func() *durationdata.DurationData) {

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.runnig {
		return
	}

	b.durationDataRegistered = append(b.durationDataRegistered, f)
}

// RegisterProc - Si occupa di registrare tutte le generiche procedure da avviare al boot dell'applicazione
func (b *boot) RegisterProc(f func()) {

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.runnig {
		return
	}

	b.procRegistered = append(b.procRegistered, f)
}

// RegisterEchoRoutes - Si occupa di avviare la registrazione delle route con Echo, è compito dell'utilizzatore di assicurarsi di passare una func che effettui la registrazione delle rotte
func (b *boot) RegisterEchoRoutes(f func(e *echo.Echo)) {

	b.mu.Lock()
	defer b.mu.Unlock()

	f(b.e)
}

// SetAppPort - Si occupa di impostare la porta su chi gira l'applicazione
func (b *boot) SetAppPort(appPort string) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.appPort = appPort
}

// SetDriverSQL - imposta il driver SQL
func (b *boot) SetDriverSQL(driver string) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.driverSQL = driver
}

// SetConnectionSting - Imposta la stringa di connessione al database
func (b *boot) SetConnectionSting(conn string) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.connection = conn
}

// UseEchoLogger - Definisce se utilizzare il logger di richieste di Echo
func (b *boot) UseEchoLogger() {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.e.Use(middleware.Logger())
}

// UseEchoRecover . Definisce se utilizzare il recover di Echo in caso si verifichi un panic
func (b *boot) UseEchoRecover() {

	b.mu.Lock()
	b.mu.Unlock()

	b.e.Use(middleware.Recover())
}
