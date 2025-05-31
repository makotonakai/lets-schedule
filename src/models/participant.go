package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/MakotoNakai/lets-schedule/config"
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

func GetParticipantListByMeetingId(db *gorm.DB, Id int) ([]Participant, error) {

	participantList := []Participant{}
	err := db.Table("participants").
		Select("*").
		Where("participants.meeting_id = ?", Id).
		Find(&participantList).Error

	if err != nil {
		return participantList, config.ErrMeetingNotFound
	}
	
	return participantList, nil

}

func GetParticipantByUserIdAndMeetingId(db *gorm.DB, userId int, meetingId int) (Participant, error) {

	p := Participant{}
	err := db.Table("participants").
		Select("*").
		Where("participants.user_id = ?", userId).
		Where("participants.meeting_id = ?", meetingId).
		Find(&p).Error

	if err != nil {
		return p, config.ErrRecordNotFound
	}

	return p, nil

}

func ConvertToParticipant(db *gorm.DB, pw ParticipantWithUserName) (Participant, error) {
	
	userId, err := GetUserIdFromUserName(db, pw.UserName) 
	if err != nil {
		return Participant{}, config.ErrUserWithUserNameNotFound
	}

	p := Participant{}
	p.UserId = userId
	p.MeetingId = pw.MeetingId
	p.IsHost = pw.IsHost
	p.HasResponded = pw.HasResponded
	return p, nil

}

func ConvertToParticipantWithUserName(db *gorm.DB, p Participant) (ParticipantWithUserName, error) {
	
	userName, err := GetUserNameFromUserId(db, p.UserId)
	if err != nil {
		return ParticipantWithUserName{}, config.ErrUserWithUserIdNotFound
	}

	pw := ParticipantWithUserName{}
	pw.UserName = userName
	pw.MeetingId = p.MeetingId
	pw.IsHost = p.IsHost
	pw.HasResponded = p.HasResponded
	return pw, nil

}

func ConvertToParticipantWithUserNameList(db *gorm.DB, pl []Participant) ([]ParticipantWithUserName, error) {

	if len(pl) == 0 {
		return []ParticipantWithUserName{}, config.ErrParticipantListIsEmpty
	}

	pwl := []ParticipantWithUserName{}
	for _, p := range pl {
		pw, err := ConvertToParticipantWithUserName(db, p)
		if err != nil {
			return pwl, config.ErrUserWithUserIdNotFound
		}
		pwl = append(pwl, pw)
	}
	return pwl, nil

}


func ConvertToParticipantList(db *gorm.DB, pwl []ParticipantWithUserName) ([]Participant, error) {

	if len(pwl) == 0 {
		return []Participant{}, config.ErrParticipantWithUserNameListIsEmpty
	}

	pl := []Participant{}
	for _, pw := range pwl {
		p, err := ConvertToParticipant(db, pw)
		if err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return pl, nil

}

func Min(a, b int) int {

	if a <= b {
			return a
	}

	return b

}

func HostIsInParticipant(pl []Participant) bool {

	if len(pl) == 0 {
		return false
	}

	for _, p := range pl {
		if p.IsHost {
			return true
		}
	}

	return false

}



