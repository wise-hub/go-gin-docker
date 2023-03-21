package routes

import (
	"ginws/config"
	"ginws/handler"

	"github.com/gin-gonic/gin"
)

func Routes(d *config.Dependencies) *gin.Engine {

	r := gin.Default()
	r.Use(handler.CORSMiddleware())
	r.POST("/api/login", handler.UserLoginHandler(d))
	r.Use(handler.LogMiddleware(d))

	api := r.Group("/api")
	api.Use(handler.AuthMiddleware(d))
	api.Use(handler.RoleMiddleware(d))

	{
		api.GET("/customer/:id", handler.CustomerHandler(d))
		api.GET("/accounts/:id", handler.AccountsHandler(d))
		api.GET("/customer-accounts/:id", handler.CustAccHandler(d))
		// cust feedback
		api.POST("/customer-feedback", handler.CustFeedbackAddHandler(d))
		api.GET("/customer-feedback/:id", handler.CustFeedbackReadHandler(d))
		// others below
	}

	return r
}
