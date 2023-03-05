package handler

import (
	"ginws/config"
	"ginws/helpers"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccountsHandler(d *config.Dependencies) gin.HandlerFunc {
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

		// fetch accounts
		accountsList, err := repository.GetAccountsRepo(d.Db, id)
		if err != nil {
			panic(err)
		}

		if len(accountsList) == 0 {
			c.JSON(http.StatusOK, gin.H{"result": "No accounts found"})
			return
		}

		c.JSON(200, gin.H{"result": "OK", "accountsData": accountsList})
	}
}
