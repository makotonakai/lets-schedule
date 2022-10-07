package models

import (
	"time"
)

type Participant struct {
	Id int `gorm:"primaryKey:not null:autoIncrement" json:"id"`	
	UserId int `json:"user_id"`
	MeetingId int `json:"meeting_id"`
	IsHost bool `json:"is_host"`
	HasResponded bool `json:"has_responded"`
	CreatedAt time.Time `gorm:"autoCreateTime:int" json:"created_at"`	
	UpdatedAt time.Time `gorm:"autoUpdateTime:int" json:"updated_at"`	
}

func GetParticipantListByMeetingId(Id int) []Participant {
	participantList := []Participant{}
	db.Table("participants").
		Select("participants.*").
		Where("participants.meeting_id = ?", Id).
		Find(&participantList)
	return participantList
}

func GetParticipantWithUserName(p Participant) ParticipantWithUserName {
	pw := ParticipantWithUserName{}
	pw.UserName = GetUserNameFromUserId(p.UserId)
	pw.MeetingId = p.MeetingId
	pw.IsHost = p.IsHost
	pw.HasResponded = p.HasResponded
	return pw
}

func GetParticipantWithUserNameList(plist []Participant) []ParticipantWithUserName {
	pwl := []ParticipantWithUserName{}
	for _, p := range plist {
		pw := GetParticipantWithUserName(p)
		pwl = append(pwl, pw)
	}
	return pwl
}



