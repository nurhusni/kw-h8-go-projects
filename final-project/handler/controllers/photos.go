package controllers

import (
	"khg-final-project/infra"
	"khg-final-project/models"
	"khg-final-project/utils"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AddPhoto(c *gin.Context) {
	db = infra.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := utils.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err = db.Debug().Create(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func GetPhotos(c *gin.Context) {
	db = infra.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// contentType := utils.GetContentType(c)
	// userId := uint(userData["id"].(float64))

	Photos := []models.Photo{}
	err = db.Debug().Preload("Comments").Preload("User").Find(&Photos).Error
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, Photos)
}

func UpdatePhoto(c *gin.Context) {

}

func DeletePhoto(c *gin.Context) {

}
