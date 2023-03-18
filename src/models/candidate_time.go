package models

import (
	"sort"
	"time"
	"reflect"
	"gorm.io/gorm"
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


func GetCandidateTimeByMeetingId(db *gorm.DB, MeetingId int) []CandidateTime {

	CandidateTimeList := []CandidateTime{}
	db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Find(&CandidateTimeList)
	return CandidateTimeList

}

func GetCandidateTimeByMeetingIdAndUserId(db *gorm.DB, MeetingId int, UserId int) []CandidateTime {

	CandidateTimeList := []CandidateTime{}
	db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Where("candidate_times.user_id = ?", UserId).
		Find(&CandidateTimeList)
	return CandidateTimeList

}

func GetAvailableTimeByMeetingId(db *gorm.DB, MeetingId int) []CandidateTime {

	candidateTimeList := []CandidateTime{}
	db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Find(&candidateTimeList)
		
	SortByStartTime(candidateTimeList)

	userIdList := CreateUserIdList(candidateTimeList)
	userIdNum := len(userIdList)

	candidateTimeNum := len(candidateTimeList)
	maxIndex := candidateTimeNum - userIdNum + 1

	availableTimeList := []CandidateTime{}

	for index := 0; index < maxIndex; index++  {

		_candidateTimeList := candidateTimeList[index:index+userIdNum]
		_userIdList := CreateUserIdList(_candidateTimeList)

		if IsSameSlice(userIdList, _userIdList) {

			startTime := GetLatestStartTime(_candidateTimeList)
			endTime := GetEarliestEndTime(_candidateTimeList)

			availableTime := CandidateTime{}
			availableTime.StartTime = startTime
			availableTime.EndTime = endTime
			availableTimeList = append(availableTimeList, availableTime)

		}
	}
	return availableTimeList
}

func Include(numList []int, num int) bool {
	for _, val := range numList {
		if val == num {
			return true
		}
	}
	return false
}

func GetLatestStartTime(candidateTimeList []CandidateTime) time.Time {
	latestStartTime := candidateTimeList[0].StartTime
	for i := 1; i < len(candidateTimeList); i++ {
		startTime := candidateTimeList[i].StartTime
		if startTime.After(latestStartTime) {
			latestStartTime = startTime
		}
	}
	return latestStartTime
}

func GetEarliestEndTime(candidateTimeList []CandidateTime) time.Time {
	earliestEndTime := candidateTimeList[0].EndTime
	for i := 1; i < len(candidateTimeList); i++ {
		endTime := candidateTimeList[i].EndTime
		if endTime.Before(earliestEndTime) {
			earliestEndTime = endTime
		}
	}
	return earliestEndTime
}

func CreateUserIdList(candidateTimeList []CandidateTime) []int {
	userIdList := []int{}
	for _, candidateTime := range candidateTimeList {
		userId := candidateTime.UserId
		if !Include(userIdList, userId) {
			userIdList = append(userIdList, userId)
		}
	}
	return userIdList
}

func IsSameSlice(slice1, slice2 []int) bool {
	sort.Ints(slice1)
	sort.Ints(slice2)
	return reflect.DeepEqual(slice1, slice2)
}

func SortByStartTime(candidateTimeList []CandidateTime) {
	sort.Slice(candidateTimeList[:], func(i, j int) bool {
		return candidateTimeList[i].StartTime.Before(candidateTimeList[j].StartTime) 
	})
}

func AvailableTimeIsNotFound(candidateTimeList []CandidateTime) bool {
	return len(candidateTimeList) == 0
}
