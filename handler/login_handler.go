package handler

import (
	"database/sql"
	"fmt"
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

		if !helpers.IsValidUsername(*&login.Username) {
			c.JSON(http.StatusOK, gin.H{"result": "Invalid username or password"})
			return
		}

		loginCheck, err := repository.UserLoginRepo(d.Db, *&login.Username, *&login.Password)

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

			// "einstein", "password"

			userDN, err := helpers.LdapAuth(*&login.Username, *&login.Password)

			if err != nil {
				fmt.Printf("Authentication failed: %v\n", err)
			} else {
				fmt.Printf("Authenticated successfully as %s\n", userDN)
			}

			fmt.Println(userDN)

			token := helpers.GenerateAccessToken()

			res["result"] = "OK"
			res["access_token"] = token
		}

		c.JSON(http.StatusOK, res)
	}
}
