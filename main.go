package main

import (
	"ginws/config"
	"ginws/routes"
)

func main() {
	d, err := config.Init()
	if err != nil {
		panic(err)
	}

	r := routes.Routes(d)
	r.Run(":" + d.Cfg.Port)
}
