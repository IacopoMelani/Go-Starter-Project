package config

import (
	"sync"

	cacheconf "github.com/IacopoMelani/Go-Starter-Project/pkg/cache_config"
)

// CacheConfig - struttura dove immagazzinare le configurazioni
type CacheConfig struct {
	cacheconf.DefaultCacheConfig
	UserTimeToRefresh int
}

var (
	cacheConfig *CacheConfig
	once        sync.Once
)

// GetInstance - restituisce l'unica istanza della struttura contenente le configurazioni
func GetInstance() *CacheConfig {
	once.Do(func() {
		cacheConfig = &CacheConfig{}
		cacheconf.LoadEnvConfig(cacheConfig)
	})
	return cacheConfig
}

// GetFieldMapper - Si occupa di restituire l'array di mappatura dell'env
func (c CacheConfig) GetFieldMapper() map[string]string {
	return map[string]string{
		"USER_TIME_TO_REFRESH": "UserTimeToRefresh",
	}
}
