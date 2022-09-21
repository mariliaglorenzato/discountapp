package responses

import (
	"discountapp/usecases/outputs"
)

type ClientsResponse struct {
	Clients []*ClientResponse `json:"clients"`
}

type ClientResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
