package main

import (
	"github.com/MakotoNakai/lets-schedule/router"
)

func main() {

	e := router.Router()
	e.Logger.Fatal(e.Start(":3000"))
	
}