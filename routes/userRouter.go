package routes

import (
	controllers "go-jwt-mongo/controllers"
	"go-jwt-mongo/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/getAllUsers", controllers.GetAllUsers())
	incomingRoutes.GET("getUserByID/:user_id", controllers.GetUserByID())

}
