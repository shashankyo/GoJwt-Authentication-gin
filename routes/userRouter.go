package routes

import (
	controller "Go-JwtAuthentication/controllers"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser)
}
