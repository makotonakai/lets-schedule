package models_test

import (
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestGetCandidateTimeByMeetingIdSuccess(t *testing.T) {

	meetingId := 1
	candidateTimeList := models.GetCandidateTimeByMeetingId(mockDB, meetingId)

	candidateTime1 := candidateTimeList[0]
	candidateTime2 := candidateTimeList[1]

	expectedMeetingId := 1
	expectedUserId1 := 1
	expectedStartTime1 := time.Time(time.Date(2022, time.September, 2, 10, 0, 0, 0, time.Local))
	expectedEndTime1 := time.Time(time.Date(2022, time.September, 2, 12, 0, 0, 0, time.Local))

	expectedUserId2 := 2
	expectedStartTime2 := time.Time(time.Date(2022, time.September, 2, 11, 0, 0, 0, time.Local))
	expectedEndTime2 := time.Time(time.Date(2022, time.September, 2, 13, 0, 0, 0, time.Local))

	assert.Equal(t, candidateTime1.MeetingId, expectedMeetingId)
	assert.Equal(t, candidateTime1.UserId, expectedUserId1)
	assert.Equal(t, candidateTime1.StartTime, expectedStartTime1)
	assert.Equal(t, candidateTime1.EndTime, expectedEndTime1)

	assert.Equal(t, candidateTime2.MeetingId, expectedMeetingId)
	assert.Equal(t, candidateTime2.UserId, expectedUserId2)
	assert.Equal(t, candidateTime2.StartTime, expectedStartTime2)
	assert.Equal(t, candidateTime2.EndTime, expectedEndTime2)

}

func TestGetCandidateTimeByMeetingIdFail(t *testing.T) {
	meetingId := 10
	candidateTimeList := models.GetCandidateTimeByMeetingId(mockDB, meetingId)
	assert.Empty(t, candidateTimeList)
}

func TestGetCandidateTimeByMeetingIdAndUserIdSuccess(t *testing.T) {

	meetingId := 1
	userId := 1

	candidateTimeList := models.GetCandidateTimeByMeetingIdAndUserId(mockDB, meetingId, userId)
	candidateTime := candidateTimeList[0]

	expectedMeetingId := 1
	expectedUserId := 1
	expectedStartTime := time.Time(time.Date(2022, time.September, 2, 10, 0, 0, 0, time.Local))
	expectedEndTime := time.Time(time.Date(2022, time.September, 2, 12, 0, 0, 0, time.Local))

	assert.Equal(t, candidateTime.MeetingId, expectedMeetingId)
	assert.Equal(t, candidateTime.UserId, expectedUserId)
	assert.Equal(t, candidateTime.StartTime, expectedStartTime)
	assert.Equal(t, candidateTime.EndTime, expectedEndTime)

}

func TestGetCandidateTimeByMeetingIdAndUserIdFail(t *testing.T) {
	meetingId := 1
	userId := 3
	candidateTimeList := models.GetCandidateTimeByMeetingIdAndUserId(mockDB, meetingId, userId)
	assert.Empty(t, candidateTimeList)
}

func TestIncludeSuccess(t *testing.T) {

	numList := []int{1, 2}
	num := 1

	result := models.Include(numList, num)
	expectedResult := true
	assert.Equal(t, result, expectedResult)
}

func TestIncludeFail(t *testing.T) {
	numList := []int{1, 2}
	num := 3

	result := models.Include(numList, num)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestGetLatestStartTime(t *testing.T) {

	meetingId := 1
	candidateTimeList := models.GetCandidateTimeByMeetingId(mockDB, meetingId)

	latestStartTime := models.GetLatestStartTime(candidateTimeList)
	expectedLatestStartTime := time.Time(time.Date(2022, time.September, 2, 11, 0, 0, 0, time.Local))
	assert.Equal(t, latestStartTime, expectedLatestStartTime)

}

func TestGetEarliestEndTime(t *testing.T) {

	meetingId := 1
	candidateTimeList := models.GetCandidateTimeByMeetingId(mockDB, meetingId)

	earliestEndTime := models.GetEarliestEndTime(candidateTimeList)
	expectedEarliestEndTime := time.Time(time.Date(2022, time.September, 2, 12, 0, 0, 0, time.Local))
	assert.Equal(t, earliestEndTime, expectedEarliestEndTime)

}

func TestCreateUserIdList(t *testing.T) {

	meetingId := 1
	candidateTimeList := models.GetCandidateTimeByMeetingId(mockDB, meetingId)

	userIdList := models.CreateUserIdList(candidateTimeList)
	expectedUserIdList := []int{1, 2}
	assert.Equal(t, userIdList, expectedUserIdList)

}

func TestIsSameSliceSuccess(t *testing.T) {

	slice1 := []int{1, 2}
	slice2 := []int{2, 1}
	result := models.IsSameSlice(slice1, slice2)

	expectedResult := true
	assert.Equal(t, result, expectedResult)

}

func TestIsSameSliceFail(t *testing.T) {

	slice1 := []int{1, 2}
	slice2 := []int{2, 3}
	result := models.IsSameSlice(slice1, slice2)

	expectedResult := false
	assert.Equal(t, result, expectedResult)

}

func TestSortByStartTime(t *testing.T) {

	candidateTimeList := []models.CandidateTime{
		models.CandidateTime{
			StartTime:  time.Time(time.Date(2022, time.September, 2, 11, 0, 0, 0, time.Local)),
			EndTime: time.Time(time.Date(2022, time.September, 2, 13, 0, 0, 0, time.Local)),
		},
		models.CandidateTime{
			StartTime:  time.Time(time.Date(2022, time.September, 2, 10, 0, 0, 0, time.Local)),
			EndTime: time.Time(time.Date(2022, time.September, 2, 12, 0, 0, 0, time.Local)),
		},
	}

	models.SortByStartTime(candidateTimeList)

	expectedCandidateTimeList := []models.CandidateTime{
		models.CandidateTime{
			StartTime:  time.Time(time.Date(2022, time.September, 2, 10, 0, 0, 0, time.Local)),
			EndTime: time.Time(time.Date(2022, time.September, 2, 12, 0, 0, 0, time.Local)),
		},
		models.CandidateTime{
			StartTime:  time.Time(time.Date(2022, time.September, 2, 11, 0, 0, 0, time.Local)),
			EndTime: time.Time(time.Date(2022, time.September, 2, 13, 0, 0, 0, time.Local)),
		},
	}

	assert.Equal(t, candidateTimeList, expectedCandidateTimeList)

}

func TestAvailableTimeIsNotFoundSuccess(t *testing.T) {
	candidateTimeList := []models.CandidateTime{}
	result := models.AvailableTimeIsNotFound(candidateTimeList)
	expectedResult := true
	assert.Equal(t, result, expectedResult)
}

func TestAvailableTimeIsNotFoundFail(t *testing.T) {
	candidateTimeList := []models.CandidateTime{
		models.CandidateTime{
			StartTime:  time.Time(time.Date(2022, time.September, 2, 10, 0, 0, 0, time.Local)),
			EndTime: time.Time(time.Date(2022, time.September, 2, 12, 0, 0, 0, time.Local)),
		},
		models.CandidateTime{
			StartTime:  time.Time(time.Date(2022, time.September, 2, 11, 0, 0, 0, time.Local)),
			EndTime: time.Time(time.Date(2022, time.September, 2, 13, 0, 0, 0, time.Local)),
		},
	}
	result := models.AvailableTimeIsNotFound(candidateTimeList)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}



