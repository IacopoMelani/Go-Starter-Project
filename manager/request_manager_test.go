package rmanager

import (
	durationdata "github.com/IacopoMelani/Go-Starter-Project/models/duration_data"
	"testing"
	"time"
)

func TestRequestManager(t *testing.T) {

	rm := GetRequestManager()

	u := durationdata.UserRemoteData{}

	for i := 0; i < 10; i++ {

		go func() {

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
