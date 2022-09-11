package controllers

import (

	"log"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
)

//----------
// Handlers
//----------

func CreateMeeting(c echo.Context) error {
	
	newMeeting := models.Meeting{}

	err := c.Bind(&newMeeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Create(&newMeeting)
	return c.JSON(http.StatusCreated, newMeeting)
	
}

func GetMeetingByUserId(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&user, id)

	meetings := []models.Meeting{}
	db.Table("meetings").
		Select("meetings.*").
		Joins("inner join participants on participants.meeting_id = meetings.id").
		Joins("inner join users on users.id = participants.user_id").
		Where("users.id = ?", id).
		Find(&meetings)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, meetings)

}

func UpdateMeeting(c echo.Context) error {

	Meeting := models.Meeting{}

	err := c.Bind(&Meeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Save(&Meeting)
	return c.JSON(http.StatusOK, Meeting)

}

func DeleteMeeting(c echo.Context) error {

	Meeting := models.Meeting{}
	
	err := c.Bind(&Meeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Delete(&Meeting)
	return c.JSON(http.StatusNoContent, Meeting)

}
