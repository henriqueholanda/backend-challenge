package main

import (
	"github.com/henriqueholanda/backend-challenge/backend/application"
	"github.com/henriqueholanda/backend-challenge/backend/handlers"
)

func main() {
	checkoutHandlers := handlers.NewCheckoutHandlers()

	router := application.SetupRouter(checkoutHandlers)

	router.Run(":80")
}
