package rabbitmq

import (
	"fmt"
	"sync"
	"time"

	"github.com/Humphryyy/docker-testing/api_consumer/global/environment"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQService interface {
	ExchangeDeclare(name, kind string) error
	Publish(exchange, key string, message []byte) error
	Close()
}

var rabbitMQ RabbitMQService

type rabbitMQService struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func Init() error {
	retries := 0

retry:
	conn, err := amqp.Dial(environment.RabbitMQUrl)
	if err != nil {
		if retries < 10 {
			retries++
			time.Sleep(5 * time.Second)
			goto retry
		}
		return err
	}

	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	rq := &rabbitMQService{
		conn:    conn,
		channel: channel,
	}

	rabbitMQ = rq

	err = DeclareExchanges()
	if err != nil {
		return err
	}

	return nil
}

func (r *rabbitMQService) ExchangeDeclare(name, kind string) error {
	err := r.channel.ExchangeDeclare(name, kind, true, true, false, false, nil)
	if err != nil {
		return err
	}

	return nil
}

func (r *rabbitMQService) Publish(exchange, key string, message []byte) error {
	err := r.channel.Publish(exchange, key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        message,
	})
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

func (r *rabbitMQService) Close() {
	r.channel.Close()
	r.conn.Close()
}

func Close() {
	if rabbitMQ == nil {
		return
	}

	rabbitMQ.Close()
}

func Publish(exchange, key string, message []byte) error {
	if rabbitMQ == nil {
		return fmt.Errorf("rabbitmq not initialized")
	}

	err := rabbitMQ.Publish(exchange, key, message)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

func DeclareExchanges() error {
	if rabbitMQ == nil {
		return fmt.Errorf("rabbitmq not initialized")
	}

	exchanges := [][]string{
		{"messages", "direct"},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(exchanges))

	var errors []error
	for _, exchange := range exchanges {
		go func(exchange []string) {
			defer wg.Done()
			err := rabbitMQ.ExchangeDeclare(exchange[0], exchange[1])
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to declare exchange %s: %w", exchange[0], err))
			}
		}(exchange)
	}

	wg.Wait()

	if len(errors) > 0 {
		return fmt.Errorf("failed to declare exchanges: %v", errors)
	}

	return nil
}
