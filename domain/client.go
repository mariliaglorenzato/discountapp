package domain

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID        int64     `gorm:"primaryKey,autoIncrement"`
	FirstName string    `validate:"required"`
	LastName  string    `validate:"required"`
	Email     string    `gorm:"unique"`
	BirthDate time.Time `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (c *Client) GetClientAge() uint {
	now := time.Now()

	return uint(now.Year() - c.BirthDate.Year())
}
