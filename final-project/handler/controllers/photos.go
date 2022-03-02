package controllers

import (
	"khg-final-project/infra"
	"khg-final-project/models"
	"khg-final-project/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AddPhoto(c *gin.Context) {
	db = infra.GetDB()
	Photo := models.Photo{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	contentType := utils.GetContentType(c)
	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userId

	err = db.Debug().Create(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Title,
		"photo_url":  Photo.PhotoURL,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func GetPhotos(c *gin.Context) {
	db = infra.GetDB()
	Photos := []models.Photo{}

	err = db.Debug().Preload("Comments").Preload("User").Find(&Photos).Error
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, Photos)
}

func UpdatePhoto(c *gin.Context) {
	db = infra.GetDB()
	OldPhoto := models.Photo{}
	NewPhoto := models.Photo{}

	paramId, _ := strconv.Atoi(c.Param("photoId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	contentType := utils.GetContentType(c)
	if contentType == appJson {
		c.ShouldBindJSON(&NewPhoto)
	} else {
		c.ShouldBind(&NewPhoto)
	}

	err = db.First(&OldPhoto, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Data Not Found",
			"message": err.Error(),
		})
	}

	NewPhoto.ID = uint(paramId)
	NewPhoto.UserID = userId

	err = db.Model(&OldPhoto).Updates(&NewPhoto).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         OldPhoto.ID,
		"title":      OldPhoto.Title,
		"caption":    OldPhoto.Caption,
		"photo_url":  OldPhoto.PhotoURL,
		"user_id":    OldPhoto.UserID,
		"updated_at": OldPhoto.UpdatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	db = infra.GetDB()
	Photo := models.Photo{}

	paramId, _ := strconv.Atoi(c.Param("photoId"))

	err = db.First(&Photo, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Data Not Found",
			"message": err.Error(),
		})
	}

	err = db.Model(&Photo).Delete(&Photo, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
