package handler

import (
	"khg-final-project/infra"
	"khg-final-project/models"
	"khg-final-project/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func RegisterUser(c *gin.Context) {
	// var (
	// 	user   models.User
	// 	result gin.H
	// )

	db := infra.GetDB()
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
	})

	// age := c.PostForm("age")
	// email := c.PostForm("email")
	// password := c.PostForm("password")
	// username := c.PostForm("username")

	// ageUint64, err := strconv.ParseUint(age, 10, 64)
	// if err != nil {
	// 	log.Fatal("Failed to parse to uint64", err)
	// }

	// hashedPassword := utils.HashPassword(password)

	// user.Age = uint(ageUint64)
	// user.Email = email
	// user.Password = hashedPassword
	// user.Username = username
	// user.CreatedAt = time.Now()
	// user.UpdatedAt = time.Time{}

	// err = db.Create(&user).Error
	// if err != nil {
	// 	result = gin.H{
	// 		"result": "Data isn't created",
	// 	}
	// }

	// result = gin.H{
	// 	"status": http.StatusCreated,
	// 	"data":   user,
	// }

	// c.JSON(http.StatusCreated, result)
}

func LoginUser(c *gin.Context) {
	db := infra.GetDB()
	contentType := utils.GetContentType(c)
	User := models.User{}
	password := ""

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	// password = User.Password

	err := db.Preload("Photos").Preload("Comments").Preload("SocialMedias").Find(&User).Error
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
		"token": token,
	})
}
