package routes

import (
	"ginws/config"
	"ginws/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(d *config.Dependencies) *gin.Engine {
	r := gin.Default()

	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.POST("/login", handler.UserLoginHandler(d))
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
