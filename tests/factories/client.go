package factories

import (
	"time"

	"discountapp/domain"
	"discountapp/utils"
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
	birthDate, _ := time.Parse(utils.DateFormat, "29/01/1980")
	return birthDate
}
