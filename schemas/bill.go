package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	ID             int
	Title          string
	Value          float32
	Description    string
	BillingDueDate time.Time
	IsPaid         bool
}

type BillResponse struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Title       string    `json:"title"`
	Value       float32   `json:"value"`
	Description string    `json:"description"`
	BillingDay  int       `json:"billingday"`
	IsPaid      bool      `json:"ispaid"`
}
