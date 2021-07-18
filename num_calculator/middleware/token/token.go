package token

import (
	"github.com/gin-gonic/gin"
)

func TokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("User-Token")
		if token != "68b1fdee-1a8e-44ef-98dc-e1250f8e23f0" {
			c.AbortWithStatus(404)
			return
		}
		c.Next()
	}
}
