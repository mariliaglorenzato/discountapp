package outputs

import "discountapp/domain"

type ClientsOutput struct {
	Clients []*domain.Client
}
