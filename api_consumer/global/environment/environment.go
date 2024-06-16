package environment

import (
	"fmt"
	"os"
)

var (
	RedisUrl     string
	RabbitMQUrl  string
	ConsumerPort string
)

func Load() error {
	variables := []string{
		"REDIS_URL",
		"RABBITMQ_URL",
		"CONSUMER_PORT",
	}

	for _, variable := range variables {
		value := os.Getenv(variable)
		if value == "" {
			return fmt.Errorf("missing %s", variable)
		}

		switch variable {
		case "REDIS_URL":
			RedisUrl = value
		case "RABBITMQ_URL":
			RabbitMQUrl = value
		case "CONSUMER_PORT":
			ConsumerPort = value
		}
	}

	return nil
}
