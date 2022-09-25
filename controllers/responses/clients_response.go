package responses

import (
	"discountapp/domain"
	"discountapp/usecases/outputs"
)

type ClientResponse struct {
	ID        int64  `json:"id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
	CreatedAt string `json:"created_at" validate:"required"`
	UpdatedAt string `json:"updated_at" validate:"required"`
}

type ClientsResponse struct {
	Clients []*ClientResponse `json:"clients"`
}

func GetClientsResponse(clientsOutput *outputs.ClientsOutput) *ClientsResponse {
	var clientsResponse []*ClientResponse
	for _, client := range clientsOutput.Clients {
		clientsResponse = append(clientsResponse, &ClientResponse{
			ID:        client.ID,
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Email:     client.Email,
			BirthDate: client.BirthDate.String(),
			CreatedAt: client.CreatedAt.String(),
			UpdatedAt: client.UpdatedAt.String(),
		})
	}

	return &ClientsResponse{
		Clients: clientsResponse,
	}
}

func GetClientResponse(client *domain.Client) *ClientResponse {
	return &ClientResponse{
		ID:        client.ID,
		FirstName: client.FirstName,
		LastName:  client.LastName,
		Email:     client.Email,
		BirthDate: client.BirthDate.String(),
		CreatedAt: client.CreatedAt.String(),
		UpdatedAt: client.UpdatedAt.String(),
	}
}
