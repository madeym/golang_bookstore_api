package middleware

import (
	"belajar-api/helper"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		clientToken := c.Request.Header.Get("Token")
		if clientToken == "" {
			helper.ResponseError(c, "Not Authorized", nil)
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			helper.ResponseError(c, err, nil)
			c.Abort()
			return
		}

		c.Set("name", claims.Name)
		c.Set("email", claims.Email)
		c.Next()
	}
}
