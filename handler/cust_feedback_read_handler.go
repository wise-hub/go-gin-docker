package handler

import (
	"database/sql"
	"fmt"
	"ginws/config"
	"ginws/helpers"
	"ginws/model_in"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CustFeedbackReadHandler(d *config.Dependencies) gin.HandlerFunc {
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

		// set the parameters from input data
		id := &model_in.T{}
		id.CustomerID = c.Param("id")

		// validate input based on struct rules
		validate := validator.New()
		if err := validate.Struct(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		/* LOGGER PRELIMINARY */
		logInfo := &repository.LogInfo{
			Username:   tokenParams.User,
			IPAddress:  helpers.GetRemoteAddr(c),
			Handler:    "customer-feedback-read",
			BodyParams: id,
		}
		// end logger

		// fetch customer data
		customerData, err := repository.ReadCustFeedback(d.Db, id.CustomerID)
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
