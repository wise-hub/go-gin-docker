package main

import (
	"fmt"
	"ginws/dep"

	"github.com/gin-gonic/gin"
	_ "github.com/sijms/go-ora/v2"
)

func router() *gin.Engine {

	r := gin.Default()

	r.GET("/api/customer/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.JSON(200, gin.H{"customer_endpoint_one": id})
	})

	return r
}

func main() {

	d, err := dep.Init()
	if err != nil {
		panic(err)
	}
	if d.Cfg.Environment == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	fmt.Println(d.Cfg.Database.DoConnect)

	r := router()
	r.RunTLS(":"+d.Cfg.Port, d.Cfg.PemLoc, d.Cfg.KeyLoc)
}
