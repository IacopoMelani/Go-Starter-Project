package cacheconf

import (
	"os"
	"reflect"
	"strconv"

	refl "github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/reflect"
)

// CacheConfigInterface - Interfaccia per implementare CacheConfig
type CacheConfigInterface interface {
	GetDefaultCacheConfig() CacheConfigInterface
}

// DefaultCacheConfig - Definisce la configurazione generica dell'aplicazione
type DefaultCacheConfig struct {
	AppName          string `config:"APP_NAME"`
	Debug            bool   `config:"DEBUG"`
	SQLDriver        string `config:"SQL_DRIVER"`
	StringConnection string `config:"STRING_CONNECTION"`
	AppPort          string `config:"APP_PORT"`
}

// ConfigTagName - Definisce il nome del tag config per la mappatura della configurazione
const ConfigTagName = "config"

// cacheConf - Defines the global instance of CacheConfigInterface
var cacheConf *DefaultCacheConfig

// config - Stringa con tutte le configurazione caricate
var config string

func loadEnvByFieldsMapper(c CacheConfigInterface, envFields []string, structFieldsName []string) {
	for i := 0; i < len(envFields); i++ {
		setField(c, structFieldsName[i], os.Getenv(envFields[i]))
		config = config + structFieldsName[i] + " -> " + os.Getenv(envFields[i]) + "\n"

	}
}

// setField - si occupa di impostare  attrun campo averso la reflection, c รจ necessario sia un puntatore a una struttura
func setField(c CacheConfigInterface, name string, value string) {

	rv := reflect.ValueOf(c)

	if rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.Struct {

		rv = rv.Elem()

		fv := rv.FieldByName(name)
		if fv.IsValid() && fv.CanSet() {

			if fv.Kind() == reflect.String {
				fv.SetString(value)
			}

			if fv.Kind() == reflect.Bool {
				content, err := strconv.ParseBool(value)
				if err == nil {
					fv.SetBool(content)
				}
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

// Debug - Returns if application is in debug
func Debug() bool {
	return cacheConf.Debug
}

// LoadEnvConfig - si occupa di caricare tutte le configurazioni dell'env nella struttura di configurazione
func LoadEnvConfig(c CacheConfigInterface) {

	config = "\n"

	d := DefaultCacheConfig{}

	envFields, structFieldsName := refl.GetStructFieldsNameAndTagByTagName(d, ConfigTagName)
	loadEnvByFieldsMapper(c, envFields, structFieldsName)

	envFields, structFieldsName = refl.GetStructFieldsNameAndTagByTagName(c, ConfigTagName)
	loadEnvByFieldsMapper(c, envFields, structFieldsName)

	cacheConf = c.GetDefaultCacheConfig().(*DefaultCacheConfig)
}

// GetDefaultCacheConfig - Return the instance of CacheConfigInterface
func (d *DefaultCacheConfig) GetDefaultCacheConfig() CacheConfigInterface {
	return d
}
