package interfaces

import (
	"discountapp/usecases/inputs"
	"discountapp/usecases/outputs"
)

type IGetDiscount interface {
	Perform(discountInput *inputs.DiscountInput) (*outputs.DiscountOutput, error)
}
