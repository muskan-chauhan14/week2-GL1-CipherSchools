package database

import (
	"log"
	_ "github.com/jinzhu/gorm/dialects/postgres" //import a package wihtout using library
	"github.com/jinzhu/gorm"
	"github.com/muskan-chauhan14/week2-GL1-CipherSchools/models"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := "postgres"
	password := "postgres"
	args := "host=" + host + " port=" + port + " user=" + username + " dbnme=" + dbName + " sslmode=disable password=" + password
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db
}
