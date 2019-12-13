package main

import (
	"github.com/mblode/gospel/db"
	"github.com/mblode/gospel/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))
}
