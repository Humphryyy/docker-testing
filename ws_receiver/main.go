package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	retries := 0

retry:
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		if retries < 10 {
			retries++
			time.Sleep(5 * time.Second)
			goto retry
		}
		panic(err)
	}
	defer conn.Close()

	amqpChan, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	var forever chan struct{}

	for i := 0; i < 1; i++ {
		go CreateSub(amqpChan)
	}

	<-forever
}

func CreateSub(amqpChan *amqp.Channel) {
	id := uuid.New().String()

	q, err := amqpChan.QueueDeclare(id, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = amqpChan.QueueBind(id, id, "messages", false, nil)

	msgs, err := amqpChan.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	go func() {
		for d := range msgs {
			log.Printf("%v | [x] %s", id, d.Body)
		}
	}()

	log.Printf("Subscribed to %s", id)
}
