package router

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)

		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:userId", middlewares.UserAuthorization(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.UserAuthorization(), controllers.DeleteUser)
	}

	return r
}
