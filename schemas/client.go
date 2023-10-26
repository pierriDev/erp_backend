package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID     int
	UserID int
	User   User
}

type ClientResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}
