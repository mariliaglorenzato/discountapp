package interfaces

import (
	"discountapp/domain"
	"discountapp/usecases/inputs"
)

type ICreateClient interface {
	Perform(clientInput *inputs.ClientInput) (*domain.Client, error)
}
