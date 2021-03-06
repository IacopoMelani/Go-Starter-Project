package boot

import (
	"os"

	bootmanager "github.com/IacopoMelani/Go-Starter-Project/pkg/manager/boot"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/log"
	"github.com/op/go-logging"

	durationmodel "github.com/IacopoMelani/Go-Starter-Project/app/models/duration_data"

	"github.com/IacopoMelani/Go-Starter-Project/config"
	"github.com/IacopoMelani/Go-Starter-Project/routes"

	"github.com/labstack/echo/v4"
)

// initEchoRoutes - Initialize all defined routes
func initEchoRoutes(e *echo.Echo) {

	routes.InitGetRoutes(e)
	routes.InitPostRoutes(e)
}

// InitServer - It takes care of launching the application with all the necessary initial operations
func InitServer() {

	config := config.GetInstance()

	bm := bootmanager.GetBootManager()

	bm.SetAppPort(config.AppPort)
	bm.SetConnectionSting(config.StringConnection)
	bm.SetDriverSQL(config.SQLDriver)

	bm.RegisterEchoRoutes(initEchoRoutes)

	if config.Debug {
		bm.UseEchoLogger()
	}

	bm.UseEchoRecover()

	bm.RegisterDDataProc(durationmodel.GetUsersData)

	bm.RegisterProc(func() {

		if _, err := os.Stat("./log"); os.IsNotExist(err) {
			if err = os.Mkdir("./log", os.ModePerm); err != nil {
				panic(err)
			}
		}

		file, err := log.NewRotateFile("log/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666, true)
		if err != nil {
			panic(err)
		}

		log.NewLogBackend(os.Stdout, "", 0, logging.DEBUG, log.DefaultLogFormatter)
		log.NewLogBackend(file, "", 0, logging.INFO, log.VerboseLogFilePathFormatter)
		log.Init(config.AppName)

		logger := log.GetLogger()

		if config.Debug {
			logger.Debug("App avviata")
			logger.Info("App avviata")
		}
	})

	bm.StartApp()
}
