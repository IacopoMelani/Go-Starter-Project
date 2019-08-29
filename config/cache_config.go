package config

import (
	"os"
	"reflect"
	"strconv"
	"sync"
)

// CacheConfig - struttura dove immagazzinare le configurazioni
type CacheConfig struct {
	StringConnection  string
	AppPort           string
	UserTimeToRefresh int
}

var arrayEnvMapper = map[string]string{
	"STRING_CONNECTION":    "StringConnection",
	"APP_PORT":             "AppPort",
	"USER_TIME_TO_REFRESH": "UserTimeToRefresh",
}

var cacheConfig *CacheConfig
var once sync.Once

// GetInstance - restituisce l'unica istanza della struttura contenente le configurazioni
func GetInstance() *CacheConfig {
	once.Do(func() {
		cacheConfig = &CacheConfig{}
		cacheConfig.loadEnvConfig()

	})
	return cacheConfig
}

// loadEnvConfig - si occupa di caricare tutte le configurazioni dell'env nella struttura di configurazione
func (c *CacheConfig) loadEnvConfig() {
	for envName, StructName := range arrayEnvMapper {
		c.setField(StructName, os.Getenv(envName))
	}
}

// setField - si occupa di impostare  attrun campo averso la reflection, c รจ necessario sia un puntatore a una struttura
func (c *CacheConfig) setField(name string, value string) {

	rv := reflect.ValueOf(c)

	// Controllo se pointer a una struct
	if rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.Struct {

		// Prelevo i campi della struct
		rv = rv.Elem()

		// Controllo che il campo esista
		fv := rv.FieldByName(name)
		if fv.IsValid() && fv.CanSet() {

			// Controllo tipo stringa
			if fv.Kind() == reflect.String {
				fv.SetString(value)
			}

			if fv.Kind() == reflect.Int {
				content, err := strconv.ParseInt(value, 10, 64)
				if err == nil {
					fv.SetInt(content)
				}
			}
		}
	}
}
