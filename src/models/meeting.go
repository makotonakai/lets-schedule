package models

import (
	"time"
)

type Meeting struct {
	Id int `json:"id" param:"id"`
	Title string `json:"string"`
	Description string `json:"description"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Type string `json:"type"`
	MeetingPlace string `json:"meeting_place"`
	MeetingUrl string `json:"meeting_url"`
}