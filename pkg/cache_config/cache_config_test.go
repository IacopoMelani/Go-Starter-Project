package cacheconf

import (
	"sync"
	"testing"

	"github.com/subosito/gotenv"
)

// CacheConfigTest - struttura dove immagazzinare le configurazioni
type CacheConfigTest struct {
	StringConnection  string
	AppPort           string
	UserTimeToRefresh int
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

// GetFieldMapper - Si occupa di restituire l'array di mappatura dell'env
func (c CacheConfigTest) GetFieldMapper() map[string]string {
	return map[string]string{
		"STRING_CONNECTION":    "StringConnection",
		"APP_PORT":             "AppPort",
		"USER_TIME_TO_REFRESH": "UserTimeToRefresh",
	}
}

func TestCacheConfigBoot(t *testing.T) {

	gotenv.Load("../../.env")

	config := GetInstance()

	if config.StringConnection == "" {
		t.Error("Errore: variabili d'ambiente non caricate correttamente")
	}

}
