package main

import (
	"log"

	"github.com/Humphryyy/docker-testing/api_consumer/api"
	"github.com/Humphryyy/docker-testing/api_consumer/global/environment"
	"github.com/Humphryyy/docker-testing/api_consumer/rabbitmq"
	"github.com/Humphryyy/docker-testing/api_consumer/redis"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting API Consumer")

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = environment.Load()
	if err != nil {
		panic(err)
	}

	log.Println("Environment loaded")

	err = rabbitmq.Init()
	if err != nil {
		panic(err)
	}

	log.Println("RabbitMQ connected")

	err = redis.Init()
	if err != nil {
		panic(err)
	}

	log.Println("Redis connected")

	defer exit()

	err = api.Run()
	if err != nil {
		panic(err)
	}
}

func exit() {
	rabbitmq.Close()
}
