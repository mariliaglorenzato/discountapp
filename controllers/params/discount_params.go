package params

type DiscountParams struct {
	ProductTitle string `json:"product_title" binding:"required"`
	ClientEmail  string `json:"client_email" binding:"required"`
}
