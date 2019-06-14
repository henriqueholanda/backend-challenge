package repository

import (
	"errors"

	"github.com/henriqueholanda/backend-challenge/backend/domain"
)

type Repository interface {
	GetByCode(code string) (domain.Product, error)
}

type Memory struct {
	products map[string]domain.Product
}

func NewProductRepository() *Memory {
	products := map[string]domain.Product{
		"VOUCHER": {
			"VOUCHER", "Cabify Voucher", 5.00,
		},
		"TSHIRT": {
			"TSHIRT", "Cabify T-Shirt", 20.00,
		},
		"MUG": {
			"MUG", "Cabify Coffee Mug", 7.50,
		},
	}

	return &Memory{products}
}

func (m *Memory) GetByCode(code string) (domain.Product, error) {
	product, ok := m.products[code]

	if !ok {
		return domain.Product{}, errors.New("product not found")
	}

	return product, nil
}

func (m *Memory) GetAll() domain.Products {
	var products domain.Products
	for _, product := range m.products {
		products = append(products, product)
	}
	return products
}
