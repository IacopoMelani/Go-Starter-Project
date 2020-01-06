package boot

import (
	"os"
	"sync"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/log"
	"github.com/op/go-logging"

	durationmodel "github.com/IacopoMelani/Go-Starter-Project/models/duration_data"

	durationdata "github.com/IacopoMelani/Go-Starter-Project/pkg/models/duration_data"

	"github.com/IacopoMelani/Go-Starter-Project/controllers"

	"github.com/IacopoMelani/Go-Starter-Project/config"
	"github.com/IacopoMelani/Go-Starter-Project/routes"

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

	wg.Add(5)

	go func() {
		defer wg.Done()
		config.GetInstance()
	}()

	go func() {
		defer wg.Done()
		durationdata.RegisterInitDurationData(durationmodel.GetUsersData)
		durationdata.InitDurationData()
	}()

	go func() {
		defer wg.Done()
		e = echo.New()
		e.Use(middleware.Recover())
		e.Use(middleware.Logger())
		initEchoRoutes(e)

	}()

	var file *os.File
	defer file.Close()
	go func(file *os.File) {
		defer wg.Done()
		if _, err := os.Stat("./log"); os.IsNotExist(err) {
			os.Mkdir("./log", os.ModePerm)
		}
		file, err := os.OpenFile("./log/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		log.NewLogBackend(os.Stdout, "", 0, logging.DEBUG, log.DefaultLogFormatter)
		log.NewLogBackend(file, "", 0, logging.WARNING, log.VerboseLogFilePathFormatter)
		log.Init()
	}(file)

	go func() {
		defer wg.Done()
		controllers.InitCustomHandler()
	}()

	wg.Wait()

	config := config.GetInstance()

	logger := log.GetLogger()
	logger.Info("Applicazione avviata!")

	e.Logger.Fatal(e.Start(config.AppPort))
}
