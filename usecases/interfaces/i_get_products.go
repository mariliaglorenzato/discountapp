package interfaces

import (
	"discountapp/usecases/outputs"
)

type IGetProducts interface {
	Perform() (*outputs.ProductsOutput, error)
}
