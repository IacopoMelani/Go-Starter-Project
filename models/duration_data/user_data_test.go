package durationdata

import (
	"testing"
	"time"

	"github.com/subosito/gotenv"
)

func TestGetUserData(t *testing.T) {

	gotenv.Load("./../../.env.test")

	d := GetUsersData()

	time.Sleep(2 * time.Second)

	d.StopDaemon()

	d.SetContent(12, 3)

	d.GetContent()

}
