package handler

import (
	"ginws/config"
	"ginws/helpers"
	"ginws/model_in"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustFeedbackAddHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		// set struct
		custFeedback := &model_in.InCustomerFeedback{}

		if err := ValidateMiddleware(c, custFeedback); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		// insert new feedback record
		if err := repository.InsertCustFeedback(d.Db, custFeedback, c.GetString("username")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		// return json to client
		c.JSON(http.StatusOK, gin.H{
			"result": "OK",
		})

	}
}
