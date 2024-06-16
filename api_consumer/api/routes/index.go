package routes

import (
	"fmt"
	"net/http"

	"github.com/Humphryyy/docker-testing/api_consumer/redis"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	redisClient := redis.Client()

	redisClient.Incr(r.Context(), "count")
	reqCount, _ := redisClient.Get(r.Context(), "count").Int64()

	fmt.Fprintf(w, "Hi you've been here %v times.", reqCount)
}
