package models

import (
	"time"
	"github.com/MakotoNakai/lets-schedule/database"
)

type Meeting struct {
	Id int `json:"id" param:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Type string `json:"type"`
	Place string `json:"place"`
	Url string `json:"url"`
	IsConfirmed bool `json:"is_confirmed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var db = database.Connect()

func GetMeetingsByUserId(UserId int) []Meeting {
	meetings := []Meeting{}
	db.Table("meetings").
		Select("meetings.*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("users.id = ?", UserId).
		Find(&meetings)
	return meetings
}

