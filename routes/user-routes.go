package routes

import (
	"belajar-api/controllers"
	"belajar-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var RegisterUserRoutes = func(router *gin.Engine, v *gin.RouterGroup) {
	// v.GET("/user", controllers.GetUsers)
	// v.GET("/user/:id", controllers.GetUserByID)
	v.Use(cors.Default())
	v.POST("/user/login", controllers.LoginUser)
	v.POST("/user/register", controllers.RegisterUser)

	// v.PUT("/user/:id", controllers.UpdateUser)
	jwt := router.Group("/v1/jwt")
	jwt.Use(middleware.Authenticate())
	jwt.GET("user", controllers.DetailUser)
	jwt.PUT("user/change-password", controllers.ChangePasswordUser)
	jwt.PUT("user", controllers.UpdateUser)
	// v.DELETE("/user/:id", controllers.DeleteUser)
}
