package durationdata

import (
	"testing"
	"time"

	"github.com/subosito/gotenv"
)

func TestGetUserData(t *testing.T) {

	gotenv.Load("./../../.env")

	RegisterInitDurationData(GetUsersData)

	InitDurationData()
	time.Sleep(2 * time.Second)

	d := GetUsersData()

	d.StopDaemon()

	d.SetContent(12, 3)

	d.GetContent()

	d.GetSafeContent()
}
