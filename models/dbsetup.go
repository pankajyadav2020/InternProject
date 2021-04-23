package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB //db declared global
var err error   //err

func ConnectDataBase() {
	DB, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&Loan{}) //creates table from struct
}
