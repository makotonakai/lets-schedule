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
	
	newParticipantList := []models.Participant{}
	err := c.Bind(&newParticipantList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	for _, newParticipant := range newParticipantList {
		newParticipant.CreatedAt = time.Now()
		newParticipant.UpdatedAt = time.Now()
	}

	db.Create(&newParticipantList)
	return c.JSON(http.StatusCreated, newParticipantList)
	
}

func GetParticipant(c echo.Context) error {

	participant := models.Participant{}
	err := c.Bind(&participant)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.First(&participant)

	return c.JSON(http.StatusOK, participant)
}

func GetHostParticipant(c echo.Context) error {

	participantList := []models.Participant{}
	userName := c.Param("username")
	db.Where("user_name = ? ", userName).Where("is_host = ?", 1).Find(&participantList)
	
	return c.JSON(http.StatusOK, participantList)

}

func GetRespondedGuestParticipant(c echo.Context) error {

	participantList := []models.Participant{}
	userName := c.Param("username")
	db.Where("user_name = ? ", userName).Where("is_host = ?", 0).Where("is_responded = ?", 1).Find(&participantList)
	
	return c.JSON(http.StatusOK, participantList)

}

func GetNotRespondedGuestParticipant(c echo.Context) error {

	participantList := []models.Participant{}
	userName := c.Param("username")
	db.Where("user_name = ? ", userName).Where("is_host = ?", 0).Where("is_responded = ?", 0).Find(&participantList)
	
	return c.JSON(http.StatusOK, participantList)

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

