package log

import (
	"os"
	"testing"
	"time"
)

func TestRotateFile(t *testing.T) {

	Now = func() time.Time { return time.Now().Add(24 * time.Hour) }

	_, err := NewRotateFile("./log-test-error/info-test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666, true)
	if err == nil {
		t.Fatal("No errors occured")
	}

	file, err := NewRotateFile("./log-test/info-test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666, true)
	if err != nil {
		t.Fatal(err)
	}

	file.closeFile()

	file.writeRaw([]byte("First log raw"))

	file.Write([]byte("First log\n"))

}
