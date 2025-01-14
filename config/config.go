package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	//Initialize the DB
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("Database initialization error: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initialize the logger
	logger = NewLogger(p)
	return logger
}
