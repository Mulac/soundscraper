package config

import "os"

type envConfig struct{}

func (c *envConfig) GetString(key string) string {
	return os.Getenv(key)
}
