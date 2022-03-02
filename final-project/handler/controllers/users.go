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

func RegisterUser(c *gin.Context) {
	db = infra.GetDB()
	contentType := utils.GetContentType(c)
	User := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Preload("Photos").Preload("Comments").Preload("SocialMedias").Find(&User).Error
	if err != nil {
		log.Fatal(err)
	}

	User.Password = utils.HashPassword(User.Password)

	err = db.Debug().Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
		"password": User.Password,
	})
}

func LoginUser(c *gin.Context) {
	var (
		User     models.User
		password string
		err      error
	)

	db = infra.GetDB()
	contentType := utils.GetContentType(c)
	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err = db.Preload("Photos").Preload("Comments").Preload("SocialMedias").Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	comparePass := utils.ComparePassword([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid password",
		})
		return
	}

	token := utils.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"age":        User.Age,
		"updated_at": User.UpdatedAt,
		"token":      token,
	})
}

func UpdateUser(c *gin.Context) {
	db = infra.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := utils.GetContentType(c)
	OldUser := models.User{}
	NewUser := models.User{}

	paramId, _ := strconv.Atoi(c.Param("userId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&NewUser)
	} else {
		c.ShouldBind(&NewUser)
	}

	err = db.First(&OldUser, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Data Not Found",
			"message": err.Error(),
		})
	}

	NewUser.ID = userId

	err = db.Model(&OldUser).Updates(&NewUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         OldUser.ID,
		"email":      OldUser.Email,
		"username":   OldUser.Username,
		"age":        OldUser.Age,
		"updated_at": OldUser.UpdatedAt,
	})
}

func DeleteUser(c *gin.Context) {
	db = infra.GetDB()
	User := models.User{}

	paramId, _ := strconv.Atoi(c.Param("userId"))

	err := db.Model(&User).Delete(&User, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
