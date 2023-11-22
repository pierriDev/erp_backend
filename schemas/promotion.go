package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Promotion struct {
	gorm.Model
	ID        int
	Title     string
	DateStart time.Time
	DateEnd   time.Time
	IsActive  bool
}

type PromotionResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Title     string    `json:"title"`
	DateStart time.Time `json:"dateStart"`
	DateEnd   time.Time `json:"dateEnd"`
	IsActive  bool      `json:"isActive"`
}
