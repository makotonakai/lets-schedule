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

func GetParticipantByUserIdAndMeetingId(userId int, meetingId int) Participant {
	p := Participant{}
	db.Table("participants").
		Select("participants.*").
		Where("participants.user_id = ?", userId).
		Where("participants.meeting_id = ?", meetingId).
		Find(&p)
	return p

}

func ConvertToParticipant(pw ParticipantWithUserName) Participant {
	p := Participant{}
	p.UserId = GetUserIdFromUserName(pw.UserName)
	p.MeetingId = pw.MeetingId
	p.IsHost = pw.IsHost
	p.HasResponded = pw.HasResponded
	return p
}

func ConvertToParticipantWithUserName(p Participant) ParticipantWithUserName {
	pw := ParticipantWithUserName{}
	pw.UserName = GetUserNameFromUserId(p.UserId)
	pw.MeetingId = p.MeetingId
	pw.IsHost = p.IsHost
	pw.HasResponded = p.HasResponded
	return pw
}

func ConvertToParticipantWithUserNameList(plist []Participant) []ParticipantWithUserName {
	pwl := []ParticipantWithUserName{}
	for _, p := range plist {
		pw := ConvertToParticipantWithUserName(p)
		pwl = append(pwl, pw)
	}
	return pwl
}

func ConvertToParticipantList(pwl []ParticipantWithUserName) []Participant {
	pl := []Participant{}
	for _, pw := range pwl {
		p := ConvertToParticipant(pw)
		pl = append(pl, p)
	}
	return pl
}



