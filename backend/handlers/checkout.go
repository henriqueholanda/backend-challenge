package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/henriqueholanda/backend-challenge/backend/domain"
	"github.com/henriqueholanda/backend-challenge/backend/handlers/response"
	"github.com/henriqueholanda/backend-challenge/backend/infrastructure/storage"
	"net/http"
)

type CheckoutHandlers struct {
	storage          storage.Storage
}


func NewCheckoutHandlers(storage storage.Storage) *CheckoutHandlers {
	return &CheckoutHandlers{
		storage:          storage,
	}
}

func (ch *CheckoutHandlers) Create(context *gin.Context) {
	basket := domain.NewBasket()

	ch.storage.Save(basket.ID.String(), basket)

	response.CreatedResponse(context, gin.H{
		"id": basket.ID,
	})
}

func (ch *CheckoutHandlers) Delete(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "Method not implemented"})
}

func (ch *CheckoutHandlers) FetchAmount(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "Method not implemented"})
}

func (ch *CheckoutHandlers) AddProduct(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "Method not implemented"})
}
