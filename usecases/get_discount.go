package usecases

import (
	"fmt"

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
	product := domain.Product{Title: discountInput.ProductTitle}
	productOutput, err := usecase.Repository.GetProduct(&product)
	if err != nil {
		return nil, err
	}

	client := domain.Client{Email: discountInput.ClientEmail}

	clientOutput, err := usecase.Repository.GetClient(&client)
	if err != nil {
		return nil, err
	}

	fmt.Println(clientOutput)

	discount := domain.NewDiscount(clientOutput)

	totalPrice := discount.GetPriceWithDiscount(productOutput.Price)

	output := outputs.DiscountOutput{
		ProductTile:          productOutput.Title,
		ProductPrice:         int64(totalPrice),
		OriginalProductPrice: int64(productOutput.Price),
		DiscountPercentage:   discount.TotalDiscount,
	}

	return &output, nil
}
