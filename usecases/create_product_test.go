package usecases

import (
	"time"

	"discountapp/domain"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creating a new Product", func() {
	var getEntityProduct func() *domain.Product

	BeforeEach(func() {
		getEntityProduct = func() *domain.Product {
			return &domain.Product{
				Title:       "Caneta de CD",
				Description: "Caneta preta não sai com alcool",
				Price:       30,
				CreatedAt:   time.Now(),
			}
		}
	})

	It("creates a new product", func() {
		Expect(getEntityProduct)
	})
})
