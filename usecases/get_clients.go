package usecases

import (
	"discountapp/usecases/interfaces"
	"discountapp/usecases/outputs"
)

type GetClients struct {
	Repository interfaces.IRepository
}

func NewGetClients(repository interfaces.IRepository) interfaces.IGetClients {
	return &GetClients{Repository: repository}
}

func (usecase *GetClients) Perform() (*outputs.ClientsOutput, error) {
	output, err := usecase.Repository.GetAllClients()
	if err != nil {
		return nil, err
	}

	clientsOutput := &outputs.ClientsOutput{
		Clients: output,
	}

	return clientsOutput, nil
}
