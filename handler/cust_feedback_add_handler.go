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

func CustFeedbackAddHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		/* TOKEN AUTHENTICATION */
		tokenParams, err := ValidateToken(c, d)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"result": err.Error()})
			return
		}

		/* REFRESH TOKEN - FOR MAIN API METHOD */
		if err := repository.UpdateTokenExpiry(d, tokenParams.User); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"result": "Authentication Error"})
			return
		}

		/* ROLE CHECK */
		if tokenParams.Role == "ADMIN" {
			fmt.Println("ADMIN ROLE")
			// do stuff
		}
		//////////////////////////////////////////////////////////////

		// set struct
		custFeedback := &model_in.InCustomerFeedback{}

		// validate keys (fields)
		if err := c.ShouldBind(custFeedback); err != nil {
			fmt.Println("Not bound JSON")
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		// validate values
		validate := validator.New()
		if err := validate.Struct(custFeedback); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		/* LOGGER PRELIMINARY */
		logInfo := &repository.LogInfo{
			Username:   tokenParams.User,
			IPAddress:  helpers.GetRemoteAddr(c),
			Handler:    "CustFeedbackAddHandler",
			BodyParams: map[string]interface{}{"l": custFeedback},
		}

		// insert new feedback record
		if err := repository.InsertCustFeedback(d.Db, custFeedback, tokenParams.User); err != nil {
			errMsg := err.Error()
			logInfo.ErrorInfo = &errMsg
			repository.SaveLog(d, logInfo)
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"result": "Error (0)"})
			}
			panic(err)
		}

		// save action log to database
		if err := repository.SaveLog(d, logInfo); err != nil {
			fmt.Println("Error logging to Oracle database:", err)
			panic(err)
		}

		// return json to client
		c.JSON(http.StatusOK, gin.H{
			"result": "OK",
		})

	}
}
