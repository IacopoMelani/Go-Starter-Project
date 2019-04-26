package boot

import (
	"Go-Starter-Project/config"
	"Go-Starter-Project/routes"
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

	wg.Add(2)

	go func() {
		defer wg.Done()
		config.GetInstance()
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
