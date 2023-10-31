package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          int
	Title       string
	Price       float32
	Code        string
	Description string

	CategoryID int
	Category   Category
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Title       string    `json:"title"`
	Price       float32   `json:"price"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	CategoryID  int       `json:"categoryId"`
}
