package handler

import (
	"database/sql"
	"ginws/config"
	"ginws/helpers"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomerHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		// /////////////////////////////////////////
		// AUTHORIZATION
		// /////////////////////////////////////////
		token := helpers.FetchValidTokenOffline(c)

		if token == "0" {
			c.JSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized (1)"})
			return
		}
		if !repository.ValidateTokenAtDb(d.Db, token) {
			c.JSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized (2)"})
			return
		}
		// ////////////////////////////////////////

		// validate customer id
		id := c.Param("id")
		if !helpers.IsValidCustomerID(id) {
			c.JSON(http.StatusOK, gin.H{"result": "Invalid customer ID"})
			return
		}

		// fetch customer data
		customerData, err := repository.GetCustomerRepo(d.Db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"result": "No customer found"})
				return
			}
			panic(err)
		}

		res := gin.H{
			"result":       "OK",
			"customerData": customerData,
		}

		c.JSON(http.StatusOK, res)

	}
}
