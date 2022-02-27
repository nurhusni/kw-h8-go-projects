package main

import (
	"khg-final-project/infra"

	_ "github.com/dgrijalva/jwt-go"
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
// tests -> jwt gp

// h := handler.NewHandler()
// _ = h

// mux := http.NewServeMux()
// mux.HandleFunc("/", h.Employee.Create)

func main() {
	infra.DBInit()
	// commentDB := &handler.CommentHandler{DB: db}
	// router := gin.Default()
}
