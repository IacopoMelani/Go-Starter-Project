package cacheconf

import (
	"os"
	"reflect"
	"strconv"

	refl "github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/reflect"
)

// CacheConfigInterface - Interface to implements CacheConfig
// No methods, all you need is to define the tag "config" for your Custom CacheConfig
type CacheConfigInterface interface {
	GetDefaultCacheConfig() CacheConfigInterface
}

// DefaultCacheConfig - Defines the standard configuration
type DefaultCacheConfig struct {
	AppName          string `config:"APP_NAME"`
	Debug            bool   `config:"DEBUG"`
	SQLDriver        string `config:"SQL_DRIVER"`
	StringConnection string `config:"STRING_CONNECTION"`
	AppPort          string `config:"APP_PORT"`
}

// ConfigTagName - Defines the tag name to permit CacheConfig stores your configurations
const ConfigTagName = "config"

// cacheConf - Defines the global instance of CacheConfigInterface
var cacheConf *DefaultCacheConfig

// config - Stores all configurations for display
var config string

// loadEnvByFieldsMapper - Loops over env fields name and sets the value to CacheConfig instance passed, with the env value read with os pkg
func loadEnvByFieldsMapper(c CacheConfigInterface, envFields []string, structFieldsName []string) {
	for i := 0; i < len(envFields); i++ {
		setField(c, structFieldsName[i], os.Getenv(envFields[i]))
		config = config + structFieldsName[i] + " -> " + os.Getenv(envFields[i]) + "\n"

	}
}

// setField - Try to sets the value passed to an CacheConfigInterface, "name" is the struct field name
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

// GetCurrentConfig - Returns the current config
func GetCurrentConfig() string {
	return config
}

// Debug - Returns if application is in debug
func Debug() bool {
	return cacheConf.Debug
}

// LoadEnvConfig - Try to init the CacheConfigInterface using DefaultCacheConfig for default config fields and the custom CacheConfig passed as CacheConfigInterface, requires that "c" is a pointer to your custom CacheConfig
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
