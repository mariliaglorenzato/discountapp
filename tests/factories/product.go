package factories

import (
	"time"

	"discountapp/domain"
)

func BuildProduct() *domain.Product {
	return &domain.Product{
		ID:          1,
		Title:       "Caneta de CD",
		Slug:        "caneta-de-cd",
		Description: "Caneta preta não sai com alcool",
		Price:       3000,
		TotalPrice:  0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
}

func BuildStubbedProduct() *domain.Product {
	return &domain.Product{
		Title:       "Caneta de CD",
		Slug:        "caneta-de-cd",
		Description: "Caneta preta não sai com alcool",
		Price:       3000,
	}
}
