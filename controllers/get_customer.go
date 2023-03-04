package controllers

import "github.com/gin-gonic/gin"

func GetCustomer(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"customer": id,
	})
	c.Done()
}
