package api

import (
	"log"
	"net/http"

	"github.com/Humphryyy/docker-testing/api_consumer/api/routes"
	"github.com/Humphryyy/docker-testing/api_consumer/global/environment"
)

func Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	for _, route := range routes.GetRoutes() {
		mux.HandleFunc(route.Path, route.Handler)
	}

	log.Println("API Consumer running on port", environment.ConsumerPort)
	err := http.ListenAndServe(":"+environment.ConsumerPort, mux)
	if err != nil {
		return err
	}

	return nil
}
