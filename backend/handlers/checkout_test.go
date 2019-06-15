package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/henriqueholanda/backend-challenge/backend/application"
	"github.com/henriqueholanda/backend-challenge/backend/domain"
	"github.com/henriqueholanda/backend-challenge/backend/domain/amount"
	"github.com/henriqueholanda/backend-challenge/backend/handlers"
	"github.com/henriqueholanda/backend-challenge/backend/infrastructure/repository"
	"github.com/henriqueholanda/backend-challenge/backend/infrastructure/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MemoryStorageMock struct {
	basket interface{}
}

func (s *MemoryStorageMock) Delete(key string) {}

func (s *MemoryStorageMock) Fetch(key string) (interface{}, error) {
	return s.basket, nil
}

func (s *MemoryStorageMock) Save(key string, value interface{}) {}

func TestBasketsCreateHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/v1/checkout/basket", nil)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	memoryStorage := storage.NewMemoryStorage()
	productRepository := repository.NewProductRepository()
	amountCalculator := amount.NewAmountCalculator(
		amount.NewSum(),
		amount.NewBuyTwoPayOnePromotion("VOUCHER"),
		amount.NewBulkDiscount("TSHIRT", 3, 19.00),
	)

	checkoutHandlers := handlers.NewCheckoutHandlers(memoryStorage, productRepository, amountCalculator)

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

func TestBasketsDeleteHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/v1/checkout/basket/1", nil)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	memoryStorage := storage.NewMemoryStorage()
	productRepository := repository.NewProductRepository()
	amountCalculator := amount.NewAmountCalculator(
		amount.NewSum(),
		amount.NewBuyTwoPayOnePromotion("VOUCHER"),
		amount.NewBulkDiscount("TSHIRT", 3, 19.00),
	)

	checkoutHandlers := handlers.NewCheckoutHandlers(memoryStorage, productRepository, amountCalculator)

	application.SetupRouter(checkoutHandlers).ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusNoContent {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusNoContent,
		)
	}
}

func TestBasketsAddItemHandlerWhenBasketNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/v1/checkout/basket/123/products", nil)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	memoryStorage := storage.NewMemoryStorage()
	productRepository := repository.NewProductRepository()
	amountCalculator := amount.NewAmountCalculator(
		amount.NewSum(),
		amount.NewBuyTwoPayOnePromotion("VOUCHER"),
		amount.NewBulkDiscount("TSHIRT", 3, 19.00),
	)

	checkoutHandlers := handlers.NewCheckoutHandlers(memoryStorage, productRepository, amountCalculator)

	application.SetupRouter(checkoutHandlers).ServeHTTP(rec, req)

	if message := rec.Body.String(); message != "{\"error\":\"key not found\"}" {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			message,
			"{\"error\":\"key not found\"}",
		)
	}
}

func TestBasketsAddItemHanderWithInvalidRequest(t *testing.T) {
	buf := bytes.NewBufferString(`{"product-code":"}`)

	req, err := http.NewRequest(http.MethodPost, "/v1/checkout/basket/1/products", buf)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	memoryStorage := new(MemoryStorageMock)
	memoryStorage.basket = domain.Basket{}

	basketsHandler := handlers.NewCheckoutHandlers(
		memoryStorage,
		repository.NewProductRepository(),
		amount.NewAmountCalculator(),
	)

	handler := application.SetupRouter(basketsHandler)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusBadRequest,
		)
	}
}

func TestBasketsAddItemHandlerWithInvalidProduct(t *testing.T) {
	buf := bytes.NewBufferString(`{"product-code":"SKATE"}`)

	req, err := http.NewRequest(http.MethodPost, "/v1/checkout/basket/1/products", buf)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	memoryStorage := new(MemoryStorageMock)
	memoryStorage.basket = domain.Basket{}

	basketsHandler := handlers.NewCheckoutHandlers(
		memoryStorage,
		repository.NewProductRepository(),
		amount.NewAmountCalculator(),
	)

	handler := application.SetupRouter(basketsHandler)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusBadRequest,
		)
	}
}

func TestBasketsAddItemHandler(t *testing.T) {
	buf := bytes.NewBufferString(`{"product-code":"MUG"}`)

	req, err := http.NewRequest(
		http.MethodPost,
		"/v1/checkout/basket/1/products",
		buf,
	)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	memoryStorage := new(MemoryStorageMock)
	memoryStorage.basket = &domain.Basket{}

	basketsHandler := handlers.NewCheckoutHandlers(
		memoryStorage,
		repository.NewProductRepository(),
		amount.NewAmountCalculator(),
	)

	handler := application.SetupRouter(basketsHandler)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusCreated,
		)
	}
}

func TestBasketsAmountHandlerWhenBasketNotFound(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodGet,
		"/v1/checkout/basket/123/amount",
		nil,
	)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	basketsHandler := handlers.NewCheckoutHandlers(
		storage.NewMemoryStorage(),
		repository.NewProductRepository(),
		amount.NewAmountCalculator(),
	)

	handler := application.SetupRouter(basketsHandler)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusNotFound {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusNotFound,
		)
	}
}

func TestBasketsAmountHandler(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodGet,
		"/v1/checkout/basket/1/amount",
		nil,
	)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	product := domain.Product{"MUG", "Cabify Coffee Mug", 7.50}
	bkt := domain.NewBasket()
	bkt.AddProduct(product)

	store := &MemoryStorageMock{}
	store.basket = bkt
	basketsHandler := handlers.NewCheckoutHandlers(
		store,
		repository.NewProductRepository(),
		amount.NewAmountCalculator(amount.NewSum()),
	)

	handler := application.SetupRouter(basketsHandler)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	res := make(map[string]float64)

	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fail()
	}

	expectedAmount := 7.50

	if res["amount"] != expectedAmount {
		t.Errorf(
			"wrong amount value: got %v want %v",
			res["amount"],
			expectedAmount,
		)
	}
}
