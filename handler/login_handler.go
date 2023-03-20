package handler

import (
	"fmt"
	"ginws/config"
	"ginws/helpers"
	"ginws/model_in"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLoginHandler(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {
		login := &model_in.InLogin{}

		if err := c.ShouldBind(&login); err != nil {
			fmt.Println("Not bound JSON")
			c.JSON(http.StatusBadRequest,
				gin.H{"result": helpers.AssertEnvForError(d.Cfg.EnvType, err)})
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
			fmt.Println(err)
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
				//fmt.Println("Authentication failed: %v\n", err)
				fmt.Println("Authenticaiton failed")
				fmt.Println(err)
				c.JSON(http.StatusOK, gin.H{"result": "Invalid username or password (3)"})
				return
			}

			token, err := EncryptToken(login.Username, role)
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

			logInfo := &repository.LogInfo{
				Username:  login.Username,
				IPAddress: helpers.GetRemoteAddr(c),
				Handler:   "login",
				//BodyParams: map[string]interface{}{"customer": id},
			}

			if err := repository.SaveLog(d, logInfo); err != nil {
				fmt.Println("Error logging to Oracle database:", err)
				panic(err)
			}

			res["result"] = "OK"
			res["access_token"] = token
		}

		c.JSON(http.StatusOK, res)
	}
}
