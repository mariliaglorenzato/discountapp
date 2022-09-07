package payloads

type ClientPayload struct {
	FirstName string `json:"name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
}
