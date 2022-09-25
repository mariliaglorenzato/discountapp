package responses

import (
	"discountapp/domain"
	"discountapp/usecases/outputs"
)

type ProductResponse struct {
	ID          uint64  `json:"id" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Slug        string  `json:"slug" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	CreatedAt   string  `json:"created_at" validate:"required"`
	UpdatedAt   string  `json:"updated_at" validate:"required"`
}

type ProductsResponse struct {
	Products []*ProductResponse `json:"products"`
}

func GetProductsResponse(productsOutput *outputs.ProductsOutput) *ProductsResponse {
	var productsResponse []*ProductResponse
	for _, product := range productsOutput.Products {
		productsResponse = append(productsResponse, &ProductResponse{
			ID:          product.ID,
			Title:       product.Title,
			Slug:        product.Slug,
			Description: product.Description,
			Price:       float64(product.Price) * float64(0.01),
			CreatedAt:   product.CreatedAt.String(),
			UpdatedAt:   product.UpdatedAt.String(),
		})
	}

	return &ProductsResponse{
		Products: productsResponse,
	}
}

func GetProductResponse(product *domain.Product) *ProductResponse {
	return &ProductResponse{
		ID:          product.ID,
		Title:       product.Title,
		Slug:        product.Slug,
		Description: product.Description,
		Price:       float64(product.Price) * float64(0.01),
	}
}
