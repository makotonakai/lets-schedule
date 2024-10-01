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

func IsCandidateTimeEmpty(ctlist []CandidateTime) (bool, error) {
	if *ctlist == nil {
		return false, errors.New("The given array is nil")
	}
	return len(ctlist) == 0, nil
}

func IsAvailableTimeEmpty(at AvailableTime) (bool, error) {
	if *at == nil {
		return false, errors.New("The given AvailableTime object is nil")
	}
	return at.ActualStartTime == time.Time{} && at.ActualEndTime == time.Time{}, nil
}

func EmptyCandidateTimeExists(ctlist []CandidateTime) (bool, error) {
	if *ctlist == nil {
		return false, errors.New("The given array is nil")
	}
	t := time.Time{}
	for _, ct := range ctlist {
		if ct.StartTime.Equal(t) || ct.EndTime.Equal(t) {
			return true, nil
		}
	}
	return false, nil
}

func PastCandidateTimeExists(ctlist []CandidateTime) (bool, error) {
	if *ctlist == nil {
		return false, errors.New("The given array is nil")
	}

	now := time.Now()
	for _, ct := range ctlist {
		if ct.StartTime.Before(now) || ct.EndTime.Before(now) {
			return true, nil
		}
	}
	return false, nil
}


func GetCandidateTimeByMeetingId(db *gorm.DB, MeetingId int) ([]CandidateTime, error) {

	CandidateTimeList := []CandidateTime{}
	err := db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Find(&CandidateTimeList).Error
	if err != nil {
		return CandidateTimeList, err
	}
	return CandidateTimeList, nil

}

func GetCandidateTimeByMeetingIdAndUserId(db *gorm.DB, MeetingId int, UserId int) ([]CandidateTime, error) {

	CandidateTimeList := []CandidateTime{}
	err := db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Where("candidate_times.user_id = ?", UserId).
		Find(&CandidateTimeList).Error
	if err != nil {
		return CandidateTimeList, err
	}
	return CandidateTimeList, nil
}

// func RegisterAvailableTime(db *gorm.DB, )

func GetAvailableTimeByMeetingId(db *gorm.DB, MeetingId int) ([]CandidateTime, error) {

	candidateTimeList := []CandidateTime{}
	err := db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Find(&candidateTimeList).Error
	if err != nil {
		return candidateTimeList, err
	}
		
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
	return availableTimeList, nil
}

func Include(numList []int, num int) (bool, error) {

	if *numList == nil {
		return false, errors.New("The given int array is nil")
	}

	for _, val := range numList {
		if val == num {
			return true, nil
		}
	}
	return false, nil
}

func GetLatestStartTime(candidateTimeList []CandidateTime) (time.Time, error) {

	if *candidateTimeList == nil {
		return false, errors.New("The given list of candidateTime is nil")
	}

	latestStartTime := candidateTimeList[0].StartTime
	for i := 1; i < len(candidateTimeList); i++ {
		startTime := candidateTimeList[i].StartTime
		if startTime.After(latestStartTime) {
			latestStartTime = startTime
		}
	}
	return latestStartTime, nil
}

func GetEarliestEndTime(candidateTimeList []CandidateTime) (time.Time, error) {

	if *candidateTimeList == nil {
		return false, errors.New("The given list of candidateTime is nil")
	}

	earliestEndTime := candidateTimeList[0].EndTime
	for i := 1; i < len(candidateTimeList); i++ {
		endTime := candidateTimeList[i].EndTime
		if endTime.Before(earliestEndTime) {
			earliestEndTime = endTime
		}
	}
	return earliestEndTime, nil
}

func CreateUserIdList(candidateTimeList []CandidateTime) ([]int, error) {

	if *candidateTimeList == nil {
		return false, errors.New("The given list of candidateTime is nil")
	}

	userIdList := []int{}
	for _, candidateTime := range candidateTimeList {
		userId := candidateTime.UserId
		if !Include(userIdList, userId) {
			userIdList = append(userIdList, userId)
		}
	}
	return userIdList, nil
}

func IsSameSlice(slice1, slice2 []int) (bool, error) {

	if len(slice1) == 0 || len(slice2) == 0 {
		return false, errors.New("The given int array is empty")
	}
	sort.Ints(slice1)
	sort.Ints(slice2)
	return reflect.DeepEqual(slice1, slice2), nil
}

func SortByStartTime(candidateTimeList []CandidateTime) {
	sort.Slice(candidateTimeList[:], func(i, j int) bool {
		return candidateTimeList[i].StartTime.Before(candidateTimeList[j].StartTime) 
	})
}

func AvailableTimeIsNotFound(candidateTimeList []CandidateTime) (bool, error) {

	if *candidateTimeList == nil {
		return false, errors.New("The given list of candidateTime is nil")
	}
	
	return len(candidateTimeList) == 0
}
