package controllers

import (

	"time"
	"strconv"
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
)

//----------
// Handlers
//----------

func CreateParticipant(c echo.Context) error {
	
	newParticipantWithUserNameList := []models.ParticipantWithUserName{}
	err := c.Bind(&newParticipantWithUserNameList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newParticipantList := []models.Participant{}

	for _, ParticipantWithUserName := range newParticipantWithUserNameList {

		Participant := models.Participant{}
		UserName := ParticipantWithUserName.UserName
		Participant.UserId = models.GetUserIdFromUserName(UserName)

		Participant.MeetingId = ParticipantWithUserName.MeetingId
		Participant.IsHost = ParticipantWithUserName.IsHost
		Participant.HasResponded = ParticipantWithUserName.HasResponded

		Participant.CreatedAt = time.Now()
		Participant.UpdatedAt = time.Now()

		newParticipantList = append(newParticipantList, Participant)

	}

	db.Create(&newParticipantList)
	return c.JSON(http.StatusCreated, newParticipantList)

}

func GetAllParticipant(c echo.Context) error {

	participantList := []models.Participant{}

	db.Find(&participantList)
	return c.JSON(http.StatusOK, participantList)

}

func GetParticipantByMeetingId(c echo.Context) error {

	pmi := c.Param("meeting_id")
	mi, err := strconv.Atoi(pmi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	pl := models.GetParticipantListByMeetingId(mi)
	pwl := models.ConvertToParticipantWithUserNameList(pl)

	return c.JSON(http.StatusOK, pwl)

}

func UpdateParticipantByUserIdAndMeetingId(c echo.Context) error {

	pmi := c.Param("meeting_id")
	mi, err := strconv.Atoi(pmi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	participantWithUserNameList := []models.ParticipantWithUserName{}
	err = c.Bind(&participantWithUserNameList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	oldParticipantList := models.GetParticipantListByMeetingId(mi)
	newParticipantList := []models.Participant{}
	for i, oldp := range oldParticipantList {
		pw := participantWithUserNameList[i]
		newp := models.ConvertToParticipant(pw)
		result := db.Model(&oldp).Updates(newp)
		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error)
		}
		newParticipantList = append(newParticipantList, oldp)
	}
	return c.JSON(http.StatusOK, newParticipantList)

}

func DeleteParticipant(c echo.Context) error {

	participant := models.Participant{}
	
	err := c.Bind(&participant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Delete(&participant)
	return c.JSON(http.StatusNoContent, participant)

}
