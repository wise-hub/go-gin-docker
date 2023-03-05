package handler

import (
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
		login := &model.Login{}

		if err := c.ShouldBind(&login); err != nil {
			fmt.Println("not bound json")
			c.JSON(http.StatusBadRequest, gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
			return
		}

		if !helpers.IsValidUsername(login.Username) {
			fmt.Println("invalid username")
			c.JSON(http.StatusOK, gin.H{"result": "Invalid username or password (0)"})
			return
		}

		//loginCheck := repository.ValidateUserAtDb(d.Db, login.Username)
		role, err := repository.GetRoleFromUser(d.Db, login.Username)

		if err != nil {
			fmt.Println("role fetch failed")
			c.JSON(http.StatusOK, gin.H{"result": "Invalid username or password (1)"})
			return
		}

		res := gin.H{
			"result": "Invalid username or password (2)",
		}

		if len(role) > 0 {

			// use this u/p for ldap login -> "einstein", "password"

			_, err := helpers.LdapAuth(d, login.Username, login.Password)

			if err != nil {
				fmt.Println("Authentication failed: %v\n", err)
				fmt.Println("auth failed")
				c.JSON(http.StatusOK, gin.H{"result": "Invalid username or password (3)"})
				return
			}
			//remoteAddr := strings.Split(c.Request.RemoteAddr, ":")[0]

			remoteAddr := helpers.GetRemoteAddr(c)
			token, err := repository.InsertNewToken(d.Db, login.Username, role, remoteAddr)

			if err != nil {
				fmt.Println("user table update failed")
				c.JSON(http.StatusOK, gin.H{"result": "wtf"})
				return
			}

			res["result"] = "OK"
			res["access_token"] = token
		}

		c.JSON(http.StatusOK, res)
	}
}
