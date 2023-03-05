package controllers

import (
	"ginws/config"
	"ginws/helpers"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func GetCustomer(c *gin.Context, d *config.Dependencies) {

	if !helpers.IsValidAccessToken(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized"})
		return
	}

	id := c.Param("id")

	re := regexp.MustCompile(`^\d{9}$`)

	res := gin.H{
		"customer": "Not Found",
	}

	// res2 := model.Customer{
	// 	Name:    "Ivan Petrov",
	// 	Address: "addr",
	// 	EGN:     "3244423434",
	// }

	if re.MatchString(id) {
		res["customer"] = id
	}

	c.JSON(200, res)
}
