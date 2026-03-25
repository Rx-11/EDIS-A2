package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CustomerSvcURL string
	BookSvcURL     string
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
