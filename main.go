package main

import (
	"fmt"
	"ginws/config"
	c "ginws/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/sijms/go-ora/v2"
)

func main() {

	d, err := config.Init()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/api/customer/:id", c.GetCustomer)
	r.POST("/api/login", c.UserLogin(*gin.Context, d.Db)) // ERROR c.UserLogin(*gin.Context, d.Db) (no value) used as valuecompilerTooManyValues

	fmt.Println("ENV: " + *&d.Cfg.EnvType)

	r.RunTLS(":"+d.Cfg.Port, d.Cfg.PemLoc, d.Cfg.KeyLoc)
}
