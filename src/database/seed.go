package database

import (
	"gorm.io/gorm"
	"github.com/MakotoNakai/lets-schedule/models"
)

func Seed(db *gorm.DB){
	SeedUser(db)
}

func SeedUser(db *gorm.DB) {
	user1 := models.User{
		Id: 1,
		UserName: "makoto",
		EmailAddress: "makoto@email.com",
		Password: "password",
		IsAdmin: true,
		CanLogin: true,
	}

	user2 := models.User{
		Id: 2,
		UserName: "minoru",
		EmailAddress: "minoru@email.com",
		Password: "password",
		IsAdmin: false,
		CanLogin: false,
	}

	db.Create(&user1)
	db.Create(&user2)
}