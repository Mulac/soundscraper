package config

// Config is how we can tweak and change our application
//
// Using an interface like this prepares us for the future when we might want to get our config from somewhere else
type Config interface {
	GetString(key string) string
}
