package params

type ProductParams struct {
	Slug string `json:"slug" binding:"required"`
}
