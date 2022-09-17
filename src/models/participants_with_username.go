package models

type ParticipantWithUserName struct {
	Id int `json:"id"`	
	MeetingId int `json:"meeting_id"`
	UserName int `json:"user_name"`
	IsHost bool `json:"is_host"`
	HasResponded bool `json:"has_reponded"`
}
