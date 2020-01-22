package config

import (
	"testing"

	"github.com/subosito/gotenv"
)

func TestCacheConfigBoot(t *testing.T) {

	if err := gotenv.Load("../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}

	config := GetInstance()

	if config.StringConnection == "" {
		t.Error("Errore: variabili d'ambiente non caricate correttamente")
	}

}
