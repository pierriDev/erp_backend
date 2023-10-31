package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	ID     int
	UserID int
	User   User
}

type SupplierResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}
