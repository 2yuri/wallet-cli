package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"github.com/hyperyuri/wallet-cli/pkg/gorm/models"
)

var DB *gorm.DB

func init(){
	db, err := gorm.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatal("F caiu o banco")
	}

	DB = db

	DB.AutoMigrate(&models.Wallet{})
	DB.AutoMigrate(&models.Address{})
}