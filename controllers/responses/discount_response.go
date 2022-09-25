package responses

import "discountapp/usecases/outputs"

type DiscountResponse struct {
	ProductTile          string  `json:"product_title" validate:"required"`
	ProductPrice         float64 `json:"product_price" validate:"required"`
	OriginalProductPrice float64 `json:"original_product_price" validate:"required"`
	DiscountPercentage   float64 `json:"discount_percentage" validate:"required"`
}

func GetDiscountResponse(discountOutput *outputs.DiscountOutput) *DiscountResponse {
	return &DiscountResponse{
		ProductTile:          discountOutput.ProductTile,
		ProductPrice:         float64(discountOutput.ProductPrice) * float64(0.01),
		OriginalProductPrice: float64(discountOutput.OriginalProductPrice) * float64(0.01),
		DiscountPercentage:   discountOutput.DiscountPercentage,
	}
}
