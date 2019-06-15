package main

import (
	"github.com/henriqueholanda/backend-challenge/backend/application"
	"github.com/henriqueholanda/backend-challenge/backend/handlers"
	"github.com/henriqueholanda/backend-challenge/backend/infrastructure/storage"
)

func main() {
	memoryStorage := storage.NewMemoryStorage()

	checkoutHandlers := handlers.NewCheckoutHandlers(memoryStorage)

	router := application.SetupRouter(checkoutHandlers)

	router.Run(":80")
}
