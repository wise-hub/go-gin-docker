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

func GetCustomerHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		/* TOKEN AUTHENTICATION */
		tokenParams, err := ValidateToken(c, d)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"result": err.Error()})
			return
		}

		/* REFRESH TOKEN - FOR MAIN API METHOD */
		errUpdated := repository.UpdateTokenExpiry(d.Db, tokenParams.User)
		if errUpdated != nil {
			fmt.Println(errUpdated)
			c.JSON(http.StatusOK, gin.H{"result": "Authentication Error"})
			return
		}

		/* ROLE CHECK */
		if tokenParams.Role == "ADMIN" {
			fmt.Println("ADMIN ROLE")
			// do stuff
		}

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
