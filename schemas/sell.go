package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Sell struct {
	gorm.Model
	ID          int
	TotalValue  float32
	LiquidValue float32
	Status      string

	PaymentMethodID int
	ClientID        int
	EmployeeID      int

	PaymentMethod PaymentMethod
	Client        Client
	Employee      Employee
}

type SellResponse struct {
	ID              int       `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	DeletedAt       time.Time `json:"deletedAt,omitempty"`
	TotalValue      float32   `json:"totalValue"`
	LiquidValue     float32   `json:"liquidValue"`
	Status          string    `json:"status"`
	PaymentMethodID int       `json:"paymentMethodId"`
	ClientID        int       `json:"clientId"`
	EmployeeID      int       `json:"employeeId"`
}
