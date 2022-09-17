package controllers

import (

	"time"
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

func GetParticipantById(c echo.Context) error {

	participant := models.Participant{}
	err := c.Bind(&participant)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.First(&participant)
	return c.JSON(http.StatusOK, participant)

}

func UpdateParticipant(c echo.Context) error {

	participant := models.Participant{}

	err := c.Bind(&participant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Save(&participant)
	return c.JSON(http.StatusOK, participant)

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
