package config

import (
	"os"

	"github.com/pierriDev/erp_backend.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"
	//Check if the database file existes
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database not created, creating it...")
		// Create the fatabase file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}
	// Create and connect DB
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrorF("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(
		&schemas.Opening{},
		&schemas.User{},
		&schemas.Address{},
		&schemas.Employee{},
		&schemas.Client{},
		&schemas.Category{},
		&schemas.Product{},
		&schemas.Stock{},
		&schemas.Supplier{},
	)
	if err != nil {
		logger.ErrorF("sqlite auto migration error: %v", err)
		return nil, err
	}

	return db, nil
}
