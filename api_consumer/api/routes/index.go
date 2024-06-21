package routes

import (
	"fmt"
	"net/http"

	"github.com/Humphryyy/docker-testing/api_consumer/redis"
)

/* Route to index */
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	_, err := redis.Incr(r.Context(), "count")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reqCountCmd, err := redis.Get(r.Context(), "count")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reqCount, err := reqCountCmd.Int64()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hi you've been here %v times.", reqCount)
}
