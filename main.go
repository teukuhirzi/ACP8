package main

import (
	"acp8/configs"
	"acp8/routes"
)

func main() {
	configs.InitDB()
	configs.InitMigrate()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
