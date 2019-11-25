package rmanager

import (
	"testing"
	"time"

	durationmodel "github.com/IacopoMelani/Go-Starter-Project/models/duration_data"
)

func TestRequestManager(t *testing.T) {

	rm := GetRequestManager()

	for i := 0; i < 10; i++ {

		go func() {

			u := durationmodel.UserRemoteData{}

			res, err := rm.AddRequest(u)

			select {
			case <-res:
				t.Log("response")
			case err := <-err:
				t.Log(err.Error())
			}

		}()
	}

	time.Sleep(5 * time.Second)
}
