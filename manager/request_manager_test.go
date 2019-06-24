package rmanager

import (
	durationdata "Go-Starter-Project/models/duration_data"
	"testing"
	"time"
)

func TestRequestManager(t *testing.T) {

	rm := GetRequestManager()

	u := durationdata.UserRemoteData{}

	go func() {

		time.Sleep(600 * time.Millisecond)

		res, err := rm.AddRequest(u)

		select {
		case <-res:
			t.Log("response")
		case err := <-err:
			t.Log(err.Error())
		}

	}()

	go func() {
		rm.StopService(true)
	}()

	time.Sleep(5 * time.Second)
}
