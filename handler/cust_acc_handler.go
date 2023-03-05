package handler

import (
	"database/sql"
	"ginws/config"
	"ginws/helpers"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustAccHandler(d *config.Dependencies) gin.HandlerFunc {
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

		// fetch customer data
		custAccData, err := repository.GetCustAccRepo(d.Db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"result": "No customer found"})
				return
			}
			panic(err)
		}

		res := gin.H{
			"result":          "OK",
			"custAccountData": custAccData,
		}

		c.JSON(http.StatusOK, res)

	}
}
