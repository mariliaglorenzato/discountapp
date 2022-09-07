package interfaces

import "discountapp/domain"

type IGetProducts interface {
	Perform() ([]*domain.Product, error)
}
