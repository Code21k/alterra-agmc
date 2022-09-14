package main

import (
	"c21k.alterra.agmc/config"
	"c21k.alterra.agmc/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	e.Logger.Fatal(e.Start(":4848"))
}
