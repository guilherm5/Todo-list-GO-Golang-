package database

import (
	"fmt"
	"log"
	"os"

	"github.com/guilherm5/crud-gingonic-mvc/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func OpenDB(DB *gorm.DB) Handler {
	return Handler{DB}
}

func Init() *gorm.DB {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err.Error(), " Error in load configs files DB")
	}

	HOST := os.Getenv("HOST")
	PORTBD := os.Getenv("PORTBD")
	USER := os.Getenv("USER")
	PASSWORD := os.Getenv("PASSWORD")
	DBNAME := os.Getenv("DBNAME")

	stringConnectionDB := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s", HOST, PORTBD, USER, PASSWORD, DBNAME)

	db, err := gorm.Open(postgres.Open(stringConnectionDB), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error(), " error in init database")
	}

	db.AutoMigrate(&models.Todo_List{})
	return db

}
