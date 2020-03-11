package cacheconf

import (
	"sync"
	"testing"

	"github.com/subosito/gotenv"
)

// CacheConfigTest - struttura dove immagazzinare le configurazioni
type CacheConfigTest struct {
	DefaultCacheConfig
	StringConnection  string `config:"APP_NAME"`
	AppPort           string `config:"STRING_CONNECTION"`
	UserTimeToRefresh int    `config:"APP_PORT"`
}

var (
	cct  *CacheConfigTest
	once sync.Once
)

// GetInstance - restituisce l'unica istanza della struttura contenente le configurazioni
func GetInstance() *CacheConfigTest {
	once.Do(func() {
		cct = &CacheConfigTest{}
		LoadEnvConfig(cct)
	})
	return cct
}

func TestCacheConfigBoot(t *testing.T) {

	if err := gotenv.Load("../../.env"); err != nil {
		t.Fatal("Errore durante il caricamento della configurazione")
	}

	config := GetInstance()

	if config.StringConnection == "" {
		t.Error("Errore: variabili d'ambiente non caricate correttamente")
	}

	GetCurrentConfig()
}
