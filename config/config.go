package config

type AppConfig struct {
	ServerPort int
	DBPort     int
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
}

func NewConfig() *AppConfig {
	return
}
