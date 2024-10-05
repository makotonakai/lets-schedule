package models

import (
	"time"
	"errors"
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

func GetParticipantListByMeetingId(db *gorm.DB, Id int) ([]Participant, error) {
	participantList := []Participant{}
	err := db.Table("participants").
		Select("participants.*").
		Where("participants.meeting_id = ?", Id).
		Find(&participantList).Error
	if err != nil {
		return participantList, err
	}
	return participantList, nil
}

func GetParticipantByUserIdAndMeetingId(db *gorm.DB, userId int, meetingId int) (Participant, error) {
	p := Participant{}
	err := db.Table("participants").
		Select("participants.*").
		Where("participants.user_id = ?", userId).
		Where("participants.meeting_id = ?", meetingId).
		Find(&p).Error
	if err != nil {
		return p, err
	}
	return p, nil

}

func ConvertToParticipant(db *gorm.DB, pw ParticipantWithUserName) (*Participant, error) {
	if &pw == nil {
		return &Participant{}, errors.New("The given ParticipantWithUserName object is nil")
	}
	p := &Participant{}
	userId, err := GetUserIdFromUserName(db, pw.UserName) 
	if err != nil {
		return nil, err
	}
	p.UserId = userId
	p.MeetingId = pw.MeetingId
	p.IsHost = pw.IsHost
	p.HasResponded = pw.HasResponded
	return p, nil
}

func ConvertToParticipantWithUserName(db *gorm.DB, p Participant) (ParticipantWithUserName, error) {
	if &p == nil {
		return ParticipantWithUserName{}, errors.New("The given Participant object is nil")
	}
	pw := ParticipantWithUserName{}
	pw.UserName, _ = GetUserNameFromUserId(db, p.UserId)
	pw.MeetingId = p.MeetingId
	pw.IsHost = p.IsHost
	pw.HasResponded = p.HasResponded
	return pw, nil
}

func ConvertToParticipantWithUserNameList(db *gorm.DB, plist []Participant) ([]ParticipantWithUserName, error) {
	if len(plist) == 0 {
		return []ParticipantWithUserName{}, errors.New("The given ParticipantWithUserName list is empty")
	}
	pwl := []ParticipantWithUserName{}
	for _, p := range plist {
		pw, err := ConvertToParticipantWithUserName(db, p)
		if err != nil {
			return pwl, err
		}
		pwl = append(pwl, pw)
	}
	return pwl, nil
}

func ConvertToParticipantList(db *gorm.DB, pwl []ParticipantWithUserName) (*[]Participant, error) {
	pl := &[]Participant{}
	for _, pw := range pwl {
		p, err := ConvertToParticipant(db, pw)
		if err != nil {
			return nil, err
		}
		*pl = append(*pl, *p)
	}
	return pl, nil
}

func Min(a, b int) (int, error) {
	if &a == nil || &b == nil {
		return -1, errors.New("The given integer is nil")
	}
	if a <= b {
			return a, nil
	}
	return b, nil
}

func HostIsInParticipant(plist []Participant) (bool, error) {
	if len(plist) == 0 {
		return false, errors.New("The given Participant list is empty")
	}
	host := Participant{}
	for _, p := range plist {
		if p.IsHost == true {
			host = p
		}
		if host.Id == p.Id {
			return true, nil
		}
	}

	return false, nil

}



