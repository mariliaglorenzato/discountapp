package usecases

import (
	domain "discountapp/domain"
	"discountapp/usecases/interfaces"
)

type GetClients struct {
	Repository interfaces.IRepository
}

func NewGetClients(repository interfaces.IRepository) interfaces.IGetClients {
	return &GetClients{Repository: repository}
}

func (usecase *GetClients) Perform() ([]*domain.Client, error) {
	output, err := usecase.Repository.GetAllClients()
	if err != nil {
		return nil, err
	}

	return output, nil
}
