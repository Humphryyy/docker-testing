package main

/*

docker build . -t go-container
docker compose up

*/
import (
	"context"
	"fmt"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
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

	fmt.Println("Hello, World!")

	opts, err := redis.ParseURL("redis://redis:6379")
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(opts)

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redisClient.Incr(context.Background(), "count")
		reqCount, _ := redisClient.Get(context.Background(), "count").Int64()

		fmt.Fprintf(w, "Hi you've been here %v times.", reqCount)
	})

	http.ListenAndServe(":8080", nil)
}
