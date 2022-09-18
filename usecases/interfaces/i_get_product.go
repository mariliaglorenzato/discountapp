package interfaces

import (
	"discountapp/domain"
	"discountapp/usecases/inputs"
)

type IGetProduct interface {
	Perform(productInput *inputs.ProductBySlugInput) (*domain.Product, error)
}
