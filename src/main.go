package main

import (
	
	"github.com/MakotoNakai/lets-schedule/router"
	
)

func main() {
	
	e := router.Initialize()
	e.Logger.Fatal(e.Start(":1323"))
	
}