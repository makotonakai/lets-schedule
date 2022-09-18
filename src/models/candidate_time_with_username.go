package models

import (
	"time"
)

type CandidateTimeWithUserName struct {
	MeetingId int `json:"meeting_id"`
	UserName string `json:"user_name"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
}
