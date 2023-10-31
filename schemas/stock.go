package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	ID         int
	Quantity   int
	PriceOfBuy float32

	ProductID  int
	Product    Product
	SupplierID int
	Supplier   Supplier
}

type StockResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
	Quantity   int       `json:"quantity"`
	PriceOfBuy float32   `json:"priceofbuy"`
	ProductID  int       `json:"productId"`
	SupplierID int       `json:"supplierId"`
}
