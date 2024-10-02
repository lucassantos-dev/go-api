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
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initialization sqlite: %v", err)
	}
	return nil
}

func GetSqlite() *gorm.DB {
	return db
}
func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
