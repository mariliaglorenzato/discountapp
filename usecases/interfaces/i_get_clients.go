package interfaces

import "discountapp/domain"

type IGetClients interface {
	Perform() ([]*domain.Client, error)
}
