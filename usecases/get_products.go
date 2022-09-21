package usecases

import (
	"discountapp/usecases/interfaces"
	"discountapp/usecases/outputs"
)

type GetProducts struct {
	Repository interfaces.IRepository
}

func NewGetProducts(repository interfaces.IRepository) interfaces.IGetProducts {
	return &GetProducts{Repository: repository}
}

func (usecase *GetProducts) Perform() (*outputs.ProductsOutput, error) {
	products, err := usecase.Repository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return &outputs.ProductsOutput{
		Products: products,
	}, nil
}
