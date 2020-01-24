package boot

import (
	"testing"

	"github.com/subosito/gotenv"

	"github.com/IacopoMelani/Go-Starter-Project/config"
)

func TestInitServer(t *testing.T) {

	if err := gotenv.Load("../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}

	go InitServer()

	config := config.GetInstance()

	if config.StringConnection == "" {
		t.Fatal("Errore durante l'avvio del server")
	}

}
