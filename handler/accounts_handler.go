package handler

import (
	"ginws/config"
	"ginws/helpers"
	"ginws/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAccountsHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		// /////////////////////////////////////////
		// AUTHORIZATION
		// /////////////////////////////////////////
		tokenData, err := helpers.ValidateTokenFull(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized (0)"})
			return
		}

		if tokenData.User == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized (1)"})
			return
		}

		if tokenData.ExpDate.Before(time.Now()) {
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

		// handle zero accounts
		if len(accountsList) == 0 {
			c.JSON(http.StatusOK, gin.H{"result": "No accounts found"})
			return
		}

		c.JSON(200, gin.H{"result": "OK", "accountsData": accountsList})
	}
}
