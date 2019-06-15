package handlers_test

import (
	"encoding/json"
	"github.com/henriqueholanda/backend-challenge/backend/application"
	"github.com/henriqueholanda/backend-challenge/backend/handlers"
	"github.com/henriqueholanda/backend-challenge/backend/infrastructure/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasketsCreateHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/v1/checkout/basket", nil)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	memoryStorage := storage.NewMemoryStorage()

	checkoutHandlers := handlers.NewCheckoutHandlers(memoryStorage)

	application.SetupRouter(checkoutHandlers).ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusCreated,
		)
	}

	res := make(map[string]string)

	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fail()
	}

	if _, ok := res["id"]; !ok {
		t.Error("response id key not found")
	}
}