package boot

import (
	"os"

	bootmanger "github.com/IacopoMelani/Go-Starter-Project/pkg/manager/boot"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/log"
	"github.com/op/go-logging"

	durationmodel "github.com/IacopoMelani/Go-Starter-Project/models/duration_data"

	"github.com/IacopoMelani/Go-Starter-Project/config"
	"github.com/IacopoMelani/Go-Starter-Project/routes"

	"github.com/labstack/echo"
)

var e *echo.Echo

func initEchoRoutes(e *echo.Echo) {

	routes.InitGetRoutes(e)
	routes.InitPostRoutes(e)
}

// InitServer - Si occupa di lanciare l'applicazione con tutte le dovute operazioni iniziali
func InitServer() {

	config := config.GetInstance()

	bm := bootmanger.GetBootManager()

	bm.SetAppPort(config.AppPort)
	bm.SetConnectionSting(config.StringConnection)
	bm.SetDriverSQL("mysql")

	bm.RegisterEchoRoutes(initEchoRoutes)

	bm.UseEchoLogger()
	bm.UseEchoRecover()

	bm.RegisterDDataProc(durationmodel.GetUsersData)

	bm.RegisterProc(func() {

		var file *os.File

		file, err := os.OpenFile("./log/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		log.NewLogBackend(os.Stdout, "", 0, logging.DEBUG, log.DefaultLogFormatter)
		log.NewLogBackend(file, "", 0, logging.WARNING, log.VerboseLogFilePathFormatter)
		log.Init(config.AppName)

		logger := log.GetLogger()

		logger.Debug("App avviata")
	})

	bm.StartApp()
}
