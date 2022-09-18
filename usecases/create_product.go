package usecases

import (
	domain "discountapp/domain"
	"discountapp/usecases/inputs"
	"discountapp/usecases/interfaces"
)

type CreateProduct struct {
	Repository interfaces.IRepository
}

func NewCreateProduct(repository interfaces.IRepository) interfaces.ICreateProduct {
	return &CreateProduct{Repository: repository}
}

func (usecase *CreateProduct) Perform(productInput *inputs.ProductInput) (*domain.Product, error) {
	product := domain.Product{
		Title:       productInput.Title,
		Description: productInput.Description,
		Price:       productInput.Price,
		Slug:        domain.TitleToSlug(productInput.Title),
	}

	productOutput, err := usecase.Repository.CreateProduct(&product)
	if err != nil {
		return nil, err
	}

	return productOutput, nil
}
