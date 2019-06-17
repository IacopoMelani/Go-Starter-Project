package durationdata

import (
	"testing"
	"time"
)

func TestGetUserData(t *testing.T) {

	d := GetUsersData(1)

	time.Sleep(2 * time.Second)

	d.StopDaemon()

	d.SetContent(12, 3)

	d.GetContent()

}
