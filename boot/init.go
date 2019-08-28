package boot

import (
	"github.com/Go-Starter-Project/config"
	durationdata "github.com/Go-Starter-Project/models/duration_data"
	"github.com/Go-Starter-Project/routes"
	"sync"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

var e *echo.Echo

func initEchoRoutes(e *echo.Echo) {

	routes.InitGetRoutes(e)
	routes.InitPostRoutes(e)
}

// InitServer - Si occupa di lanciare l'applicazione con tutte le dovute operazioni iniziali
func InitServer() {

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		config.GetInstance()
	}()

	go func() {
		defer wg.Done()
		durationdata.InitDurationData()
	}()

	go func() {
		defer wg.Done()

		e = echo.New()
		e.Use(middleware.Recover())
		e.Use(middleware.Logger())

		initEchoRoutes(e)

	}()

	wg.Wait()

	config := config.GetInstance()

	e.Logger.Fatal(e.Start(config.AppPort))
}
