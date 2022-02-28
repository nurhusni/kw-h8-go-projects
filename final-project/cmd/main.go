package main

import (
	controllers "khg-final-project/handler/controllers"
	"khg-final-project/infra"

	"github.com/gin-gonic/gin"
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
	const PORT = ":3030"

	infra.StartDB()

	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	router.Run(PORT)
}
