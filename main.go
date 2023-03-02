package main

import (
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/customer/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.JSON(200, gin.H{"customer": id})
	})

	return r
}

func main() {
	r := router()
	r.Run("0.0.0.0:44400")
}
