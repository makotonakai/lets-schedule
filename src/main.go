package main

import (
	
	"github.com/MakotoNakai/lets-schedule/router"
	"github.com/MakotoNakai/lets-schedule/models"
	"github.com/MakotoNakai/lets-schedule/database"
)

func main() {

	db := database.Connect()
	db.AutoMigrate(&models.User{})
	database.Seed(db)
	
	e := router.Initialize()
	e.Logger.Fatal(e.Start(":3000"))
	
}