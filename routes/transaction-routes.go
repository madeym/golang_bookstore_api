package routes

import (
	"belajar-api/controllers"
	"belajar-api/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterTransactionRoutes = func(router *gin.Engine, v *gin.RouterGroup) {
	jwt := router.Group("/v1/jwt")
	jwt.Use(middleware.Authenticate())
	jwt.GET("/transaction", controllers.GetTransactions)
	jwt.GET("/transaction/:id", controllers.GetTransactionByID)
}
