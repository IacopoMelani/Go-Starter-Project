package log

import (
	"os"
	"testing"

	"github.com/op/go-logging"
)

var (
	testLogDirName  = "log-test"
	testLogFileName = "info-test.log"
)

func TestLogger(t *testing.T) {

	if _, err := os.Stat(testLogDirName); os.IsNotExist(err) {
		err = os.Mkdir(testLogDirName, os.ModePerm)
		if err != nil {
			t.Fatal("Errore durante la creazione della cartella di test")
		}
	}
	file, err := os.OpenFile(testLogDirName+"/"+testLogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	NewLogBackend(os.Stderr, "", 0, logging.DEBUG, nil)
	NewLogBackend(file, "", 0, logging.WARNING, VerboseLogFilePathFormatter)
	NewLogBackend(os.Stdout, "", 0, logging.INFO, LowVerboseLogFormatter)
	Init("Test-App")

	testLogger := GetLogger()
	testLogger.Info("Test info")
	testLogger.Debug("Test debug")
	testLogger.Warning("Test warning")
}
