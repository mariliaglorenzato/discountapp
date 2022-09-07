package interfaces

import (
	"discountapp/domain"
	"discountapp/usecases/inputs"
)

type IGetProduct interface {
	Perform(productInput *inputs.ProductByTitleInput) (*domain.Product, error)
}
