package handler

import (
	"khg-final-project/infra"
	"khg-final-project/models"
	"khg-final-project/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func RegisterUser(c *gin.Context) {
	var (
		user   models.User
		result gin.H
	)

	db := infra.GetDB()

	age := c.PostForm("age")
	email := c.PostForm("email")
	password := c.PostForm("password")
	username := c.PostForm("username")

	ageUint64, err := strconv.ParseUint(age, 10, 64)
	if err != nil {
		log.Fatal("Failed to parse to uint64", err)
	}

	hashedPassword := utils.HashPassword(password)

	user.Age = uint(ageUint64)
	user.Email = email
	user.Password = hashedPassword
	user.Username = username
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Time{}

	err = db.Preload("Photos").Preload("Comments").Preload("SocialMedias").Find(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Create(&user).Error
	if err != nil {
		result = gin.H{
			"result": "Data isn't created",
		}
	}

	result = gin.H{
		"status": http.StatusCreated,
		"data":   user,
	}

	c.JSON(http.StatusCreated, result)
}
