package boot

import (
	durationdata "github.com/IacopoMelani/Go-Starter-Project/pkg/models/duration_data"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
)

// BManager -
type BManager struct {
	driverSQL              string
	connection             string
	durationDataRegistered []func() *durationdata.DurationData
}

func (b BManager) initDbConnection() {
	db.InitConnection(b.driverSQL, b.connection)
}

func (b BManager) initDurationData() {
	for _, d := range b.durationDataRegistered {
		durationdata.RegisterInitDurationData(d)
	}
	durationdata.InitDurationData()
}
