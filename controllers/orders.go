package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) CreateOrder(c *gin.Context) {

}

func (idb *InDB) GetOrders(c *gin.Context) {

}

func (idb *InDB) UpdateOrder(c *gin.Context) {

}

func (idb *InDB) DeleteOrder(c *gin.Context) {

}
