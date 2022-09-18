package params

type DiscountParams struct {
	ProductSlug string `json:"product_title" binding:"required"`
	ClientEmail string `json:"client_email" binding:"required"`
}
