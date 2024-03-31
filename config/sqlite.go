package config

import (
	"os"

	"github.com/CintiaSilva7300/go-opportunity/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {

	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"
	//Check if DB exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("SQLite database not found. Creating a new one...")

		//Create DB file and directory
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

	//Create DB end connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("SQLite connection error: %v", err)
		return nil, err
	}

	//Migrate schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("SQLite auto migration error: %v", err)
		return nil, err
	}

	return db, nil
}
