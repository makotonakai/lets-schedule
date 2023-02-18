package models

import (
	"time"
)

type Meeting struct {
	Id int `gorm:"primaryKey:not null:autoIncrement" json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Type string `json:"type"`
	Place string `json:"place"`
	Url string `json:"url"`
	AllParticipantsResponded bool `json:"all_participants_responded"`
	IsConfirmed bool `json:"is_confirmed"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Hour int `json:"hour"`
	CreatedAt time.Time `gorm:"autoCreateTime:int" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:int" json:"updated_at"`
}

func IsTitleBlank(m Meeting) bool {
	return m.Title == ""
}

func IsOnsiteButNoPlaceSpecified(m Meeting) bool {
	return m.Type == "物理開催" && m.Place == "なし"
}

func IsOnlineButNoURLSpecified(m Meeting) bool {
	return m.Type == "オンライン開催" && m.Url == "なし"
}

func IsHybridButNeitherPlaceOrURLSpecified(m Meeting) bool {
	return m.Type == "ハイブリッド開催" && m.Place == "なし" && m.Url == "なし"
}

func GetMeetingById(Id int) Meeting {
	meeting := Meeting{}
	db.Table("meetings").
		Select("meetings.*").
		Where("meetings.id = ?", Id).
		Find(&meeting)
	return meeting
}

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

func GetConfirmedMeetingsForHostByUserId(UserId int) []Meeting {
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

func GetNotConfirmedMeetingsForHostByUserId(UserId int) []Meeting {
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

func GetNotRespondedMeetingsForHostByUserId(UserId int) []Meeting {
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

func GetConfirmedMeetingsForGuestByUserId(UserId int) []Meeting {
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

func GetNotConfirmedMeetingsForGuestByUserId(UserId int) []Meeting {
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

func GetNotRespondedMeetingsForGuestByUserId(UserId int) []Meeting {
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


