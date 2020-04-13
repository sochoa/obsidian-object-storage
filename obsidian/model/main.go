package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Initialize(cfg config.ObjectStorageConfig) {
	db, err := gorm.Open("sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Object{})
}
