package models

import (
	"time"
)

type CandidateTime struct {
	Id int `json:"id" param:"id"`
	MeetingId int `json:"meeting_id"`
	UserId int `json:"user_id"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetCandidateTimeByMeetingId(MeetingId int) []CandidateTime {

	CandidateTimeList := []CandidateTime{}
	db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Find(&CandidateTimeList)
	return CandidateTimeList

}
