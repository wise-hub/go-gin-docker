package main

import (
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
	r.Run(":8000")
}
