package controllers

import (

	"strconv"
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
)

//----------
// Handlers
//----------

func CreateCandidateTime(c echo.Context) error {
	
	newCandidateTime := []models.CandidateTime{}
	err := c.Bind(&newCandidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Create(&newCandidateTime)
	return c.JSON(http.StatusCreated, newCandidateTime)
	
}

func GetCandidateTimeWithUserNameByMeetingId(c echo.Context) error {

	MeetingIdString := c.Param("meeting_id")
	MeetingId, err := strconv.Atoi(MeetingIdString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newCandidateTimeList := []models.CandidateTime{}
	db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Find(&newCandidateTimeList)

	CandidateTimeWithUserNameList := []models.CandidateTimeWithUserName{}

		for _, CandidateTime := range newCandidateTimeList {

			CandidateTimeWithUserName := models.CandidateTimeWithUserName{}
			UserId := CandidateTime.UserId

			CandidateTimeWithUserName.UserName = models.GetUserNameFromUserId(UserId)
			CandidateTimeWithUserName.MeetingId = CandidateTime.MeetingId
			CandidateTimeWithUserName.StartTime = CandidateTime.StartTime
			CandidateTimeWithUserName.EndTime = CandidateTime.EndTime

			CandidateTimeWithUserNameList = append(CandidateTimeWithUserNameList, CandidateTimeWithUserName)
	
		}

	return c.JSON(http.StatusOK, CandidateTimeWithUserNameList)
}

func GetCandidateTimeByUserIdAndMeetingId(c echo.Context) error {

	UserIdString := c.Param("user_id")
	UserId, err := strconv.Atoi(UserIdString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	MeetingIdString := c.Param("meeting_id")
	MeetingId, err := strconv.Atoi(MeetingIdString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	CandidateTimeList := []models.CandidateTime{}
	CandidateTimeList = models.GetCandidateTimeByMeetingIdAndUserId(MeetingId, UserId)

	return c.JSON(http.StatusOK, CandidateTimeList)

}

func UpdateCandidateTimeByUserIdAndMeetingId(c echo.Context) error {

	UserIdString := c.Param("user_id")
	UserId, err := strconv.Atoi(UserIdString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	MeetingIdString := c.Param("meeting_id")
	MeetingId, err := strconv.Atoi(MeetingIdString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	OldCandidateTimeList := []models.CandidateTime{}
	OldCandidateTimeList = models.GetCandidateTimeByMeetingIdAndUserId(MeetingId, UserId)
	db.Delete(OldCandidateTimeList)

	NewCandidateTimeList := []models.CandidateTime{}
	err = c.Bind(&NewCandidateTimeList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Create(&NewCandidateTimeList)
	return c.JSON(http.StatusOK, NewCandidateTimeList)

}

func DeleteCandidateTime(c echo.Context) error {

	CandidateTime := models.CandidateTime{}
	
	err := c.Bind(&CandidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Delete(&CandidateTime)
	return c.JSON(http.StatusNoContent, CandidateTime)

}

