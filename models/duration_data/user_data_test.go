package durationmodel

import (
	"testing"
	"time"

	durationdata "github.com/IacopoMelani/Go-Starter-Project/pkg/models/duration_data"

	"github.com/subosito/gotenv"
)

func TestGetUserData(t *testing.T) {

	gotenv.Load("./../../.env")

	durationdata.RegisterInitDurationData(GetUsersData)

	durationdata.InitDurationData()
	time.Sleep(2 * time.Second)

	d := GetUsersData()

	d.StopDaemon()

	d.SetContent(12, 3)

	d.GetContent()

	d.GetSafeContent()
}
