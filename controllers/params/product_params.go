package params

type ProductParams struct {
	Title string `json:"title" binding:"required"`
}
