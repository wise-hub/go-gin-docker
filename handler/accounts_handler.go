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
		// TOKEN AUTHENTICATION
		// /////////////////////////////////////////

		// OFFLINE - decrypts and checks
		tokenData, err := helpers.ValidateTokenFull(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"result": err.Error()})
			return
		}

		// ONLINE - checks in DB. COMMENT it for fewer DB I/O
		if !repository.ValidateTokenAtDb(d.Db, tokenData.Token) {
			c.JSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized (1)"})
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
