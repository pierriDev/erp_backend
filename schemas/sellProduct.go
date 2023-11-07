package schemas

import (
	"time"

	"gorm.io/gorm"
)

type SellProduct struct {
	gorm.Model
	ID       int
	Quantity int

	SellID    int
	ProductID int

	Sell    Sell    `gorm:"foreignKey:SellID"`
	Product Product `gorm:"foreignKey:ProductID"`
}

type SellProductResponse struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Quantity  int       `json:"quantity"`
	SellID    int       `json:"sellId"`
	ProductID int       `json:"productId"`
}
