package controllers

import (

	"log"
	"strconv"
	"net/http"
	
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

func GetAllMeetings(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&user, id)

	meetings := models.GetMeetingsByUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, meetings)

}

func GetMeetingById(c echo.Context) error {

	meeting := models.Meeting{}
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&meeting, id)

	return c.JSON(http.StatusOK, meeting)

}

func GetConfirmedMeetingsForHost(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&user, id)

	confirmedMeetingsForHost := models.GetConfirmedMeetingsForHostByUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func GetNotConfirmedMeetingsForHost(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&user, id)

	confirmedMeetingsForHost := models.GetNotConfirmedMeetingsForHostByUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func GetNotRespondedMeetingsForHost(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&user, id)

	confirmedMeetingsForHost := models.GetNotRespondedMeetingsForHostByUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func GetConfirmedMeetingsForGuest(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&user, id)

	confirmedMeetingsForHost := models.GetConfirmedMeetingsForGuestByUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func GetNotConfirmedMeetingsForGuest(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&user, id)

	confirmedMeetingsForHost := models.GetNotConfirmedMeetingsForGuestByUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func GetNotRespondedMeetingsForGuest(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&user, id)

	confirmedMeetingsForHost := models.GetNotRespondedMeetingsForGuestByUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func UpdateMeetingById(c echo.Context) error {

	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	oldMeeting := models.Meeting{}
	db.First(&oldMeeting, id)

	newMeeting := models.Meeting{}
	err = c.Bind(&newMeeting)
	
	db.Model(&oldMeeting).Updates(newMeeting)
	return c.JSON(http.StatusOK, oldMeeting)

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
