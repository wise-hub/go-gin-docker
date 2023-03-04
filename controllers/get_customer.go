package controllers

import (
	"database/sql"
	r "ginws/repo"

	"github.com/gin-gonic/gin"
)

func GetCustomer(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"customer": id,
	})
	c.Done()
}

func UserLogin(c *gin.Context, db *sql.DB) {
	//userid := c.PostForm("username")
	//message := c.PostForm("password")

	result := r.UserLoginRepo(*&db)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": result,
	})
}
