package schemas

import (
	"time"

	"gorm.io/gorm"
)

type PromotionProduct struct {
	gorm.Model
	ID          int
	PromotionID int
	Promotion   Promotion
	Value       float32

	ProductID int
	Product   Product
}

type PromotionProductResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	PromotionID int       `json:"promotionId"`
	ProductID   int       `json:"productId"`
	Value       float32   `json:"value"`
}
