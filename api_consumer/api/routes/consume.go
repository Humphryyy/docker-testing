package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Humphryyy/docker-testing/api_consumer/rabbitmq"
)

/* Route to consume and process all events from a workspace */
func ConsumeRoute(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rabbitmq.Publish("messages", "test", body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Message sent!")
}
