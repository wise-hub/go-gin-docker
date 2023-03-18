package handler

import (
	"fmt"
	"ginws/config"
	"ginws/helpers"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccountsHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		// /////////////////////////////////////////
		// TOKEN AUTHENTICATION
		// /////////////////////////////////////////

		tokenParams, err := ValidateToken(c, d)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"result": err.Error()})
			return
		}
		// ////////////////////////////////////////

		// /////////////////////////////////////////
		// ROLE CHECK
		// /////////////////////////////////////////

		if tokenParams.Role == "ADMIN" {
			fmt.Println("ADMIN ROLE")
		}

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

		// handle zero accounts
		if len(accountsList) == 0 {
			c.JSON(http.StatusOK, gin.H{"result": "No accounts found"})
			return
		}

		c.JSON(200, gin.H{"result": "OK", "accountsData": accountsList})
	}
}
