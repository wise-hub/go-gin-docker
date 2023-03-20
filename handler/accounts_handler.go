package handler

import (
	"fmt"
	"ginws/config"
	"ginws/helpers"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccountsHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		/* TOKEN AUTHENTICATION */
		tokenParams, err := ValidateToken(c, d)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"result": err.Error()})
			return
		}

		/* ROLE CHECK */
		if tokenParams.Role == "ADMIN" {
			fmt.Println("ADMIN ROLE")
			// do stuff
		}

		// validate customer id
		id := c.Param("id")

		/* LOGGER PRELIMINARY */
		logInfo := &repository.LogInfo{
			Username:   tokenParams.User,
			IPAddress:  helpers.GetRemoteAddr(c),
			Handler:    "AccountsHandler",
			BodyParams: id,
		}
		// end logger

		if !helpers.IsValidCustomerID(id) {
			errMsg := "Invalid customer ID"
			logInfo.ErrorInfo = &errMsg
			repository.SaveLog(d, logInfo)
			c.JSON(http.StatusOK, gin.H{"result": "Invalid customer ID"})
			return
		}
		//////////////////////////////////////////////////////////////

		// fetch accounts
		accountsList, err := repository.AccountsRepo(d.Db, id)
		if err != nil {
			errMsg := err.Error()
			logInfo.ErrorInfo = &errMsg
			repository.SaveLog(d, logInfo)
			panic(err)
		}

		// handle zero accounts
		if len(accountsList) == 0 {
			errMsg := "No accounts foundD"
			logInfo.ErrorInfo = &errMsg
			repository.SaveLog(d, logInfo)
			c.JSON(http.StatusOK, gin.H{"result": errMsg})
			return
		}

		if err := repository.SaveLog(d, logInfo); err != nil {
			fmt.Println("Error logging to Oracle database:", err)
			panic(err)
		}

		c.JSON(200, gin.H{"result": "OK", "accountsData": accountsList})
	}
}
