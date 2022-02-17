package main

import (
	"go-api-project/config"
	"go-api-project/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/orders", inDB.CreateOrder)
	router.GET("/orders", inDB.GetOrders)
	router.PUT("/orders/:orderId", inDB.UpdateOrder)
	router.DELETE("/orders/:orderId", inDB.DeleteOrder)
	router.DELETE("/orders", inDB.DeleteTable)
	router.Run(":3000")
}
