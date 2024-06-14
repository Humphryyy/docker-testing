package main

/*

docker build . -t go-container
docker compose up

*/
import (
	"fmt"
	"io"
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
		redisClient.Incr(r.Context(), "count")
		reqCount, _ := redisClient.Get(r.Context(), "count").Int64()

		fmt.Fprintf(w, "Hi you've been here %v times.", reqCount)
	})

	amqpChan, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	amqpChan.ExchangeDeclare("messages", "fanout", true, true, false, false, nil)

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for i := 0; i < 1; i++ {
			err = amqpChan.Publish("messages", "", false, false, amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(string(body) + " : " + fmt.Sprint(i)),
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		fmt.Fprint(w, "Message sent!")
	})

	http.ListenAndServe(":8080", nil)
}
