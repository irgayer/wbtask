package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"wb/models"
)

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	env := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(sqlite.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//err = DB.Migrator().DropTable(&models.User{}, &models.Comment{})
	err = DB.AutoMigrate(&models.User{}, &models.Comment{})
	if err != nil {
		log.Fatal(err)
	}

}
