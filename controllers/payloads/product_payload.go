package payloads

type ProductPayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       uint64 `json:"price" validate:"required"`
}
