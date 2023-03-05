package handler

import (
	"database/sql"
	"ginws/config"
	"ginws/helpers"
	"ginws/model"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLoginHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {
		login := model.Login{}

		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		if !helpers.IsValidUsername(login.Username) {
			c.JSON(http.StatusOK, gin.H{"result": "Invalid username or password"})
			return
		}

		loginCheck, err := repository.UserLoginRepo(d.Db, login.Username, login.Password)

		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"result": "Invalid username or password"})
				return
			}
			panic(err)
		}

		res := gin.H{
			"result": "Invalid username or password",
		}

		if loginCheck >= 1 {
			token := helpers.GenerateAccessToken()
			// create token service, save it and return it
			// to do

			res["result"] = "OK"
			res["access_token"] = token
		}

		c.JSON(http.StatusOK, res)
	}
}
