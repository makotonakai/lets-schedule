package models

import (
	"time"
)

type CandidateTime struct {
	Id int `json:"id" param:"id"`
	MeetingId int `json:"meeting_id"`
	UserId int `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
