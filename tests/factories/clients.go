package factories

import (
	"time"

	"discountapp/domain"
)

func BuildClients() []*domain.Client {
	return []*domain.Client{
		{
			ID:        1,
			FirstName: "Maria",
			LastName:  "Oliveira",
			Email:     "email@gmail.com",
			BirthDate: getBirthDate(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
		},
		{
			ID:        2,
			FirstName: "Pedro",
			LastName:  "Silva",
			Email:     "email2@gmail.com",
			BirthDate: getBirthDate(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
		},
	}
}

func BuildStubbedClients() []*domain.Client {
	return []*domain.Client{
		{
			FirstName: "Maria",
			LastName:  "Oliveira",
			Email:     "email@gmail.com",
			BirthDate: getBirthDate(),
		},
		{
			FirstName: "Pedro",
			LastName:  "Silva",
			Email:     "email2@gmail.com",
			BirthDate: getBirthDate(),
		},
	}
}
