package usecases

import (
	domain "discountapp/domain"
	"discountapp/usecases/inputs"
	"discountapp/usecases/interfaces"
)

type GetProduct struct {
	Repository interfaces.IRepository
}

func NewGetProduct(repository interfaces.IRepository) interfaces.IGetProduct {
	return &GetProduct{Repository: repository}
}

func (usecase *GetProduct) Perform(ProductBySlugInput *inputs.ProductBySlugInput) (*domain.Product, error) {
	product := domain.Product{Slug: ProductBySlugInput.Slug}
	output, err := usecase.Repository.GetProduct(&product)
	if err != nil {
		return nil, err
	}

	return output, nil
}
