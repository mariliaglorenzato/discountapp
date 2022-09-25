package domain_test

import (
	"discountapp/domain"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Discount Domain", func() {
	BeforeEach(func() {
	})
	JustBeforeEach(func() {
	})

	When("#GetTotalDiscount", func() {
		totalDiscount := domain.GetTotalDiscount(30, nil)
		It("returns expected discount structure", func() {
			Expect(totalDiscount).To(Equal(float64(0.03)))
		})
	})
	When("#GetPriceWithDiscount", func() {
		When("Today is not special a day and no event discount is loaded ", func() {
			// age 38 = 0,038 = 3,8%
			// totalPrice: 3000 - 3000 * 0,038 =
			totalPrice, totalDiscount := domain.GetPriceWithDiscount(3000, 38, nil)
			It("returns expected discount structure", func() {
				Expect(totalDiscount).To(Equal(float64(0.038)))
				Expect(totalPrice).To(Equal(uint64(2886)))
			})
		})
		When("Today is a special day and events discount are added to total discount", func() {
		})
	})
})
