package amount

import (
	"github.com/henriqueholanda/backend-challenge/backend/domain"
)

type Sum struct {
}

func NewSum() *Sum {
	return &Sum{}
}

func (s *Sum) Calculate(basket *domain.Basket, amount float64) float64 {
	for _, product := range basket.Products {
		amount = amount + product.Price
	}

	return amount
}
