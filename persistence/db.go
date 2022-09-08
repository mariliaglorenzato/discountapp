package persistence

import (
	"fmt"

	"discountapp/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Persistence struct {
	DB *gorm.DB
}

func NewPersistence() *Persistence {
	database, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	database.AutoMigrate(&domain.Product{}, &domain.Client{})

	return &Persistence{DB: database}
}
