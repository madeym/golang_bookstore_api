package routes

import (
	"belajar-api/controllers"
	"belajar-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var RegisterBookRoutes = func(router *gin.Engine, v *gin.RouterGroup) {
	jwt := router.Group("/v1/jwt")
	v.Use(cors.Default())
	jwt.Use(middleware.Authenticate())
	v.GET("/book", controllers.GetBooks)
	v.GET("/book/:id", controllers.GetBookByID)
	jwt.POST("/book", controllers.CreateBook)
	jwt.POST("/book/:id/borrow", controllers.BorrowBook)
	jwt.POST("/book/:id/buy", controllers.BuyBook)
	jwt.PUT("/book/:id", controllers.UpdateBook)
	jwt.DELETE("/book/:id", controllers.DeleteBook)
}
