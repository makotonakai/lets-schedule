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

func GetCandidateTimeByMeetingIdAndUserId(MeetingId int, UserId int) []CandidateTime {

	CandidateTimeList := []CandidateTime{}
	db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Where("candidate_times.user_id = ?", UserId).
		Find(&CandidateTimeList)
	return CandidateTimeList

}

func GetAvailableTimeByMeetingId(MeetingId int) []CandidateTime {

	candidateTimeList := []CandidateTime{}
	db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Find(&candidateTimeList)

	availableTimeList := []CandidateTime{}

	if len(candidateTimeList) == 2 {

		time1 := candidateTimeList[0]
		time2 := candidateTimeList[1]

		availableTime := CandidateTime{}
		availableTime.StartTime = max(time1.StartTime, time2.StartTime)
		availableTime.EndTime = min(time1.EndTime, time2.EndTime)

		availableTimeList = append(availableTimeList, availableTime)
		return availableTimeList

	}else{

		return nil

	}
	
}

func min(Time1, Time2 time.Time) time.Time {
	if Time1.Equal(Time2) {
		return Time1
	}else if Time1.Before(Time2) {
		return Time1
	}else{
		return Time2
	}
}

func max(Time1, Time2 time.Time) time.Time {
	if Time1.Equal(Time2) {
		return Time1
	}else if Time1.Before(Time2) {
		return Time2
	}else{
		return Time1
	}
}
