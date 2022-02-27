package main

import (
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

// h := handler.NewHandler()
// _ = h

// mux := http.NewServeMux()
// mux.HandleFunc("/", h.Employee.Create)

func main() {
	const PORT = ":5432"

	infra.StartDB()
	// commentDB := &handler.CommentHandler{DB: db}
	router := gin.Default()

	router.Run(PORT)
}
