package interfaces

import (
	"discountapp/usecases/outputs"
)

type IGetClients interface {
	Perform() (*outputs.ClientsOutput, error)
}
