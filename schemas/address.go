package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID           int
	Adress       string
	Number       uint
	Neighborhood string
	CEP          string
	City         string
	State        string
	Country      string
}

type AdressResponse struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt,omitempty"`
	Adress       string    `json:"adress"`
	Number       uint      `json:"number"`
	Neighborhood string    `json:"neighborhood"`
	CEP          string    `json:"cep"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Country      string    `json:"country"`
}
