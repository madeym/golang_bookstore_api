package routes

import (
	"belajar-api/controllers"
	"belajar-api/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterBookBorrowRoutes = func(router *gin.Engine, v *gin.RouterGroup) {
	jwt := router.Group("/v1/jwt")
	jwt.Use(middleware.Authenticate())
	jwt.GET("/book/borrow", controllers.GetBookBorrows)
	jwt.PUT("/book/borrow/:id/return-book", controllers.ReturnBookBorrow)
	// jwt.POST("/book", controllers.CreateBook)
	// jwt.POST("/book/:id/borrow", controllers.BorrowBook)
	// jwt.POST("/book/:id/buy", controllers.BuyBook)
	// jwt.PUT("/book/:id", controllers.UpdateBook)
	// jwt.DELETE("/book/:id", controllers.DeleteBook)
}
