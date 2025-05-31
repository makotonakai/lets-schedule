package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/MakotoNakai/lets-schedule/config"
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

func IsTitleEmpty(m Meeting) bool {
	return m.Title == ""
}

func IsHourEmpty(m Meeting) bool {
	if m.Hour < 0 {
		return false
	}
	return m.Hour == 0
}

func IsOnsiteButNoPlaceSpecified(m Meeting) bool {
	return m.IsOnsite == true && m.IsOnline == false && m.Place == ""
}

func IsOnlineButNoURLSpecified(m Meeting) bool {
	return m.IsOnsite == false && m.IsOnline == true && m.Url == ""
}

func IsHybridButNeitherPlaceOrURLSpecified(m Meeting) bool {
	return m.IsOnsite == true && m.IsOnline == true && m.Place == "" && m.Url == ""
}

func IsHybridButNoPlaceSpecified(m Meeting) bool {
	return m.IsOnsite == true && m.IsOnline == true && m.Place == ""
}

func IsHybridButNoURLSpecified(m Meeting) bool {
	return m.IsOnsite == true && m.IsOnline == true && m.Url == ""
}

func GetMeetingById(db *gorm.DB, Id int) (Meeting, error) {
	meeting := Meeting{}
	err := db.Table("meetings").
		Select("*").
		Where("meetings.id = ?", Id).
		Find(&meeting).Error
	if err != nil {
		return meeting, config.ErrRecordNotFound
	}
	return meeting, nil
}

func GetMeetingsByUserId(db *gorm.DB, UserId int) ([]Meeting, error) {
	meetings := []Meeting{}
	err := db.Table("meetings").
		Select("*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("users.id = ?", UserId).
		Find(&meetings).Error
	if err != nil {
		return meetings, config.ErrRecordNotFound
	}
	return meetings, nil
}

func GetConfirmedMeetingsForHostByUserId(db *gorm.DB, UserId int) ([]Meeting, error) {
	meetings := []Meeting{}
	err := db.Table("meetings").
		Select("*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ?", UserId, 1).
		Where("meetings.is_confirmed = ?", 1).
		Find(&meetings).Error
	if err != nil {
		return meetings, config.ErrRecordNotFound
	}
	return meetings, nil
}

func GetNotConfirmedMeetingsForHostByUserId(db *gorm.DB, UserId int) ([]Meeting, error) {
	meetings := []Meeting{}
	err := db.Table("meetings").
		Select("*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ?", UserId, 1).
		Where("meetings.is_confirmed = ? AND meetings.all_participants_responded = ?", 0, 1).
		Find(&meetings).Error
	if err != nil {
		return meetings, config.ErrRecordNotFound
	}
	return meetings, nil
}

func GetNotRespondedMeetingsForHostByUserId(db *gorm.DB, UserId int) ([]Meeting, error) {
	meetings := []Meeting{}
	err := db.Table("meetings").
		Select("*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ?", UserId, 1).
		Where("meetings.is_confirmed = ? AND meetings.all_participants_responded = ?", 0, 0).
		Find(&meetings).Error
	if err != nil {
		return meetings, config.ErrRecordNotFound
	}
	return meetings, nil
}

func GetConfirmedMeetingsForGuestByUserId(db *gorm.DB, UserId int) ([]Meeting, error) {
	meetings := []Meeting{}
	err := db.Table("meetings").
		Select("*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ?", UserId, 0).
		Where("meetings.is_confirmed = ?", 1).
		Find(&meetings).Error
	if err != nil {
		return meetings, config.ErrRecordNotFound
	}
	return meetings, nil
}

func GetNotConfirmedMeetingsForGuestByUserId(db *gorm.DB, UserId int) ([]Meeting, error) {
	meetings := []Meeting{}
	err := db.Table("meetings").
		Select("*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ? AND participants.has_responded = ?", UserId, 0, 1).
		Where("meetings.is_confirmed = ?", 0).
		Find(&meetings).Error
	if err != nil {
		return meetings, config.ErrRecordNotFound
	}
	return meetings, nil
}

func GetNotRespondedMeetingsForGuestByUserId(db *gorm.DB, UserId int) ([]Meeting, error) {
	meetings := []Meeting{}
	err := db.Table("meetings").
		Select("*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("participants.user_id = ? AND participants.is_host = ? AND participants.has_responded = ?", UserId, 0, 0).
		Where("meetings.is_confirmed = ?", 0).
		Find(&meetings).Error
	if err != nil {
		return meetings, config.ErrRecordNotFound
	}
	return meetings, nil
}


