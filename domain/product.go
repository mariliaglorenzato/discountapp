package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint64 `gorm:"primary_key,autoIncrement"`
	Title       string
	Description string
	Price       uint64 `validate:"gte=5"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
}
