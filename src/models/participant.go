package models

import (
	"time"
)

type Participant struct {	
	UserMeetingId int `json:"user_meeting_id"`
	IsHost bool `gorm:"not null" json:"is_host"`
	HasResponded bool `gorm:"not null" json:"has_reponded"`
	CreatedAt time.Time `gorm:"autoCreateTime:int" json:"created_at"`	
	UpdatedAt time.Time `gorm:"autoUpdateTime:int" json:"updated_at"`	
}
