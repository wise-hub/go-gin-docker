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

func CustomerHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		/* TOKEN AUTHENTICATION */
		tokenParams, err := ValidateToken(c, d)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"result": err.Error()})
			return
		}

		/* REFRESH TOKEN - FOR MAIN API METHOD */
		if err := repository.UpdateTokenExpiry(d, tokenParams.User); err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "Authentication Error"})
			return
		}

		/* ROLE CHECK */
		if tokenParams.Role == "ADMIN" {
			fmt.Println("ADMIN ROLE")
			// do stuff
		}
		//////////////////////////////////////////////////////////////

		// validate customer id
		id := c.Param("id")

		/* LOGGER PRELIMINARY */
		logInfo := &repository.LogInfo{
			Username:   tokenParams.User,
			IPAddress:  helpers.GetRemoteAddr(c),
			Handler:    "GetCustomerHandler",
			BodyParams: map[string]interface{}{"customer": id},
		}
		// end logger

		if !helpers.IsValidCustomerID(id) {
			errMsg := "Invalid customer ID"
			logInfo.ErrorInfo = &errMsg
			repository.SaveLog(d, logInfo)
			c.JSON(http.StatusOK, gin.H{"result": &errMsg})
			return
		}

		// fetch customer data
		customerData, err := repository.CustomerRepo(d.Db, id)
		if err != nil {
			errMsg := err.Error()
			logInfo.ErrorInfo = &errMsg
			repository.SaveLog(d, logInfo)
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"result": "No customer found"})
			}
			panic(err)
		}

		if err := repository.SaveLog(d, logInfo); err != nil {
			fmt.Println("Error logging to Oracle database:", err)
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"result":       "OK",
			"customerData": customerData,
		})

	}
}
