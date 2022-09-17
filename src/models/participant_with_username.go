package models

type ParticipantWithUserName struct {
	UserName string `json:"user_name"`
	MeetingId int `json:"meeting_id"`
	IsHost bool `json:"is_host"`
	HasResponded bool `json:"has_responded"`
}
