package domain

import (
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint64 `gorm:"primary_key,autoIncrement"`
	Title       string
	Slug        string `gorm:"index:index_products_slug,unique"`
	Description string
	Price       uint64 `validate:"gte=500"`
	TotalPrice  uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
}

func TitleToSlug(title string) string {
	return slug.Make(title)
}
