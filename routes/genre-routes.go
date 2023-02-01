package routes

import (
	"belajar-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var RegisterGenreRoutes = func(router *gin.Engine, v *gin.RouterGroup) {
	// jwt := router.Group("v1/jwt")
	// jwt.Use(middleware.Authenticate())
	v.Use(cors.Default())
	v.GET("/genre", controllers.GetGenres)
	v.GET("/genre/:id", controllers.GetGenreByID)
	v.POST("/genre", controllers.CreateGenre)
	v.PUT("/genre/:id", controllers.UpdateGenre)
	v.DELETE("/genre/:id", controllers.DeleteGenre)
}
