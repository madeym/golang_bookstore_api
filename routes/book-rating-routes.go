package routes

import (
	"belajar-api/controllers"
	"belajar-api/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterBookRatingRoutes = func(router *gin.Engine, v *gin.RouterGroup) {
	jwt := router.Group("v1/jwt")
	jwt.Use(middleware.Authenticate())

	v.GET("/rating/:id", controllers.GetRatings)
	jwt.POST("/rating", controllers.CreateBookRating)
}
