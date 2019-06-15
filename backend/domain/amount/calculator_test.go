package amount

import (
	"github.com/henriqueholanda/backend-challenge/backend/domain"
	"testing"
)

func TestAmountCalculator(t *testing.T) {
	calculator := NewAmountCalculator(
		NewSum(),
		NewBuyTwoPayOnePromotion("VOUCHER"),
		NewBuyTwoPayOnePromotion("CHOCOLATE"),
		NewBulkDiscount("TSHIRT", 3, 19.0),
		NewBulkDiscount("PANTS", 5, 10.0),
	)

	basket := domain.NewBasket()
	basket.AddProduct(domain.Product{"VOUCHER", "Cabify Voucher", 5.00})
	basket.AddProduct(domain.Product{"VOUCHER", "Cabify Voucher", 5.00})
	basket.AddProduct(domain.Product{"VOUCHER", "Cabify Voucher", 5.00})
	basket.AddProduct(domain.Product{"TSHIRT", "Cabify T-Shirt", 20.00})
	basket.AddProduct(domain.Product{"TSHIRT", "Cabify T-Shirt", 20.00})
	basket.AddProduct(domain.Product{"TSHIRT", "Cabify T-Shirt", 20.00})
	basket.AddProduct(domain.Product{"MUG", "Cabify Coffee Mug", 7.50})

	if result := calculator.Calculate(basket); result != 74.5 {
		t.Errorf("Amount calculation is wrong: got %v want %v", result, 74.5)
	}
}
