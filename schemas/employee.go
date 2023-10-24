package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID       int
	Password string
	Picture  string
	UserID   int
	User     User
}

type EmployeeResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Password  string    `json:"adress"`
	Picture   string    `json:"cep"`
}
