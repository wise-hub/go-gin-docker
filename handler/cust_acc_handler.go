package handler

import (
	"database/sql"
	"fmt"
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
