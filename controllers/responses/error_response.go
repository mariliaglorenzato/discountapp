package responses

type ErrorResponse struct {
	Code uint
	Data interface{}
}

func InternalServerError(errors ...string) *ErrorResponse {
	return &ErrorResponse{}
}
