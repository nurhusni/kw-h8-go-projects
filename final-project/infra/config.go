package infra

import (
	"fmt"
	"log"

	"khg-final-project/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbName   = "postgres"
	port     = "5432"
	db       *gorm.DB
	err      error
)

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connected")

	db.AutoMigrate(entity.User{}, entity.Photo{}, entity.Comment{}, entity.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
