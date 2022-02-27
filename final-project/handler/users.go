package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (idb *UserHandler) RegisterUser(c *gin.Context) {

}
