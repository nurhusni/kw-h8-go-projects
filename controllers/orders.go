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
		// item   structs.Items
		order  structs.Orders
		result gin.H
	)

	customerName := c.PostForm("customer_name")
	itemCode := c.PostForm("item_code")
	description := c.PostForm("description")
	quantity := c.PostForm("quantity")

	order.CustomerName = customerName
	order.OrderedAt = time.Now()
	order.Items.ItemCode, _ = strconv.ParseInt(itemCode, 10, 64)
	order.Items.Description = description
	order.Items.Quantity, _ = strconv.ParseInt(quantity, 10, 64)

	idb.DB.Create(&order)

	result = gin.H{
		"order": order,
		// "item":  item,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		allOrders []structs.Orders
		// allItems  []structs.Items
		result gin.H
	)

	idb.DB.Find(&allOrders)
	// idb.DB.Find(&allItems)

	if len(allOrders) < 1 {
		result = gin.H{
			"result": nil,
		}
	} else {
		result = gin.H{
			"result": allOrders,
			// "orders": allOrders,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	var (
		order    structs.Orders
		newOrder structs.Orders
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

	order.CustomerName = customerName
	order.OrderedAt = time.Now()
	order.Items.ItemCode, _ = strconv.ParseInt(itemCode, 10, 64)
	order.Items.Description = description
	order.Items.Quantity, _ = strconv.ParseInt(quantity, 10, 64)

	err = idb.DB.Model(&order).Updates(&newOrder).Error
	if err != nil {
		result = gin.H{
			"result": "Update failed",
		}
	} else {
		result = gin.H{
			"result": "Successfully updated data",
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		order  structs.Orders
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
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
	idb.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&[]structs.Items{})
	idb.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&[]structs.Orders{})
}
