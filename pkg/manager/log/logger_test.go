package log

import (
	"os"
	"testing"
	"time"

	"github.com/op/go-logging"
)

var (
	testLogDirName  = "log-test"
	testLogFileName = "info-test.log"
)

func TestLogger(t *testing.T) {

	Now = func() time.Time { return time.Now().Add(24 * time.Hour) }

	if _, err := os.Stat(testLogDirName); os.IsNotExist(err) {
		err = os.Mkdir(testLogDirName, os.ModePerm)
		if err != nil {
			t.Fatal("Errore durante la creazione della cartella di test")
		}
	}

	file, err := NewRotateFile(testLogDirName+"/"+testLogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644, true)
	if err != nil {
		panic(err)
	}

	NewLogBackend(os.Stderr, "", 0, logging.DEBUG, nil)
	NewLogBackend(file, "", 0, logging.WARNING, VerboseLogFilePathFormatter)
	NewLogBackend(os.Stdout, "", 0, logging.INFO, LowVerboseLogFormatter)
	Init("Test-App")

	testLogger := GetLogger()
	testLogger.Info("Test info")
	testLogger.Debug("Test debug")
	testLogger.Warning("Test warning")
}
