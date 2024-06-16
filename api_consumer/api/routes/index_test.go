package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Humphryyy/docker-testing/api_consumer/redis"
)

func TestIndexRoute(t *testing.T) {
	redis.InitTest()

	t.Run("Test IndexRoute", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(IndexRoute)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := "Hi you've been here 1 times."
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
