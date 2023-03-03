package main

import (
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {

	r := gin.Default()

	r.GET("/api/customer/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.JSON(200, gin.H{"customer_endpoint_one": id})
	})

	return r
}

func main() {

	//gin.SetMode(gin.ReleaseMode)
	//
	r := router()
	r.RunTLS(":8443", "./cert/server.pem", "./cert/server.key")
}
