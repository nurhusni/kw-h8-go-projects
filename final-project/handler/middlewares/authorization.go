package middlewares

import (
	"khg-final-project/infra"
	"khg-final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := infra.GetDB()
		paramId, err := strconv.Atoi(c.Param("userId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		User := models.User{}

		err = db.Select("id").First(&User, uint(paramId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if User.ID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unathorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
