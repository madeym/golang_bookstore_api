package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes() {
	router := gin.Default()

	v1 := router.Group("/v1")

	RegisterBookRoutes(router, v1)
	RegisterGenreRoutes(router, v1)
	RegisterUserRoutes(router, v1)
	RegisterTransactionRoutes(router, v1)
	RegisterBookBorrowRoutes(router, v1)
	RegisterBookRatingRoutes(router, v1)
	router.Run()
}
