package persistence

import (
	"fmt"

	"discountapp/domain"

	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Persistence struct {
	DB *gorm.DB
}

func NewPersistence() *Persistence {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// todo: this should not be here, move as soon as possible
	database.AutoMigrate(&domain.Product{}, &domain.Client{})

	return &Persistence{DB: database}
}
