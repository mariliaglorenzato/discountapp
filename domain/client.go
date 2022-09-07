package domain

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID        int64     `gorm:"primary_key,autoIncrement"`
	FirstName string    `validate:"required"`
	LastName  string    `validate:"required"`
	BirthDate time.Time `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (c *Client) getClientAge() uint {
	now := time.Now()
	return uint(now.Sub(c.BirthDate))
}
