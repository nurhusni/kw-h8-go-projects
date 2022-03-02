package main

import (
	"khg-final-project/handler/controllers"
	"khg-final-project/handler/middlewares"
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
		photoRouter.POST("/", controllers.AddPhoto)
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComments)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.AddSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedias)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	router.Run(PORT)
}
