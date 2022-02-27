package handler

import (
	"khg-final-project/entity"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (userDB *UserHandler) RegisterUser(c *gin.Context) {
	var (
		user   entity.User
		result gin.H
	)

	age := c.PostForm("age")
	email := c.PostForm("email")
	password := c.PostForm("password")
	username := c.PostForm("username")

	ageUint64, err := strconv.ParseUint(age, 10, 64)
	if err != nil {
		log.Fatal("Failed to parse to uint64", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal("Failed to hash password", err)
	}

	user.Age = uint(ageUint64)
	user.Email = email
	user.Password = string(hashedPassword)
	user.Username = username
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = userDB.DB.Preload("Photos").Preload("Comments").Preload("SocialMedias").Find(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	err = userDB.DB.Create(&user).Error
	if err != nil {
		result = gin.H{
			"result": "Data isn't created",
		}
	}

	result = gin.H{
		"result": user,
	}

	c.JSON(http.StatusOK, result)
}
