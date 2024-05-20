package routes

import (
	contollers "go-jwt-mongo/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", contollers.Signup())
	incomingRoutes.POST("/users/login", contollers.Login())

}
