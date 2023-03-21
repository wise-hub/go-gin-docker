package routes

import (
	"ginws/config"
	"ginws/handler"
	"ginws/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(d *config.Dependencies) *gin.Engine {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/login", handler.UserLoginHandler(d))
	r.Use(middleware.LogMiddleware(d))

	{
		api := r.Group("/api")
		api.Use(middleware.AuthMiddleware(d))
		api.Use(middleware.RoleMiddleware(d))

		api.GET("/customer/:id", handler.CustomerHandler(d))
		api.GET("/accounts/:id", handler.AccountsHandler(d))
		api.GET("/customer-accounts/:id", handler.CustAccHandler(d))
		api.GET("/customer-feedback/:id", handler.CustFeedbackReadHandler(d))
		api.POST("/customer-feedback", handler.CustFeedbackAddHandler(d))
	}

	return r
}
