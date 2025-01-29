package env

import "os"

const (
	Port          = "BACKEND_PORT"
	DefaultPort   = "8080"
	Host          = "BACKEND_HOST"
	DefaultHost   = "localhost"
	DefaultDB     = "lore-keeper-db"
	DefaultMode   = "dev"
	DefaultDBType = "sqlite"
)

func GetEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
