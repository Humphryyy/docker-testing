package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Humphryyy/docker-testing/api_consumer/rabbitmq"
)

func TestConsumeRoute(t *testing.T) {
	rabbitmq.InitTest()

	t.Run("Test ConsumeRoute", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/consume", bytes.NewBuffer([]byte("test")))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ConsumeRoute)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := "Message sent!"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
