package models_test

import (
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestGetParticipantListByMeetingIdSuccess(t *testing.T) {
	meetingId := 1
	participantList := models.GetParticipantListByMeetingId(mockDB, meetingId)
	expectedParticipantList := []models.Participant{
		models.Participant{
			Id: 1,
			UserId: 1,
			MeetingId: 1,
			IsHost: true,
			HasResponded: true,
			CreatedAt: participantList[0].CreatedAt,
			UpdatedAt: participantList[0].UpdatedAt,
		},
		models.Participant{
			Id: 2,
			UserId: 2,
			MeetingId: 1,
			IsHost: false,
			HasResponded: true,
			CreatedAt: participantList[1].CreatedAt,
			UpdatedAt: participantList[1].UpdatedAt,
		},
	}
	assert.Equal(t, participantList, expectedParticipantList)
}

func TestGetParticipantListByMeetingIdFail(t *testing.T) {
	meetingId := 100
	participantList := models.GetParticipantListByMeetingId(mockDB, meetingId)
	assert.Empty(t, participantList)
}

func TestGetParticipantByUserIdAndMeetingId(t *testing.T) {
	userId := 1
	meetingId := 1
	participant := models.GetParticipantByUserIdAndMeetingId(mockDB, userId, meetingId)

	expectedUserId := 1
	expectedMeetingId := 1
	assert.Equal(t, participant.UserId, expectedUserId)
	assert.Equal(t, participant.MeetingId, expectedMeetingId)
}

func TestMinBigger(t *testing.T) {
	result := models.Min(0, 100)
	expectedResult := 0
	assert.Equal(t, result, expectedResult)
}

func TestMinEqual(t *testing.T) {
	result := models.Min(0, 0)
	expectedResult := 0
	assert.Equal(t, result, expectedResult)
}

func TestHostIsInParticipantSuccess(t *testing.T) {
	participantList := []models.Participant{
		models.Participant{
			Id: 1,
			UserId: 1,
			MeetingId: 1,
			IsHost: true,
			HasResponded: true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	result := models.HostIsInParticipant(participantList)
	expectedResult := true
	assert.Equal(t, result, expectedResult)	
}

func TestHostIsInParticipantFail(t *testing.T) {
	participantList := []models.Participant{
		models.Participant{
			Id: 1,
			UserId: 1,
			MeetingId: 1,
			IsHost: false,
			HasResponded: true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	result := models.HostIsInParticipant(participantList)
	expectedResult := false
	assert.Equal(t, result, expectedResult)	
}
