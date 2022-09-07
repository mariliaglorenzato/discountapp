package usecases

import (
	domain "discountapp/domain"
	"discountapp/usecases/inputs"
	"discountapp/usecases/interfaces"
)

type CreateClient struct {
	Repository interfaces.IRepository
}

func NewCreateClient(repository interfaces.IRepository) interfaces.ICreateClient {
	return &CreateClient{Repository: repository}
}

func (usecase *CreateClient) Perform(clientInput *inputs.ClientInput) (*domain.Client, error) {
	client := domain.Client{
		FirstName: clientInput.FirstName,
		LastName:  clientInput.LastName,
		BirthDate: clientInput.BirthDate,
	}

	_, err := usecase.Repository.CreateClient(&client)
	if err != nil {
		return nil, err
	}

	return &client, nil
}
