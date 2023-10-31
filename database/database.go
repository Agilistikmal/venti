package database

import (
	"github.com/Agilistikmal/venti/handler"
	"github.com/Agilistikmal/venti/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func CreateConnection() *gorm.DB {
	log.Print("Connecting to database...")
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	handler.HandleError(err)
	err = db.AutoMigrate(&model.Product{}, &model.Stock{}, &model.Voucher{})
	handler.HandleError(err)
	DB = db
	log.Print("Connected to database")

	return db
}
