package boot

import (
	"github.com/IacopoMelani/Go-Starter-Project/models/duration_data"
	"testing"
	"time"

	"github.com/subosito/gotenv"

	"github.com/IacopoMelani/Go-Starter-Project/config"
)

func TestInitServer(t *testing.T) {

	gotenv.Load("../.env")

	go InitServer()

	config := config.GetInstance()

	time.Sleep(1000 * time.Millisecond)

	if config.StringConnection == "" || durationdata.GetUsersData().GetSafeContent() == nil {
		t.Fatal("Errore durante l'avvio del server")
	}

}
