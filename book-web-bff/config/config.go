package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConfig       DBConfig
	CustomerSvcURL string
	BookSvcURL     string
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
	config.CustomerSvcURL = os.Getenv("CUSTOMER_SVC_URL")
	config.BookSvcURL = os.Getenv("BOOK_SVC_URL")

	initFiberClient()
}

func GetConfig() Config {
	return config
}
