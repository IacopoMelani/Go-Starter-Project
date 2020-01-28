package boot

import (
	"testing"
	"time"

	"github.com/subosito/gotenv"

	"github.com/IacopoMelani/Go-Starter-Project/config"
)

func TestInitServer(t *testing.T) {

	if err := gotenv.Load("../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}

	go InitServer()

	time.Sleep(1 * time.Second)

	config := config.GetInstance()

	if config.StringConnection == "" {
		t.Fatal("Errore durante l'avvio del server")
	}

}
