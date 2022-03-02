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

func CreateComment(c *gin.Context) {
	db = infra.GetDB()
	Comment := models.Comment{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	contentType := utils.GetContentType(c)
	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userId

	err = db.Debug().Create(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func GetComments(c *gin.Context) {
	db = infra.GetDB()
	Comments := []models.Comment{}

	err = db.Debug().Preload("User").Preload("Photo").Find(&Comments).Error
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, Comments)
}

func UpdateComment(c *gin.Context) {
	db = infra.GetDB()
	OldComment := models.Comment{}
	NewComment := models.Comment{}

	paramId, _ := strconv.Atoi(c.Param("commentId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	contentType := utils.GetContentType(c)
	if contentType == appJson {
		c.ShouldBindJSON(&NewComment)
	} else {
		c.ShouldBind(&NewComment)
	}

	err = db.First(&OldComment, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Data Not Found",
			"message": err.Error(),
		})
	}

	err = db.Debug().Preload("User").Preload("Photo").Find(&OldComment).Error
	if err != nil {
		log.Fatal(err)
	}

	NewComment.ID = uint(paramId)
	NewComment.UserID = userId

	err = db.Model(&OldComment).Updates(&NewComment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request for Comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         OldComment.ID,
		"message":    OldComment.Message,
		"title":      OldComment.Photo.Title,
		"caption":    OldComment.Photo.Caption,
		"photo_url":  OldComment.Photo.PhotoURL,
		"user_id":    OldComment.UserID,
		"updated_at": OldComment.UpdatedAt,
	})
}

func DeleteComment(c *gin.Context) {
	db = infra.GetDB()
	Comment := models.Comment{}

	paramId, _ := strconv.Atoi(c.Param("commentId"))

	err = db.First(&Comment, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Data Not Found",
			"message": err.Error(),
		})
	}

	err = db.Model(&Comment).Delete(&Comment, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
