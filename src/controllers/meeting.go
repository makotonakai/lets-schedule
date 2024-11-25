package controllers

import (

	"log"
	"strconv"
	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

//----------
// Handlers
//----------

// CreateMeeting creates a new meeting
// @Summary Create a new meeting
// @Description Create a meeting with title, type (onsite, online, or hybrid), and other details
// @Tags meetings
// @Accept json
// @Produce json
// @Param meeting body models.Meeting true "Meeting details"
// @Success 201 {object} models.Meeting
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/meetings/new [post]
func CreateMeeting(c echo.Context) error {
	
	newMeeting := models.Meeting{}
	err := c.Bind(&newMeeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrFailedToBindMeeting)
	}

	errorMessageListAboutMeeting := []string{}

	if models.IsTitleEmpty(newMeeting) {
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.TitleIsEmpty)
	}

	if models.IsOnsiteButNoPlaceSpecified(newMeeting) {
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.PlaceIsNotSpecified)
	}

	if models.IsOnlineButNoURLSpecified(newMeeting) {
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.URLIsNotSpecified)
		return c.JSON(http.StatusBadRequest, config.URLIsNotSpecified)
	}

	if models.IsHybridButNoPlaceSpecified(newMeeting) {
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.PlaceIsNotSpecified)
	}

	if models.IsHybridButNoURLSpecified(newMeeting) {
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.URLIsNotSpecified)
	}

	if models.IsHybridButNeitherPlaceOrURLSpecified(newMeeting) {
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.PlaceIsNotSpecified)
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.URLIsNotSpecified)
	}

	if models.ErrorsExist(errorMessageListAboutMeeting) {
		return c.JSON(http.StatusBadRequest, errorMessageListAboutMeeting)
	}

	err = db.Create(&newMeeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrFailedToCreateMeeting)
	}

	return c.JSON(http.StatusCreated, newMeeting)
	
}

// CreateMeeting creates a new meeting
// @Summary Create a new meeting
// @Description Create a meeting with title, type (onsite, online, or hybrid), and other details
// @Tags meetings
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} models.Meeting
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/meetings/user/{user_id} [get]
func GetMeetingsByUserId(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}
	
	err = db.First(&user, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrUserNotFound)
	} 

	meetings, err := models.GetMeetingsByUserId(db, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrRecordNotFound)
	}

	return c.JSON(http.StatusOK, meetings)

}


func GetMeetingById(c echo.Context) error {

	meeting := models.Meeting{}
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	err = db.First(&meeting, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrMeetingNotFound)
	}

	return c.JSON(http.StatusOK, meeting)

}

func GetConfirmedMeetingsForHost(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}
	
	err = db.First(&user, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrUserNotFound)
	}

	confirmedMeetingsForHost, err := models.GetConfirmedMeetingsForHostByUserId(db, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrMeetingNotFound)
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func GetNotConfirmedMeetingsForHost(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	err = db.First(&user, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrUserNotFound)
	}

	confirmedMeetingsForHost := models.GetNotConfirmedMeetingsForHostByUserId(db, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrMeetingNotFound)
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func GetNotRespondedMeetingsForHost(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	err = db.First(&user, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrUserNotFound)
	}

	confirmedMeetingsForHost := models.GetNotRespondedMeetingsForHostByUserId(db, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrMeetingNotFound)
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func GetConfirmedMeetingsForGuest(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	err = db.First(&user, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrUserNotFound)
	}

	confirmedMeetingsForHost := models.GetConfirmedMeetingsForGuestByUserId(db, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrMeetingNotFound)
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func GetNotConfirmedMeetingsForGuest(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	err := db.First(&user, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrUserNotFound)
	}

	confirmedMeetingsForHost := models.GetNotConfirmedMeetingsForGuestByUserId(db, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrMeetingNotFound)
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func GetNotRespondedMeetingsForGuest(c echo.Context) error {

	user := models.User{}
	id_str := c.Param("user_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	err = db.First(&user, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrUserNotFound)
	}

	confirmedMeetingsForHost := models.GetNotRespondedMeetingsForGuestByUserId(db, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, confirmedMeetingsForHost)

}

func UpdateMeetingById(c echo.Context) error {

	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	oldMeeting := models.Meeting{}
	err = db.First(&oldMeeting, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrMeetingNotFound)
	}

	newMeeting := models.Meeting{}
	err := c.Bind(&newMeeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrFailedToBindMeeting)
	}

	errorMessageListAboutMeeting := []string{}

	if models.IsTitleEmpty(newMeeting) {
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.TitleIsEmpty)
	}

	if models.IsHourEmpty(newMeeting) {
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.HourIsEmpty)
	}

	if models.IsOnsiteButNoPlaceSpecified(newMeeting) {
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.PlaceIsNotSpecified)
	}

	if models.IsOnlineButNoURLSpecified(newMeeting) {
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.URLIsNotSpecified)
	}

	if models.IsHybridButNeitherPlaceOrURLSpecified(newMeeting) {
		errorMessageListAboutMeeting = append(errorMessageListAboutMeeting, config.NeitherPlaceOrURLIsSpecified)
	}

	if models.ErrorsExist(errorMessageListAboutMeeting) {
		return c.JSON(http.StatusBadRequest, errorMessageListAboutMeeting)
	}
	
	err = db.Model(&oldMeeting).Updates(newMeeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToUpdateMeeting)
	}

	return c.JSON(http.StatusOK, newMeeting)

}

func DeleteMeeting(c echo.Context) error {

	Meeting := models.Meeting{}
	
	err := c.Bind(&Meeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrFailedToBindMeeting)
	}

	err = db.Delete(&Meeting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrFailedToDeleteMeeting)
	}
	
	return c.JSON(http.StatusNoContent, Meeting)

}
