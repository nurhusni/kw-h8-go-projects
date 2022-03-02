package main

import (
	controllers "khg-final-project/handler/controllers"
	middlewares "khg-final-project/handler/middlewares"
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

		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:userId", middlewares.UserAuthorization(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.UserAuthorization(), controllers.DeleteUser)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", middlewares.PhotoAuthorization(), controllers.AddPhoto)
		photoRouter.GET("/", middlewares.UserAuthorization(), controllers.GetPhotos)
		photoRouter.PUT("/:userId", middlewares.UserAuthorization(), controllers.UpdateUser)
		photoRouter.DELETE("/:userId", middlewares.UserAuthorization(), controllers.DeleteUser)
	}

	router.Run(PORT)
}
