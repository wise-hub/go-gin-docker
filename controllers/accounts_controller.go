package controllers

import (
	"ginws/config"
	"ginws/helpers"
	"ginws/model"
	"ginws/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccounts(c *gin.Context, d *config.Dependencies) {

	if !helpers.IsValidAccessToken(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized"})
		return
	}

	id := c.Param("id")

	if !helpers.IsValidCustomerID(id) {
		c.JSON(http.StatusOK, gin.H{"result": "Invalid customer ID"})
		return
	}

	accountsList, err := repository.GetAccountsRepo(d.Db, id)
	if err != nil {
		panic(err)
	}

	res := gin.H{}

	if len(accountsList) >= 1 {
		var accounts []model.Account
		for _, account := range accountsList {
			accountObj := model.Account{
				IBAN:    account.IBAN,
				Balance: account.Balance,
			}
			accounts = append(accounts, accountObj)
		}

		res["result"] = "OK"
		res["accounts"] = accounts
	} else {
		res["result"] = "No records found"
		//fmt.Println(len(accountsList))
	}

	c.JSON(http.StatusOK, res)

}
