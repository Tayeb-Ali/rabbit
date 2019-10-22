package main

import (
	"os"
)

type Config struct {
	serviceName string
	dotLogs     string
}

type RabbitConfig struct {
	uri          string
	storageQueue string
	gatewayQueue string
}

var config = Config{
	serviceName: getEnv("SERVICE_NAME", "storage"),
	dotLogs:     getEnv("DOT_LOGS", ".logs"),
}

var rabbitConfig = RabbitConfig{
	uri:          getEnv("RABBIT_URI", "amqp://admin:admin@localhost:5672/"),
	gatewayQueue: getEnv("RABBIT_GATEWAY_QUEUE", "gateway"),
	storageQueue: getEnv("RABBIT_STORAGE_QUEUE", "storage"),
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
