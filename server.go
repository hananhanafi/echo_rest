package main

import (
	"echo-rest/routes"

	"echo-rest/database"
)

func main() {
	database.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
