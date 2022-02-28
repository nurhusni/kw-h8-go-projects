package main

import (
	"khg-final-project/handler"
	"khg-final-project/infra"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	_ "golang.org/x/crypto/bcrypt"
	_ "gorm.io/gorm"
)

// cmd -> main.go
// entity -> struct, data type, method dari struct
// handler -> isinya hanya untuk handler (rest api, grpc, consumer, middleware, dll)
// infra -> config, env, secret, init db, dll
// repo -> pemanggilan ke database atau ke service lain
// usecase -> business logic (validasi user)
// utils -> helper (jwt go, generate token, validate token)
// tests -> jwt go

func main() {
	const PORT = ":3000"

	infra.StartDB()

	db := infra.GetDB()
	userDB := &handler.UserHandler{DB: db}
	router := gin.Default()

	router.POST("/users/register", userDB.RegisterUser)
	router.Run(PORT)
}
