package usecases

import (
	domain "discountapp/domain"
	"discountapp/usecases/inputs"
	"discountapp/usecases/interfaces"
	"discountapp/usecases/outputs"
)

type GetDiscount struct {
	Repository interfaces.IRepository
}

func NewGetDiscount(repository interfaces.IRepository) interfaces.IGetDiscount {
	return &GetDiscount{Repository: repository}
}

func (usecase *GetDiscount) Perform(discountInput *inputs.DiscountInput) (*outputs.DiscountOutput, error) {
	product := domain.Product{Slug: discountInput.ProductSlug}
	productOutput, err := usecase.Repository.GetProduct(&product)
	if err != nil {
		return nil, err
	}

	client := domain.Client{Email: discountInput.ClientEmail}

	clientOutput, err := usecase.Repository.GetClient(&client)
	if err != nil {
		return nil, err
	}

	totalPrice, totalDiscount := domain.GetPriceWithDiscount(
		productOutput.Price,
		clientOutput.GetClientAge(),
		nil,
	)

	output := outputs.DiscountOutput{
		ProductTile:          productOutput.Title,
		ProductPrice:         int64(totalPrice),
		OriginalProductPrice: int64(productOutput.Price),
		DiscountPercentage:   totalDiscount,
	}

	return &output, nil
}
