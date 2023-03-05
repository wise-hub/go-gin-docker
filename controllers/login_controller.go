package controllers

import (
	"ginws/config"
	"ginws/helpers"
	"ginws/model"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context, d *config.Dependencies) {

	login := model.Login{}

	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
		return
	}

	res := gin.H{
		"result": "Invalid username or password",
	}

	if !helpers.IsValidUsername(login.Username) {
		c.JSON(200, res)
		return
	}

	loginCheck, err := repository.UserLoginRepo(d.Db, login.Username, login.Password)
	if err != nil {
		panic(err)
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
