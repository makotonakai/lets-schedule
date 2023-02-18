package controllers

import (
	"strconv"
	"net/http"
	"gorm.io/gorm/clause"
	"github.com/labstack/echo/v4"
	"github.com/MakotoNakai/lets-schedule/config"
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

	pui := c.Param("user_id")
	ui, err := strconv.Atoi(pui)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	pmi := c.Param("meeting_id")
	mi, err := strconv.Atoi(pmi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	oldCTList := []models.CandidateTime{}
	oldCTList = models.GetCandidateTimeByMeetingIdAndUserId(mi, ui)

	newCTList := []models.CandidateTime{}
	err = c.Bind(&newCTList)

	ctList := []models.CandidateTime{}

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Clauses(
		clause.Locking{Strength: "UPDATE"},
	).Find(&models.CandidateTime{})

	oldlen := len(oldCTList)
	newlen := len(newCTList)
	minlen := models.Min(oldlen, newlen)

	tx := db.Begin()

	for i := 0; i < minlen; i++ {
		oldct := oldCTList[i]
		newct := newCTList[i]
		tx.Model(&oldct).Updates(newct)
		ctList = append(ctList, oldct)
	}

	if oldlen < newlen {
		newct := newCTList[oldlen:newlen]
		db.Create(&newct)
	}

	if len(oldCTList) > len(newCTList) {
		for i := len(newCTList); i < len(oldCTList); i++ {
			oldp := oldCTList[i]
			tx.Delete(&oldp)
		}
	}

	tx.Commit()
	return c.JSON(http.StatusOK, ctList)

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

func GetAvailableTimeByMeetingId(c echo.Context) error {

	pmi := c.Param("meeting_id")
	mi, err := strconv.Atoi(pmi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	availableTimeList := models.GetAvailableTimeByMeetingId(mi)
	if models.AvailableTimeIsNotFound(availableTimeList) {
		return c.JSON(http.StatusBadRequest, config.AvailableTimeNotFound)
	}
	return c.JSON(http.StatusOK, availableTimeList)
}

