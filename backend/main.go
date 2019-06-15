package main

import (
	"github.com/henriqueholanda/backend-challenge/backend/application"
	"github.com/henriqueholanda/backend-challenge/backend/domain/amount"
	"github.com/henriqueholanda/backend-challenge/backend/handlers"
	"github.com/henriqueholanda/backend-challenge/backend/infrastructure/repository"
	"github.com/henriqueholanda/backend-challenge/backend/infrastructure/storage"
)

func main() {
	memoryStorage := storage.NewMemoryStorage()
	productRepository := repository.NewProductRepository()
	amountCalculator := amount.NewAmountCalculator(
		amount.NewSum(),
		amount.NewBuyTwoPayOnePromotion("VOUCHER"),
		amount.NewBulkDiscount("TSHIRT", 3, 19.00),
	)

	checkoutHandlers := handlers.NewCheckoutHandlers(memoryStorage, productRepository, amountCalculator)

	router := application.SetupRouter(checkoutHandlers)

	router.Run(":80")
}
