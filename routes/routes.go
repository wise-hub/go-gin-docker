package routes

import (
	"ginws/config"
	"ginws/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(d *config.Dependencies) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/login", func(c *gin.Context) {
			controllers.UserLogin(c, d)
		})

		api.GET("/customer/:id", func(c *gin.Context) {
			controllers.GetCustomer(c, d)
		})

		api.GET("/accounts/:id", func(c *gin.Context) {
			controllers.GetAccounts(c, d)
		})
	}

	return r
}
