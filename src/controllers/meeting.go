package controllers

import (

	"time"
	"strconv"
	"net/http"
	
	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
)

func CreateMeeting(c echo.Context) error {
	
	newMeeting := models.Meeting{}
	err := c.Bind(&newMeeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newMeeting.CreatedAt = time.Now()
	newMeeting.UpdatedAt = time.Now()

	db.Create(&newMeeting)
	return c.JSON(http.StatusCreated, newMeeting)
	
}

func GetMeetingsByUserId(c echo.Context) error {

	userId, _ := strconv.Atoi(c.Param("user_id"))
	meetingList:= []models.Meeting{}
	db.Table("meetings").Select("meetings.id", "meetings.title", "meetings.description", "meetings.type", "meetings.place", "meetings.url").Joins("inner join user_meetings on meetings.id = user_meetings.meeting_id").Where("user_meetings.user_id = ?", userId).Find(&meetingList)
	
	return c.JSON(http.StatusOK, meetingList)

}

func GetConfirmedMeetingsForHost(c echo.Context) error {

	userId, _ := strconv.Atoi(c.Param("user_id"))
	meetingList:= []models.Meeting{}
	db.Table("meetings").Select("meetings.id", "meetings.title", "meetings.description", "meetings.type", "meetings.place", "meetings.url").Joins("inner join user_meetings on meetings.id = user_meetings.meeting_id").Joins("inner join participants on participants.user_meeting_id = user_meetings.id").Where("user_meetings.user_id = ?", userId).Where("participants.is_host = ?", 1).Where("meetings.is_confirmed = ?", 1).Find(&meetingList)
	return c.JSON(http.StatusOK, meetingList)

}

func GetNotYetConfirmedMeetingsForHost(c echo.Context) error {

	userId, _ := strconv.Atoi(c.Param("user_id"))
	meetingList:= []models.Meeting{}
	db.Table("meetings").Select("meetings.id", "meetings.title", "meetings.description", "meetings.type", "meetings.place", "meetings.url").Joins("inner join user_meetings on meetings.id = user_meetings.meeting_id").Joins("inner join participants on participants.user_meeting_id = user_meetings.id").Where("user_meetings.user_id = ?", userId).Where("participants.is_host = ?", 1).Where("meetings.is_confirmed = ?", 0).Find(&meetingList)
	return c.JSON(http.StatusOK, meetingList)

}

func GetConfirmedMeetingsForGuest(c echo.Context) error {

	userId, _ := strconv.Atoi(c.Param("user_id"))
	meetingList:= []models.Meeting{}
	// db.Table("meetings").Select("meetings.id", "meetings.title", "meetings.description", "meetings.type", "meetings.place", "meetings.url").Joins("inner join participants on meetings.id = participants.meeting_id").Where("participants.user_id = ?", userId).Where("participants.is_host = ?", 0).Where("meetings.is_confirmed = ?", 1).Find(&meetingList)
	db.Table("meetings").Select("meetings.id", "meetings.title", "meetings.description", "meetings.type", "meetings.place", "meetings.url").Joins("inner join user_meetings on meetings.id = user_meetings.meeting_id").Joins("inner join participants on participants.user_meeting_id = user_meetings.id").Where("user_meetings.user_id = ?", userId).Where("participants.is_host = ?", 0).Where("meetings.is_confirmed = ?", 1).Find(&meetingList)
	return c.JSON(http.StatusOK, meetingList)

}

func GetRespondedMeetingsForGuest(c echo.Context) error {

	userId, _ := strconv.Atoi(c.Param("user_id"))
	meetingList:= []models.Meeting{}
	// db.Table("meetings").Select("meetings.id", "meetings.title", "meetings.description", "meetings.type", "meetings.place", "meetings.url").Joins("inner join participants on meetings.id = participants.meeting_id").Where("participants.user_id = ?", userId).Where("participants.is_host = ?", 0).Where("meetings.is_confirmed = ?", 0).Where("participants.has_responded = ?", 1).Find(&meetingList)
	db.Table("meetings").Select("meetings.id", "meetings.title", "meetings.description", "meetings.type", "meetings.place", "meetings.url").Joins("inner join user_meetings on meetings.id = user_meetings.meeting_id").Joins("inner join participants on participants.user_meeting_id = user_meetings.id").Where("user_meetings.user_id = ?", userId).Where("participants.is_host = ?", 0).Where("meetings.is_confirmed = ?", 0).Where("participants.has_responded = ?", 1).Find(&meetingList)
	return c.JSON(http.StatusOK, meetingList)

}

func GetNotYetRespondedMeetingsForGuest(c echo.Context) error {

	userId, _ := strconv.Atoi(c.Param("user_id"))
	meetingList:= []models.Meeting{}
	// db.Table("meetings").Select("meetings.id", "meetings.title", "meetings.description", "meetings.type", "meetings.place", "meetings.url").Joins("inner join participants on meetings.id = participants.meeting_id").Where("participants.user_id = ?", userId).Where("participants.is_host = ?", 0).Where("meetings.is_confirmed = ?", 0).Where("participants.has_responded = ?", 0).Find(&meetingList)
	db.Table("meetings").Select("meetings.id", "meetings.title", "meetings.description", "meetings.type", "meetings.place", "meetings.url").Joins("inner join user_meetings on meetings.id = user_meetings.meeting_id").Joins("inner join participants on participants.user_meeting_id = user_meetings.id").Where("user_meetings.user_id = ?", userId).Where("participants.is_host = ?", 0).Where("meetings.is_confirmed = ?", 0).Where("participants.has_responded = ?", 0).Find(&meetingList)
	return c.JSON(http.StatusOK, meetingList)

}

func UpdateMeeting(c echo.Context) error {
	
	meetingId, _ := strconv.Atoi(c.Param("id"))
	meeting := models.Meeting{}
	err := c.Bind(&meeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db.Table("meetings").Where("id = ?", meetingId).Updates(&meeting)
	return c.JSON(http.StatusOK, meeting)
}

func DeleteMeeting(c echo.Context) error {

	meeting:= models.Meeting{}
	err := c.Bind(&meeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Delete(&meeting)

	return c.JSON(http.StatusNoContent, meeting)
}

