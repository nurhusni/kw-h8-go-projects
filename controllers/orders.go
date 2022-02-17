package controllers

import (
	"go-api-project/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		items  structs.Items
		orders structs.Orders
		result gin.H
	)

	description := c.PostForm("description")
	quantity := c.PostForm("quantity")
	customerName := c.PostForm("customer_name")
	itemCode := c.PostForm("item_code")

	items.ItemCode, _ = strconv.ParseInt(itemCode, 10, 64)
	items.Description = description
	items.Quantity, _ = strconv.ParseInt(quantity, 10, 64)
	orders.CustomerName = customerName
	orders.OrderedAt = time.Now()

	idb.DB.Create(&items)
	idb.DB.Create(&orders)

	result = gin.H{
		"orderedAt":    orders.OrderedAt,
		"customerName": orders.CustomerName,
		"items":        []structs.Items{items},
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetOrders(c *gin.Context) {

}

func (idb *InDB) UpdateOrder(c *gin.Context) {

}

func (idb *InDB) DeleteOrder(c *gin.Context) {

}
