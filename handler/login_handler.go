package handler

import (
	"fmt"
	"ginws/config"
	"ginws/helpers"
	"ginws/middleware"
	"ginws/model_in"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLoginHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		login := &model_in.InLogin{}

		if err := middleware.ValidateMiddleware(c, login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		if err := middleware.LogMiddleware(d, c, login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		role, err := repository.GetRoleFromUser(d.Db, login.Username)

		if err != nil {
			fmt.Println("role fetch failed")
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"result": "Invalid username or password (1)"})
			return
		}

		res := gin.H{}

		if len(role) > 0 {

			// use this u/p for ldap login -> "einstein", "password"

			_, err := helpers.LdapAuth(d, login.Username, login.Password)

			if err != nil {
				//fmt.Println("Authentication failed: %v\n", err)
				fmt.Println("Authenticaiton failed")
				fmt.Println(err)
				c.JSON(http.StatusOK, gin.H{"result": "Invalid username or password (3)"})
				return
			}

			token, err := middleware.EncryptToken(login.Username, role)
			if err != nil {
				panic(err)
			}

			remoteAddr := helpers.GetRemoteAddr(c)
			tokenInserted, err := repository.InsertNewToken(d, login.Username, token, remoteAddr)

			if err != nil || !tokenInserted {
				fmt.Println("user table update failed")
				c.JSON(http.StatusOK, gin.H{"result": "wtf"})
				return
			}

			res["access_token"] = token
		}

		c.JSON(http.StatusOK, res)
	}
}
