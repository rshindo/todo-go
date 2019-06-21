package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB contails persistent data.
var DB *gorm.DB

// Init opens connection to DB.
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect db")
	}
	DB = db
	return DB
}

// Close connection
func Close() {
	if DB != nil {
		DB.Close()
	}
}
