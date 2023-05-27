package models_test

import (
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/MakotoNakai/lets-schedule/models"
)

var meeting = models.Meeting {
	Id: 1, 
	Title: "meeting",
	Description: "hoge",
	Type: "物理開催",
	Place: "tokyo",
	Url: "",
	AllParticipantsResponded: true,
	IsConfirmed: true,
	StartTime: time.Now(),
	EndTime: time.Now(),
	Hour: 1, 
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestIsTitleEmptySuccess(t *testing.T) {
	emptyMeeting := models.Meeting {
		Title: "",
	}
	result := models.IsTitleEmpty(emptyMeeting)
	expectedResult := true
	assert.Equal(t, result, expectedResult)
}

func TestIsTitleEmptyFail(t *testing.T) {
	result := models.IsTitleEmpty(meeting)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestIsOnsiteButNoPlaceSpecifiedSuccess(t *testing.T) {
	onsiteMeeting := models.Meeting{
		Id: 1, 
		Title: "meeting",
		Description: "hoge",
		Type: "物理開催",
		Place: "なし",
		Url: "",
		AllParticipantsResponded: true,
		IsConfirmed: true,
		StartTime: time.Now(),
		EndTime: time.Now(),
		Hour: 1, 
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := models.IsOnsiteButNoPlaceSpecified(onsiteMeeting)
	expectedResult := true
	assert.Equal(t, result, expectedResult)
}

func TestIsOnsiteButNoPlaceSpecifiedFail(t *testing.T) {
	onsiteMeeting := models.Meeting{
		Id: 1, 
		Title: "meeting",
		Description: "hoge",
		Type: "物理開催",
		Place: "東京",
		Url: "",
		AllParticipantsResponded: true,
		IsConfirmed: true,
		StartTime: time.Now(),
		EndTime: time.Now(),
		Hour: 1, 
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := models.IsOnsiteButNoPlaceSpecified(onsiteMeeting)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestIsOnlineButNoURLSpecifiedSuccess(t *testing.T) {
	onlineMeeting := models.Meeting{
		Id: 1, 
		Title: "meeting",
		Description: "hoge",
		Type: "オンライン開催",
		Place: "",
		Url: "なし",
		AllParticipantsResponded: true,
		IsConfirmed: true,
		StartTime: time.Now(),
		EndTime: time.Now(),
		Hour: 1, 
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := models.IsOnlineButNoURLSpecified(onlineMeeting)
	expectedResult := true
	assert.Equal(t, result, expectedResult)
}

func TestIsOnlineButNoURLSpecifiedFail(t *testing.T) {
	onlineMeeting := models.Meeting{
		Id: 1, 
		Title: "meeting",
		Description: "hoge",
		Type: "オンライン開催",
		Place: "",
		Url: "www.hoge.com",
		AllParticipantsResponded: true,
		IsConfirmed: true,
		StartTime: time.Now(),
		EndTime: time.Now(),
		Hour: 1, 
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := models.IsOnlineButNoURLSpecified(onlineMeeting)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestIsHybridButNeitherPlaceOrURLSpecified(t *testing.T) {
	hybridMeeting := models.Meeting{
		Id: 1, 
		Title: "meeting",
		Description: "hoge",
		Type: "ハイブリッド開催",
		Place: "なし",
		Url: "なし",
		AllParticipantsResponded: true,
		IsConfirmed: true,
		StartTime: time.Now(),
		EndTime: time.Now(),
		Hour: 1, 
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := models.IsHybridButNeitherPlaceOrURLSpecified(hybridMeeting)
	expectedResult := true
	assert.Equal(t, result, expectedResult)
}

func TestIsHybridAndOnlyPlaceIsSpecified(t *testing.T) {
	hybridMeeting := models.Meeting{
		Id: 1, 
		Title: "meeting",
		Description: "hoge",
		Type: "ハイブリッド開催",
		Place: "東京",
		Url: "なし",
		AllParticipantsResponded: true,
		IsConfirmed: true,
		StartTime: time.Now(),
		EndTime: time.Now(),
		Hour: 1, 
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := models.IsHybridButNeitherPlaceOrURLSpecified(hybridMeeting)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestIsHybridButOnlyUrlIsNotSpecified(t *testing.T) {
	hybridMeeting := models.Meeting{
		Id: 1, 
		Title: "meeting",
		Description: "hoge",
		Type: "ハイブリッド開催",
		Place: "なし",
		Url: "www.hoge.com",
		AllParticipantsResponded: true,
		IsConfirmed: true,
		StartTime: time.Now(),
		EndTime: time.Now(),
		Hour: 1, 
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := models.IsHybridButNeitherPlaceOrURLSpecified(hybridMeeting)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestIsHybridAndBothPlaceAndURLAreSpecified(t *testing.T) {
	hybridMeeting := models.Meeting{
		Id: 1, 
		Title: "meeting",
		Description: "hoge",
		Type: "ハイブリッド開催",
		Place: "東京",
		Url: "www.hoge.com",
		AllParticipantsResponded: true,
		IsConfirmed: true,
		StartTime: time.Now(),
		EndTime: time.Now(),
		Hour: 1, 
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := models.IsHybridButNeitherPlaceOrURLSpecified(hybridMeeting)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestGetMeetingByIdSuccess(t *testing.T) {
	id := 1
	meeting := models.GetMeetingById(mockDB, id)
	expectedId := 1
	assert.Equal(t, meeting.Id, expectedId)
}

func TestGetMeetingByIdFail(t *testing.T) {
	id := 10
	meeting := models.GetMeetingById(mockDB, id)
	expectedId := 10
	assert.NotEqual(t, meeting.Id, expectedId)
}

func TestGetMeetingsByUserIdSuccess(t *testing.T) {
	userId := 3
	meetings := models.GetMeetingsByUserId(mockDB, userId)
	
	expectedMeetingId := 2
	assert.Equal(t, meetings[0].Id, expectedMeetingId)
}

func TestGetMeetingsByUserIdFail(t *testing.T) {
	userId := 10
	meetings := models.GetMeetingsByUserId(mockDB, userId)
	assert.Empty(t, meetings)
}

func TestGetConfirmedMeetingsForHostByUserIdSuccess(t *testing.T) {
	userId := 1
	meetings := models.GetMeetingsByUserId(mockDB, userId)
	expectedMeetingId := 1
	assert.Equal(t, meetings[0].Id, expectedMeetingId)
}

func TestGetConfirmedMeetingsForHostByUserIdFail(t *testing.T) {
	userId := 10
	meetings := models.GetMeetingsByUserId(mockDB, userId)
	assert.Empty(t, meetings)
}

func TestGetNotConfirmedMeetingsForHostByUserIdSuccess(t *testing.T) {
	userId := 1
	meetings := models.GetNotConfirmedMeetingsForGuestByUserId(mockDB, userId)
	expectedMeetingId := 5
	assert.Equal(t, meetings[0].Id, expectedMeetingId)
}

func TestGetNotConfirmedMeetingsForHostByUserIdFail(t *testing.T) {
	userId := 10
	meetings := models.GetNotConfirmedMeetingsForGuestByUserId(mockDB, userId)
	assert.Empty(t, meetings)
}

func TestGetNotRespondedMeetingsForHostByUserIdSuccess(t *testing.T) {
	userId := 1
	meetings := models.GetNotRespondedMeetingsForHostByUserId(mockDB, userId)
	expectedMeetingId := 3
	assert.Equal(t, meetings[0].Id, expectedMeetingId)
}

func TestGetNotRespondedMeetingsForHostByUserIdFail(t *testing.T) {
	userId := 10
	meetings := models.GetNotConfirmedMeetingsForGuestByUserId(mockDB, userId)
	assert.Empty(t, meetings)
}

func TestGetConfirmedMeetingsForGuestByUserIdSuccess(t *testing.T) {
	userId := 2
	meetings := models.GetConfirmedMeetingsForGuestByUserId(mockDB, userId)
	expectedMeetingId := 1
	assert.Equal(t, meetings[0].Id, expectedMeetingId)
}

func TestGetConfirmedMeetingsForGuestByUserIdFail(t *testing.T) {
	userId := 10
	meetings := models.GetConfirmedMeetingsForGuestByUserId(mockDB, userId)
	assert.Empty(t, meetings)
}

func TestGetNotConfirmedMeetingsForGuestByUserIdSuccess(t *testing.T) {
	userId := 1
	meetings := models.GetNotConfirmedMeetingsForGuestByUserId(mockDB, userId)
	expectedMeetingId := 5
	assert.Equal(t, meetings[0].Id, expectedMeetingId)
}

func TestGetNotConfirmedMeetingsForGuestByUserIdFail(t *testing.T) {
	userId := 10
	meetings := models.GetNotConfirmedMeetingsForGuestByUserId(mockDB, userId)
	assert.Empty(t, meetings)
}

func TestGetNotRespondedMeetingsForGuestByUserIdSuccess(t *testing.T) {
	userId := 1
	meetings := models.GetNotRespondedMeetingsForGuestByUserId(mockDB, userId)
	expectedMeetingId := 6
	assert.Equal(t, meetings[0].Id, expectedMeetingId)
}

func TestGetNotRespondedMeetingsForGuestByUserIdFail(t *testing.T) {
	userId := 10
	meetings := models.GetNotRespondedMeetingsForGuestByUserId(mockDB, userId)
	assert.Empty(t, meetings)
}










