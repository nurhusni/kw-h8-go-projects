package main

import (
	"go-api-project/config"
	"go-api-project/controllers"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	_ = inDB
}
