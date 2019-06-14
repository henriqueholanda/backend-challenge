package domain

import "github.com/google/uuid"

type Basket struct {
	ID       uuid.UUID `json:"id"`
	Products Products  `json:"products"`
}

func NewBasket() *Basket {
	return &Basket{
		ID: uuid.New(),
	}
}

func (basket *Basket) AddProduct(product Product) {
	basket.Products = append(basket.Products, product)
}
