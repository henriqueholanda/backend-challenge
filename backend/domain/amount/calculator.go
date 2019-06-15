package amount

import (
	"github.com/henriqueholanda/backend-challenge/backend/domain"
)

type CalculatorStrategy interface {
	Calculate(basket *domain.Basket, amount float64) float64
}

type Calculator struct {
	strategies []CalculatorStrategy
}

func NewAmountCalculator(strategies ...CalculatorStrategy) Calculator {
	return Calculator{strategies}
}

func (c *Calculator) Calculate(basket *domain.Basket) float64 {
	var amount float64

	for _, strategy := range c.strategies {
		amount = strategy.Calculate(basket, amount)
	}

	return amount
}
