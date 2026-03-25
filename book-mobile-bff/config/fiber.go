package config

import (
	"time"

	"github.com/gofiber/fiber/v3/client"
)

var fiberClient *client.Client

func initFiberClient() {
	fiberClient = client.New()
	fiberClient.SetTimeout(10 * time.Second)
}

func GetFiberClient() *client.Client {
	if fiberClient == nil {
		initFiberClient()
	}
	return fiberClient
}
