package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID    int
	Title string
}

type CategoryResponse struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Title     string    `json:"title"`
}
