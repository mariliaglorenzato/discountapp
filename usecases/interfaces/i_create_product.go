package interfaces

import (
	"discountapp/domain"
	"discountapp/usecases/inputs"
)

type ICreateProduct interface {
	Perform(productInput *inputs.ProductInput) (*domain.Product, error)
}
