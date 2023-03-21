package middleware

import (
	"ginws/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetString("role") == "RANDOMxxx" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": "Insufficient Rights (0)"})
			return
		}

		// add more logic
		c.Next()

	}
}
