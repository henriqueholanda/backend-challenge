package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/henriqueholanda/backend-challenge/backend/domain"
	"github.com/henriqueholanda/backend-challenge/backend/domain/amount"
	"github.com/henriqueholanda/backend-challenge/backend/handlers/response"
	"github.com/henriqueholanda/backend-challenge/backend/infrastructure/repository"
	"github.com/henriqueholanda/backend-challenge/backend/infrastructure/storage"
	"strconv"
)

type CheckoutHandlers struct {
	storage          storage.Storage
	repository       *repository.Memory
	amountCalculator amount.Calculator
}

type RequestParams struct {
	ProductCode string `json:"product-code"`
	Quantity string `json:"quantity"`
}

func NewCheckoutHandlers(
	storage storage.Storage,
	repository *repository.Memory,
	calculator amount.Calculator,
) *CheckoutHandlers {
	return &CheckoutHandlers{
		storage:          storage,
		repository:       repository,
		amountCalculator: calculator,
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
	ch.storage.Delete(context.Param("id"))

	response.NoContentResponse(context)
}

func (ch *CheckoutHandlers) FetchAmount(context *gin.Context) {

	basketStored, err := ch.storage.Fetch(context.Param("id"))
	if err != nil {
		response.NotFoundResponse(context, err.Error())
		return
	}

	basket := basketStored.(*domain.Basket)

	response.OkResponse(context, gin.H{
		"amount": ch.amountCalculator.Calculate(basket),
	})
}

func (ch *CheckoutHandlers) AddProduct(context *gin.Context) {
	var requestParams RequestParams

	context.BindJSON(&requestParams)

	productCode := requestParams.ProductCode
	quantity, _ := strconv.Atoi(requestParams.Quantity)

	if quantity == 0 {
		quantity = 1
	}

	basketStored, err := ch.storage.Fetch(context.Param("id"))
	if err != nil {
		response.NotFoundResponse(context, err.Error())
		return
	}

	product, err := ch.repository.GetByCode(productCode)
	if err != nil {
		response.BadRequestResponse(context, err.Error())
		return
	}

	basket := basketStored.(*domain.Basket)

	for i := 0; i < quantity; i++ {
		basket.AddProduct(product)
	}

	ch.storage.Save(basket.ID.String(), basket)

	response.CreatedResponse(context, gin.H{"basket": basket})
}
