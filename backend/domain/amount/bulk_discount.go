package amount

import (
	"github.com/henriqueholanda/backend-challenge/backend/domain"
)

type BulkDiscount struct {
	code     string
	quantity int
	value    float64
}

func NewBulkDiscount(code string, quantity int, value float64) *BulkDiscount {
	return &BulkDiscount{
		code:     code,
		quantity: quantity,
		value:    value,
	}
}

func (bd *BulkDiscount) Calculate(basket *domain.Basket, amount float64) float64 {
	var total int
	var price float64

	for _, product := range basket.Products {
		if product.Code == bd.code {
			total = total + 1
			price = product.Price
		}
	}

	if total >= bd.quantity {
		amount = amount - (price * float64(total))
		amount = amount + (bd.value * float64(total))
	}

	return amount
}
