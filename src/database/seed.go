package database

import (
	"time"

	"gorm.io/gorm"
	"github.com/MakotoNakai/lets-schedule/models"
)

func Seed(db *gorm.DB){
	SeedUser(db)
	SeedMeeting(db)
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

func SeedMeeting(db *gorm.DB) {

	now := time.Now()
	meeting1 := models.Meeting{
		Title: "lab meeting",
		Description: "Just a regular weekly meeting",
		StartTime: now,
		EndTime: now.Add(1 * time.Hour),
		Type: "hybrid",
		MeetingPlace: "Delta 1F",
		MeetingUrl: "https://aqua-meeting.com",
	}

	meeting2 := models.Meeting{
		Title: "meeting with Aram-san",
		Description: "progres report",
		StartTime: now,
		EndTime: now.Add(1 * time.Hour),
		Type: "online",
		MeetingPlace: "",
		MeetingUrl: "https://gaiax-meeting.com",
	}

	db.Create(&meeting1)
	db.Create(&meeting2)
}
