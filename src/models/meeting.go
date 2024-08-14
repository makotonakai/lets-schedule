package models

import (
	"time"
	"gorm.io/gorm"
)

type Meeting struct {
	Id int `gorm:"primaryKey:not null:autoIncrement" json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsOnsite bool `json:"is_onsite"`
	IsOnline bool `json:"is_online"`
	Place string `json:"place"`
	Url string `json:"url"`
	AllParticipantsResponded bool `json:"all_participants_responded"`
	IsConfirmed bool `json:"is_confirmed"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	ActualStartTime time.Time `json:"actual_start_time"`
	ActualEndTime time.Time `json:"actual_end_time"`
	Hour int `json:"hour"`
	CreatedAt time.Time `gorm:"autoCreateTime:int" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:int" json:"updated_at"`
}

type AvailableTime struct {
	ActualStartTime time.Time `json:"actual_start_time"`
	ActualEndTime   time.Time `json:"actual_end_time"`
}

func IsTitleEmpty(m Meeting) bool {
	return m.Title == ""
}

func IsHourEmpty(m Meeting) bool {
	return m.Hour == 0
}

func IsOnsiteButNoPlaceSpecified(m Meeting) bool {
	return m.IsOnsite == true && m.Place == "なし"
}

func IsOnlineButNoURLSpecified(m Meeting) bool {
	return m.IsOnline == true && m.Url == "なし"
}

func IsHybridButNeitherPlaceOrURLSpecified(m Meeting) bool {
	return m.IsOnsite == true && m.IsOnline == true && m.Place == "なし" && m.Url == "なし"
}

func GetMeetingById(db *gorm.DB, Id int) Meeting {
	meeting := Meeting{}
	db.Table("meetings").
		Select("meetings.*").
		Where("meetings.id = ?", Id).
		Find(&meeting)
	return meeting
}

func GetMeetingsByUserId(db *gorm.DB, UserId int) []Meeting {
	meetings := []Meeting{}
	db.Table("meetings").
		Select("meetings.*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("users.id = ?", UserId).
		Find(&meetings)
	return meetings
}

func GetConfirmedMeetingsForHostByUserId(db *gorm.DB, UserId int) []Meeting {
	meetings := []Meeting{}
	db.Table("meetings").
		Select("meetings.*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ?", UserId, 1).
		Where("meetings.is_confirmed = ?", 1).
		Find(&meetings)
	return meetings
}

func GetNotConfirmedMeetingsForHostByUserId(db *gorm.DB, UserId int) []Meeting {
	meetings := []Meeting{}
	db.Table("meetings").
		Select("meetings.*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ?", UserId, 1).
		Where("meetings.is_confirmed = ? AND meetings.all_participants_responded = ?", 0, 1).
		Find(&meetings)
	return meetings
}

func GetNotRespondedMeetingsForHostByUserId(db *gorm.DB, UserId int) []Meeting {
	meetings := []Meeting{}
	db.Table("meetings").
		Select("meetings.*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ?", UserId, 1).
		Where("meetings.is_confirmed = ? AND meetings.all_participants_responded = ?", 0, 0).
		Find(&meetings)
	return meetings
}

func GetConfirmedMeetingsForGuestByUserId(db *gorm.DB, UserId int) []Meeting {
	meetings := []Meeting{}
	db.Table("meetings").
		Select("meetings.*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ?", UserId, 0).
		Where("meetings.is_confirmed = ?", 1).
		Find(&meetings)
	return meetings
}

func GetNotConfirmedMeetingsForGuestByUserId(db *gorm.DB, UserId int) []Meeting {
	meetings := []Meeting{}
	db.Table("meetings").
		Select("meetings.*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ? AND participants.has_responded = ?", UserId, 0, 1).
		Where("meetings.is_confirmed = ?", 0).
		Find(&meetings)
	return meetings
}

func GetNotRespondedMeetingsForGuestByUserId(db *gorm.DB, UserId int) []Meeting {
	meetings := []Meeting{}
	db.Table("meetings").
		Select("meetings.*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ? AND participants.has_responded = ?", UserId, 0, 0).
		Where("meetings.is_confirmed = ?", 0).
		Find(&meetings)
	return meetings
}


