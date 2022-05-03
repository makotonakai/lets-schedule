package models

import (
	"time"
)

type Participant struct {
	Id int `gorm:"primaryKey:autoIncrement" json:"id" param:"id"`	
	MeetingId int `json:"meeting_id" param:"meeting_id"`
	UserId int `json:"user_id"`
	IsHost bool `json:"is_host"`
	HasResponded bool `json:"has_reponded"`
	CreatedAt time.Time `gorm:"autoCreateTime:int" json:"created_at"`	
	UpdatedAt time.Time `gorm:"autoUpdateTime:int" json:"updated_at"`	
}