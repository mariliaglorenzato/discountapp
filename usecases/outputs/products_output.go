package outputs

import "discountapp/domain"

type ProductsOutput struct {
	Products []*domain.Product
}
