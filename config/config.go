package config

import (
	"fmt"
	"go-api-project/structs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connected")

	db.AutoMigrate(structs.Item{}, structs.Order{})
	return db
}
