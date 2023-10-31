package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         int
	Name       string
	NationalID string
	Phone      string
	Email      string
	AddressID  int
	Address    Address
}

type UserResponse struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
	Name       string    `json:"name"`
	NationalID string    `json:"nationalId"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
}
