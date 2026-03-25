package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConfig     DBConfig
}

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBMigrate  bool
}

var config Config

func Init() {
	_ = godotenv.Load()
	config.DbConfig.DBUser = os.Getenv("DB_USER")
	config.DbConfig.DBPassword = os.Getenv("DB_PASSWORD")
	config.DbConfig.DBHost = os.Getenv("DB_HOST")
	config.DbConfig.DBPort = os.Getenv("DB_PORT")
	config.DbConfig.DBName = os.Getenv("DB_NAME")
	config.DbConfig.DBMigrate = os.Getenv("DB_MIGRATE") == "true"

}

func GetConfig() Config {
	return config
}
