package cacheconf

import (
	"os"
	"reflect"
	"strconv"
)

// CacheConfigInterface - Interfaccia per implementare CacheConfig
type CacheConfigInterface interface {
	GetFieldMapper() map[string]string
}

// DefaultCacheConfig - Definisce la configurazione generica dell'aplicazione
type DefaultCacheConfig struct {
	AppName          string
	StringConnection string
	AppPort          string
}

// config - Stringa con tutte le configurazione caricate
var config string

// setField - si occupa di impostare  attrun campo averso la reflection, c รจ necessario sia un puntatore a una struttura
func setField(c CacheConfigInterface, name string, value string) {

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

// GetCurrentConfig - Restituisce l'attuale configurazione
func GetCurrentConfig() string {
	return config
}

// LoadEnvConfig - si occupa di caricare tutte le configurazioni dell'env nella struttura di configurazione
func LoadEnvConfig(c CacheConfigInterface) {

	config = "\n"

	d := DefaultCacheConfig{}
	for envName, StructName := range d.GetFieldMapper() {
		setField(c, StructName, os.Getenv(envName))
		config = config + StructName + " -> " + os.Getenv(envName) + "\n"
	}

	for envName, StructName := range c.GetFieldMapper() {
		setField(c, StructName, os.Getenv(envName))
		config = config + StructName + " -> " + os.Getenv(envName) + "\n"
	}
}

// GetFieldMapper - Restitusice la mappatura tra l'env e i campi di configurazione
func (d DefaultCacheConfig) GetFieldMapper() map[string]string {
	return map[string]string{
		"APP_NAME":          "AppName",
		"STRING_CONNECTION": "StringConnection",
		"APP_PORT":          "AppPort",
	}
}
