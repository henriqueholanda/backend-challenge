package amount

import (
	"github.com/henriqueholanda/backend-challenge/backend/domain"
)

type BuyTwoPayOne struct {
	code string
}

func NewBuyTwoPayOnePromotion(code string) *BuyTwoPayOne {
	return &BuyTwoPayOne{
		code: code,
	}
}

func (fp *BuyTwoPayOne) Calculate(basket *domain.Basket, amount float64) float64 {
	var total int
	var price float64

	for _, product := range basket.Products {
		if product.Code == fp.code {
			total = total + 1
			price = product.Price
		}
	}

	if total >= 2 {
		amount = amount - price
	}

	return amount
}
