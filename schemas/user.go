package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	CPF   string
	Phone string
	Email string
}

type UserResponse struct {
	ID        uint      `json: "id"`
	CreatedAt time.Time `json: "createdAt"`
	UpdatedAt time.Time `json: "updatedAt"`
	DeletedAt time.Time `json: "deletedAt,omitempty"`
	Name      string    `json: "name"`
	CPF       string    `json: "cpf"`
	Phone     string    `json: "phone"`
	Email     bool      `json: "email"`
}
