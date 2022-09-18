package responses

type DiscountResponse struct {
	ProductTile          string  `json:"product_title" validate:"required"`
	ProductPrice         int64   `json:"product_price" validate:"required"`
	OriginalProductPrice int64   `json:"original_product_price" validate:"required"`
	DiscountPercentage   float64 `json:"discount_percentage" validate:"required"`
}
