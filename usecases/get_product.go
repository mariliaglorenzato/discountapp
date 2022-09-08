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

func (usecase *GetProduct) Perform(ProductByTitleInput *inputs.ProductByTitleInput) (*domain.Product, error) {
	product := domain.Product{Title: ProductByTitleInput.Title}
	output, err := usecase.Repository.GetProduct(&product)
	if err != nil {
		return nil, err
	}

	return output, nil
}
