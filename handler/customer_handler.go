package handler

import (
	"database/sql"
	"ginws/config"
	"ginws/helpers"
	"ginws/model_in"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		cust := &model_in.T{CustomerID: c.Param("id")}

		if err := ValidateMiddleware(c, cust); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		// fetch customer data
		resultSet, err := repository.CustomerRepo(d.Db, cust.CustomerID)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"result": "No results found"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result":        "OK",
			"customer_data": resultSet,
		})

	}
}
