package routes

import (
	"ginws/config"
	"ginws/handler"

	"github.com/gin-gonic/gin"
)

func Routes(d *config.Dependencies) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/login", handler.UserLoginHandler(d))
		api.GET("/customer/:id", handler.CustomerHandler(d))
		api.GET("/accounts/:id", handler.AccountsHandler(d))
		api.GET("/customer-accounts/:id", handler.CustAccHandler(d))
	}

	return r
}
