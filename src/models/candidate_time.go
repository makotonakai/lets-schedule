package models

import (
	"sort"
	"time"
	"errors"
	"reflect"
	"gorm.io/gorm"
	"github.com/MakotoNakai/lets-schedule/config"
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

type AvailableTime struct {
	ActualStartTime time.Time `json:"actual_start_time"`
	ActualEndTime   time.Time `json:"actual_end_time"`
}

func IsCandidateTimeEmpty(ctlist *[]CandidateTime) (bool, error) {
	if ctlist == nil {
		return false, config.ErrArrayIsNil
	}

	return len(*ctlist) == 0, nil
}

func IsAvailableTimeEmpty(at *AvailableTime) (bool, error) {
	if at == nil {
		return false, config.ErrAvailableTimeIsNil
	}
	return at.ActualStartTime == time.Time{} && at.ActualEndTime == time.Time{}, nil
}

func EmptyCandidateTimeExists(ctlist *[]CandidateTime) (bool, error) {
	if ctlist == nil {
		return false, config.ErrArrayIsNil
	}

	if len(*ctlist) == 0 {
		return false, config.ErrArrayIsEmpty
	}

	t := time.Time{}
	for _, ct := range *ctlist {
		if ct.StartTime.Equal(t) || ct.EndTime.Equal(t) {
			return true, nil
		}
	}
	return false, nil
}

func PastCandidateTimeExists(ctlist *[]CandidateTime) (bool, error) {
	if ctlist == nil {
		return false, config.ErrArrayIsNil
	}

	if len(*ctlist) == 0 {
		return false, config.ErrArrayIsEmpty
	}

	now := time.Now()
	for _, ct := range *ctlist {
		if ct.StartTime.Before(now) || ct.EndTime.Before(now) {
			return true, nil
		}
	}
	return false, nil
}


func GetCandidateTimeByMeetingId(db *gorm.DB, MeetingId int) ([]CandidateTime, error) {

	CandidateTimeList := []CandidateTime{}
	err := db.Table("candidate_times").
		Select("*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Find(&CandidateTimeList).Error
	if err != nil {
		return CandidateTimeList, config.ErrRecordNotFound
	}
	return CandidateTimeList, nil

}

func GetCandidateTimeByMeetingIdAndUserId(db *gorm.DB, MeetingId int, UserId int) ([]CandidateTime, error) {

	CandidateTimeList := []CandidateTime{}
	err := db.Table("candidate_times").
		Select("*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Where("candidate_times.user_id = ?", UserId).
		Find(&CandidateTimeList).Error
	if err != nil {
		return CandidateTimeList, config.ErrRecordNotFound
	}
	return CandidateTimeList, nil
}

func GetAvailableTimeByMeetingId(db *gorm.DB, MeetingId int) ([]CandidateTime, error) {

	candidateTimeList := []CandidateTime{}
	err := db.Table("candidate_times").
		Select("*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Find(&candidateTimeList).Error
	if err != nil {
		return candidateTimeList, config.ErrRecordNotFound
	}
		
	SortByStartTime(candidateTimeList)

	userIdList, err := CreateUserIdList(&candidateTimeList)
	if err != nil {
		return candidateTimeList, err
	}

	userIdNum := len(userIdList)

	candidateTimeNum := len(candidateTimeList)
	maxIndex := candidateTimeNum - userIdNum + 1

	availableTimeList := []CandidateTime{}

	for index := 0; index < maxIndex; index++  {

		_candidateTimeList := candidateTimeList[index:index+userIdNum]
		_userIdList, err := CreateUserIdList(&_candidateTimeList)
		if err != nil {
			return _candidateTimeList, err
		}

		sameSlice, err := IsSameSlice(&userIdList, &_userIdList)
		if err != nil {
			return _candidateTimeList, err
		} 

		if sameSlice {

			startTime, err := GetLatestStartTime(&_candidateTimeList)
			if err != nil {
				return _candidateTimeList, err
			}

			endTime, err := GetEarliestEndTime(&_candidateTimeList)
			if err != nil {
				return _candidateTimeList, err
			}

			availableTime := CandidateTime{}
			availableTime.StartTime = startTime
			availableTime.EndTime = endTime
			availableTimeList = append(availableTimeList, availableTime)

		}
	}
	return availableTimeList, nil
}

func Include(numList *[]int, num *int) (bool, error) {

	if numList == nil {
		return false, errors.New("The given int array is nil")
	}

	if num == nil {
		return false, errors.New("The given int is nil")
	}

	if len(*numList) == 0 {
		return false, errors.New("The given int array is empty")
	}

	for _, val := range *numList {
		if val == *num {
			return true, nil
		}
	}
	return false, nil
}

func GetLatestStartTime(candidateTimeList *[]CandidateTime) (time.Time, error) {

	if candidateTimeList == nil {
		return time.Time{}, errors.New("The given list of candidateTime is nil")
	}

	if len(*candidateTimeList) == 0 {
		return time.Time{}, errors.New("The given list of candidateTime is empty")
	}

	cl := *candidateTimeList
	latestStartTime := cl[0].StartTime
	for i := 1; i < len(cl); i++ {
		startTime := cl[i].StartTime
		if startTime.After(latestStartTime) {
			latestStartTime = startTime
		}
	}
	return latestStartTime, nil
}

func GetEarliestEndTime(candidateTimeList *[]CandidateTime) (time.Time, error) {

	if candidateTimeList == nil {
		return time.Time{}, errors.New("The given list of candidateTime is nil")
	}

	if len(*candidateTimeList) == 0 {
		return time.Time{}, errors.New("The given list of candidateTime is empty")
	}

	cl := *candidateTimeList
	earliestEndTime := cl[0].EndTime
	for i := 1; i < len(cl); i++ {
		endTime := cl[i].EndTime
		if endTime.Before(earliestEndTime) {
			earliestEndTime = endTime
		}
	}
	return earliestEndTime, nil
}

func CreateUserIdList(candidateTimeList *[]CandidateTime) ([]int, error) {

	if candidateTimeList == nil {
		return []int{}, errors.New("The given list of candidateTime is nil")
	}

	if len(*candidateTimeList) == 0 {
		return []int{}, errors.New("The given list of candidateTime is empty")
	}

	userIdList := []int{}
	for _, candidateTime := range *candidateTimeList {
		userId := candidateTime.UserId
		include, err := Include(&userIdList, &userId)
		if err != nil {
			return userIdList, err
		}

		if !include {
			userIdList = append(userIdList, userId)
		}
	}
	return userIdList, nil
}

func IsSameSlice(slice1, slice2 *[]int) (bool, error) {

	if slice1 == nil || slice2 == nil {
		return false, errors.New("The given int array is nil")
	}
	
	if len(*slice1) == 0 || len(*slice2) == 0 {
		return false, errors.New("The given int array is empty")
	}

	sort.Ints(*slice1)
	sort.Ints(*slice2)
	return reflect.DeepEqual(*slice1, *slice2), nil
}

func SortByStartTime(candidateTimeList []CandidateTime) {
	sort.Slice(candidateTimeList[:], func(i, j int) bool {
		return candidateTimeList[i].StartTime.Before(candidateTimeList[j].StartTime) 
	})
}

func AvailableTimeIsNotFound(candidateTimeList *[]CandidateTime) (bool, error) {

	if candidateTimeList == nil {
		return false, errors.New("The given list of candidateTime is nil")
	}
	
	return len(*candidateTimeList) == 0, nil
}
