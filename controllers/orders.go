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
		order  structs.Order
		result gin.H
	)

	customerName := c.PostForm("customer_name")
	itemCode := c.PostForm("item_code")
	description := c.PostForm("description")
	quantity := c.PostForm("quantity")

	order.CustomerName = customerName
	order.OrderedAt = time.Now()

	order.Item.ItemCode, _ = strconv.ParseInt(itemCode, 10, 64)
	order.Item.Description = description
	order.Item.Quantity, _ = strconv.ParseInt(quantity, 10, 64)

	err := idb.DB.Create(&order).Error
	if err != nil {
		result = gin.H{
			"result": "Order data isn't created",
		}
	}

	result = gin.H{
		"result": order,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []structs.Order
		result gin.H
	)

	idb.DB.Preload("Item").Find(&orders)

	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
		}
	} else {
		result = gin.H{
			"result": orders,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	var (
		order    structs.Order
		newOrder structs.Order
		result   gin.H
	)

	id := c.Query("id")

	description := c.PostForm("description")
	quantity := c.PostForm("quantity")
	customerName := c.PostForm("customer_name")
	itemCode := c.PostForm("item_code")

	err := idb.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	}

	newOrder.CustomerName = customerName
	newOrder.OrderedAt = time.Now()

	newOrder.Item.ItemCode, _ = strconv.ParseInt(itemCode, 10, 64)
	newOrder.Item.Description = description
	newOrder.Item.Quantity, _ = strconv.ParseInt(quantity, 10, 64)

	assocErr := idb.DB.Model(&order.Item).Updates(&newOrder.Item).Error
	rootErr := idb.DB.Model(&order).Updates(&newOrder).Error
	if rootErr != nil && assocErr != nil {
		result = gin.H{
			"result": "Updating order data failed",
		}
	}

	result = gin.H{
		"result": "Successfully updated data",
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data wasn't found",
		}
	}

	err = idb.DB.Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "Failed to delete data",
		}
	} else {
		result = gin.H{
			"result": "Data was deleted successfully",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteTable(c *gin.Context) {
	var result gin.H

	orderErr := idb.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&[]structs.Order{}).Error

	if orderErr != nil {
		result = gin.H{
			"result": "Deletion is unsuccessful",
		}
	} else {
		result = gin.H{
			"result": "Tables are deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
