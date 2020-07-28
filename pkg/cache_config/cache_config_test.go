package cacheconf

import (
	"sync"
	"testing"

	"github.com/subosito/gotenv"
)

// CacheConfigTest - Example CacheConfig test struct
type CacheConfigTest struct {
	DefaultCacheConfig
	UserTimeToRefresh int `config:"USER_TIME_TO_REFRESH"`
}

var (
	cct  *CacheConfigTest
	once sync.Once
)

// GetInstance - Return the one struct for tes
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

	if !Debug() {
		t.Error("Errore: assicurarsi che il flag debug sia a true")
	}
}
