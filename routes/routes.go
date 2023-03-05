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
		api.GET("/customer/:id", handler.GetCustomerHandler(d))
		api.GET("/accounts/:id", handler.GetAccountsHandler(d))
		api.GET("/customer-accounts/:id", handler.GetCustAccHandler(d))
	}

	return r
}
