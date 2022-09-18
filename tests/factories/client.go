package factories

import (
	"time"

	"discountapp/domain"
)

func BuildClient() *domain.Client {
	return &domain.Client{
		ID:        1,
		FirstName: "Maria",
		LastName:  "Oliveira",
		Email:     "email@gmail.com",
		BirthDate: getBirthDate(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}
}

func BuildStubbedClient() *domain.Client {
	return &domain.Client{
		FirstName: "Maria",
		LastName:  "Oliveira",
		Email:     "email@gmail.com",
		BirthDate: getBirthDate(),
	}
}

func getBirthDate() time.Time {
	birthDate, _ := time.Parse("2006-01-02", "1980-01-29")
	return birthDate
}
