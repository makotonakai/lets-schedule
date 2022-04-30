package models

import (
	"time"
)

type Participant struct {
	Id int `json:"id"`
	MeetingId int `json:"meeting_id" param:"meeting_id"`
	UserId int `json:"user_id"`
	IsHost bool `json:"is_host"`
	HasResponded bool `json:"has_reponded"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}