package handler

import (
	"ginws/config"
	"ginws/helpers"
	"ginws/model_in"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccountsHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		cust := &model_in.T{CustomerID: c.Param("id")}

		if err := ValidateMiddleware(c, cust); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		// fetch accounts
		resultSet, err := repository.AccountsRepo(d.Db, cust.CustomerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		// handle zero accounts
		if len(resultSet) == 0 {
			c.JSON(http.StatusOK, gin.H{"result": "No results found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result":       "OK",
			"accountsData": resultSet,
		})
	}
}
