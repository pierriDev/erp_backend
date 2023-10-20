package schemas

import (
	"gorm.io/gorm"
)

// DEFINICAO DOS CAMPOS DO BANCO
type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
