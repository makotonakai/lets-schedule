package database

import (
	"time"

	"gorm.io/gorm"
	"github.com/MakotoNakai/lets-schedule/models"
)

func Seed(db *gorm.DB){
	SeedUser(db)
	SeedMeeting(db)
	SeedParticipant(db)
	SeedCandidateTime(db)
}

func SeedUser(db *gorm.DB) {
	user1 := models.User{
		Id: 1,
		UserName: "makoto",
		EmailAddress: "makoto@email.com",
		Password: "password",
		IsAdmin: true,
		CanLogin: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user2 := models.User{
		Id: 2,
		UserName: "minoru",
		EmailAddress: "minoru@email.com",
		Password: "password",
		IsAdmin: false,
		CanLogin: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.Create(&user1)
	db.Create(&user2)
}

func SeedMeeting(db *gorm.DB) {

	now := time.Now()

	meeting1 := models.Meeting{
		Id: 1,
		Title: "lab meeting",
		Description: "Just a regular weekly meeting",
		StartTime: now,
		EndTime: now.Add(1 * time.Hour),
		Type: "hybrid",
		MeetingPlace: "Delta 1F",
		MeetingUrl: "https://aqua-meeting.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	meeting2 := models.Meeting{
		Id: 2,
		Title: "meeting with Aram-san",
		Description: "progres report",
		StartTime: now,
		EndTime: now.Add(1 * time.Hour),
		Type: "online",
		MeetingPlace: "",
		MeetingUrl: "https://gaiax-meeting.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.Create(&meeting1)
	db.Create(&meeting2)
}

func SeedParticipant(db *gorm.DB) {
	participant1 := models.Participant{
		Id:1, 
		MeetingId: 1,
		UserId:1,
		IsHost:true,
		HasResponded:true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	participant2 := models.Participant{
		Id:2, 
		MeetingId: 1,
		UserId:2,
		IsHost:false,
		HasResponded:false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.Create(&participant1)
	db.Create(&participant2)
}

func SeedCandidateTime(db *gorm.DB) {
	candidateTime1 := models.CandidateTime{
		Id:1,
		MeetingId:1,
		UserId:1,
		StartTime:time.Now(),
		EndTime:time.Now(),
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
	}

	candidateTime2 := models.CandidateTime{
		Id:2,
		MeetingId:1,
		UserId:1,
		StartTime:time.Now(),
		EndTime:time.Now(),
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
	}

	db.Create(&candidateTime1)
	db.Create(&candidateTime2)
}


