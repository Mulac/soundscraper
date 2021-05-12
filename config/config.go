package config

import "sync"

var configSingleton Config
var once sync.Once

// Manager gives us our ConfigManager which just uses environment variables on our local machine... for now
func Manager() Config {
	once.Do(func() {
		configSingleton = &envConfig{}
	})

	return configSingleton
}
