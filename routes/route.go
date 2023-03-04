package routes

import (
	c "ginws/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.GET("/api/customer/:id", c.GetCustomer)

	return r
}
