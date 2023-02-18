package models

import (
	"sort"
	"time"
	"reflect"
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
		
	sortByStartTime(candidateTimeList)

	userIdList := createUserIdList(candidateTimeList)
	userIdNum := len(userIdList)

	candidateTimeNum := len(candidateTimeList)
	maxIndex := candidateTimeNum - userIdNum + 1

	availableTimeList := []CandidateTime{}

	for index := 0; index < maxIndex; index++  {

		_candidateTimeList := candidateTimeList[index:index+userIdNum]
		_userIdList := createUserIdList(_candidateTimeList)

		if isSameSlice(userIdList, _userIdList) {

			startTime := getLatestStartTime(_candidateTimeList)
			endTime := getEarliestEndTime(_candidateTimeList)

			availableTime := CandidateTime{}
			availableTime.StartTime = startTime
			availableTime.EndTime = endTime
			availableTimeList = append(availableTimeList, availableTime)

		}
	}
	return availableTimeList
}

func include(numList []int, num int) bool {
	for _, val := range numList {
		if val == num {
			return true
		}
	}
	return false
}

func getLatestStartTime(candidateTimeList []CandidateTime) time.Time {
	latestStartTime := candidateTimeList[0].StartTime
	for i := 1; i < len(candidateTimeList); i++ {
		startTime := candidateTimeList[i].StartTime
		if startTime.After(latestStartTime) {
			latestStartTime = startTime
		}
	}
	return latestStartTime
}

func getEarliestEndTime(candidateTimeList []CandidateTime) time.Time {
	earliestEndTime := candidateTimeList[0].EndTime
	for i := 1; i < len(candidateTimeList); i++ {
		endTime := candidateTimeList[i].EndTime
		if endTime.Before(earliestEndTime) {
			earliestEndTime = endTime
		}
	}
	return earliestEndTime
}

func createUserIdList(candidateTimeList []CandidateTime) []int {
	userIdList := []int{}
	for _, candidateTime := range candidateTimeList {
		userId := candidateTime.UserId
		if !include(userIdList, userId) {
			userIdList = append(userIdList, userId)
		}
	}
	return userIdList
}

func isSameSlice(slice1, slice2 []int) bool {
	sort.Ints(slice1)
	sort.Ints(slice2)
	return reflect.DeepEqual(slice1, slice2)
}

func sortByStartTime(candidateTimeList []CandidateTime) {
	sort.Slice(candidateTimeList[:], func(i, j int) bool {
		return candidateTimeList[i].StartTime.Before(candidateTimeList[j].StartTime) 
	})
}

func overlapExist(candidateTimeList []CandidateTime) bool {
	sortByStartTime(candidateTimeList)
	return true
}

func AvailableTimeIsNotFound(candidateTimeList []CandidateTime) bool {
	return len(candidateTimeList) == 0
}
