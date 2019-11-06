package boot

import (
	"testing"

	"github.com/subosito/gotenv"

	"github.com/IacopoMelani/Go-Starter-Project/config"
)

func TestInitServer(t *testing.T) {

	gotenv.Load("../.env")

	go InitServer()

	config := config.GetInstance()

	if config.StringConnection == "" {
		t.Fatal("Errore durante l'avvio del server")
	}

}
