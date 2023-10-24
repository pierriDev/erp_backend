package schemas

import "time"

type UserAdress struct {
	User         User `gorm:"foreignKey:UserID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Adress       string
	Number       string
	Neighborhood string
	cep          string
}

type UserAdressResponse struct {
	ID           uint      `json: "id"`
	CreatedAt    time.Time `json: "createdAt"`
	UpdatedAt    time.Time `json: "updatedAt"`
	DeletedAt    time.Time `json: "deletedAt,omitempty"`
	User         uint      `json: "user_id"`
	Adress       string    `json: "adress"`
	Number       uint      `json: "number"`
	Neighborhood string    `json: "neighborhood"`
	cep          string    `json: "cep"`
}
