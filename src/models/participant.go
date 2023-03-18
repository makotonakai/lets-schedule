package models

import (
	"time"
	"gorm.io/gorm"
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

func GetParticipantListByMeetingId(db *gorm.DB, Id int) []Participant {
	participantList := []Participant{}
	db.Table("participants").
		Select("participants.*").
		Where("participants.meeting_id = ?", Id).
		Find(&participantList)
	return participantList
}

func GetParticipantByUserIdAndMeetingId(db *gorm.DB, userId int, meetingId int) Participant {
	p := Participant{}
	db.Table("participants").
		Select("participants.*").
		Where("participants.user_id = ?", userId).
		Where("participants.meeting_id = ?", meetingId).
		Find(&p)
	return p

}

func ConvertToParticipant(db *gorm.DB, pw ParticipantWithUserName) Participant {
	p := Participant{}
	p.UserId = GetUserIdFromUserName(db, pw.UserName)
	p.MeetingId = pw.MeetingId
	p.IsHost = pw.IsHost
	p.HasResponded = pw.HasResponded
	return p
}

func ConvertToParticipantWithUserName(db *gorm.DB, p Participant) ParticipantWithUserName {
	pw := ParticipantWithUserName{}
	pw.UserName = GetUserNameFromUserId(db, p.UserId)
	pw.MeetingId = p.MeetingId
	pw.IsHost = p.IsHost
	pw.HasResponded = p.HasResponded
	return pw
}

func ConvertToParticipantWithUserNameList(db *gorm.DB, plist []Participant) []ParticipantWithUserName {
	pwl := []ParticipantWithUserName{}
	for _, p := range plist {
		pw := ConvertToParticipantWithUserName(db, p)
		pwl = append(pwl, pw)
	}
	return pwl
}

func ConvertToParticipantList(db *gorm.DB, pwl []ParticipantWithUserName) []Participant {
	pl := []Participant{}
	for _, pw := range pwl {
		p := ConvertToParticipant(db, pw)
		pl = append(pl, p)
	}
	return pl
}

func Min(a, b int) int {
	if a <= b {
			return a
	}
	return b
}

func HostIsInParticipant(plist []Participant) bool {

	host := Participant{}
	for _, p := range plist {
		if p.IsHost == true {
			host = p
		}
		if host.Id == p.Id {
			return true
		}
	}

	return false

}



