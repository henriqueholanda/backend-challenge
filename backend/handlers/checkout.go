package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckoutHandlers struct {

}


func NewCheckoutHandlers() *CheckoutHandlers {
	return &CheckoutHandlers{}
}

func (ch *CheckoutHandlers) Create(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "Method not implemented"})
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
