package controllers

import (

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

	db.Create(&newMeeting)
	return c.JSON(http.StatusCreated, newMeeting)
	
}

func GetMeeting(c echo.Context) error {

	meeting := models.Meeting{}
	err := c.Bind(&meeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db.First(&meeting)

	return c.JSON(http.StatusOK, meeting)
}

func GetMeetings(c echo.Context) error {

	meetingList:= []models.Meeting{}
	db.Find(&meetingList)
	
	return c.JSON(http.StatusOK, meetingList)

}

func UpdateMeeting(c echo.Context) error {

	meeting := models.Meeting{}
	err := c.Bind(&meeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db.Save(&meeting)

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

