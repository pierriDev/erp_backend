package schemas

import (
	"time"

	"gorm.io/gorm"
)

type ProductSupplier struct {
	gorm.Model
	ID       int
	BuyPrice float32

	ProductID int
	Product   Product

	SupplierID int
	Supplier   Supplier
}

type ProductSupplierResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
	BuyPrice   float32   `json:"price"`
	ProductID  int       `json:"productId"`
	SupplierID int       `json:"supplierId"`
}
