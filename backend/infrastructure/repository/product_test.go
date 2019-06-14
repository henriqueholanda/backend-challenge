package repository

import (
	"testing"
)

func TestProductRepoGetByCode(t *testing.T) {
	repo := NewProductRepository()

	if _, err := repo.GetByCode("SHOE"); err == nil {
		t.Errorf(
			"method should returned an error: got %v want %v",
			err.Error(),
			"product not found",
		)
	}

	product, err := repo.GetByCode("MUG")

	if err != nil {
		t.Fatal()
	}

	if product.Code != "MUG" {
		t.Error("invalid product code")
	}
}
