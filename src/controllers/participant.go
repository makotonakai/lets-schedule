package controllers

import (

	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
)

//----------
// Handlers
//----------

func CreateParticipant(c echo.Context) error {
	
	newParticipant := models.Participant{}

	err := c.Bind(&newParticipant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Create(&newParticipant)
	return c.JSON(http.StatusCreated, newParticipant)
	
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

func GetParticipants(c echo.Context) error {

	participantList := []models.Participant{}

	db.Find(&participantList)
	
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

