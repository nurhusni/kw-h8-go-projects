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

func CreateSocialMedia(c *gin.Context) {
	db = infra.GetDB()
	SocialMedia := models.SocialMedia{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	contentType := utils.GetContentType(c)
	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userId

	err = db.Debug().Create(&SocialMedia).Error

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaURL,
		"user_id":          SocialMedia.UserID,
	})
}

func GetSocialMedias(c *gin.Context) {
	db = infra.GetDB()
	SocialMedias := []models.SocialMedia{}

	err = db.Debug().Preload("User").Find(&SocialMedias).Error
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, SocialMedias)
}

func UpdateSocialMedia(c *gin.Context) {
	db = infra.GetDB()
	OldSocialMedia := models.SocialMedia{}
	NewSocialMedia := models.SocialMedia{}

	paramId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	contentType := utils.GetContentType(c)
	if contentType == appJson {
		c.ShouldBindJSON(&NewSocialMedia)
	} else {
		c.ShouldBind(&NewSocialMedia)
	}

	err = db.First(&OldSocialMedia, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Data Not Found",
			"message": err.Error(),
		})
	}

	NewSocialMedia.ID = uint(paramId)
	NewSocialMedia.UserID = userId

	err = db.Model(&OldSocialMedia).Updates(&NewSocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               OldSocialMedia.ID,
		"name":             OldSocialMedia.Name,
		"social_media_url": OldSocialMedia.SocialMediaURL,
		"user_id":          OldSocialMedia.UserID,
		"updated_at":       OldSocialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db = infra.GetDB()
	SocialMedia := models.SocialMedia{}

	paramId, _ := strconv.Atoi(c.Param("socialMediaId"))

	err = db.First(&SocialMedia, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Data Not Found",
			"message": err.Error(),
		})
	}

	err = db.Model(&SocialMedia).Delete(&SocialMedia, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
