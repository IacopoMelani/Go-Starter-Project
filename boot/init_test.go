package boot

import (
	"github.com/subosito/gotenv"
	"testing"
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/config"
	durationdata "github.com/IacopoMelani/Go-Starter-Project/models/duration_data"
)

func TestInitServer(t *testing.T) {

	gotenv.Load("../.env")

	go InitServer()

	config := config.GetInstance()
	
	time.Sleep(1000 * time.Millisecond)

	if config.StringConnection == "" || durationdata.GetUsersData().Content == nil {
		t.Fatal("Errore durante l'avvio del server")
	}

}
