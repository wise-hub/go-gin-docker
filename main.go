package main

import (
	"fmt"
	"ginws/config"
	"ginws/routes"

	_ "github.com/sijms/go-ora/v2"
)

func main() {

	d, err := config.Init()
	if err != nil {
		panic(err)
	}

	r := routes.Routes(d)

	fmt.Println("ENV: " + *&d.Cfg.EnvType)

	r.RunTLS(":"+d.Cfg.Port, d.Cfg.PemLoc, d.Cfg.KeyLoc)
}
