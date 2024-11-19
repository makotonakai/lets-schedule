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
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTime)
	}

	errorMessageListAboutCandidateTime := []string{}

	if models.IsCandidateTimeEmpty(newCandidateTime) == true {
		errorMessageListAboutCandidateTime = append(errorMessageListAboutCandidateTime, config.CandidateTimeIsEmpty)
	}

	if models.EmptyCandidateTimeExists(newCandidateTime) == true {
		errorMessageListAboutCandidateTime = append(errorMessageListAboutCandidateTime, config.CandidateTimeIsEmpty)
	}

	if models.PastCandidateTimeExists(newCandidateTime) == true {
		errorMessageListAboutCandidateTime = append(errorMessageListAboutCandidateTime, config.CandidateTimeIsPast)
	}

	if models.ErrorsExist(errorMessageListAboutCandidateTime) {
		return c.JSON(http.StatusBadRequest, errorMessageListAboutCandidateTime)
	}

	err = db.Create(&newCandidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToCreateCandidateTime)
	}

	return c.JSON(http.StatusCreated, newCandidateTime)
	
}

func GetCandidateTimeWithUserNameByMeetingId(c echo.Context) error {

	MeetingIdString := c.Param("meeting_id")
	MeetingId, err := strconv.Atoi(MeetingIdString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	newCandidateTimeList := []models.CandidateTime{}
	err := db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Find(&newCandidateTimeList)

	if err != nil {
		c.JSON(http.StatusBadRequest, config.ErrCandidateTimeNotFound)
	}

	CandidateTimeWithUserNameList := []models.CandidateTimeWithUserName{}

		for _, CandidateTime := range newCandidateTimeList {

			CandidateTimeWithUserName := models.CandidateTimeWithUserName{}
			UserId := CandidateTime.UserId

			CandidateTimeWithUserName.UserName, _ = models.GetUserNameFromUserId(db, UserId)
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
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	MeetingIdString := c.Param("meeting_id")
	MeetingId, err := strconv.Atoi(MeetingIdString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	CandidateTimeList := []models.CandidateTime{}
	CandidateTimeList = models.GetCandidateTimeByMeetingIdAndUserId(db, MeetingId, UserId)

	return c.JSON(http.StatusOK, CandidateTimeList)

}

func UpdateCandidateTimeByUserIdAndMeetingId(c echo.Context) error {

	pui := c.Param("user_id")
	ui, err := strconv.Atoi(pui)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}
	
	pmi := c.Param("meeting_id")
	mi, err := strconv.Atoi(pmi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	oldCTList := []models.CandidateTime{}
	oldCTList = models.GetCandidateTimeByMeetingIdAndUserId(db, mi, ui)

	newCTList := []models.CandidateTime{}
	err = c.Bind(&newCTList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTimeList)
	}

	errorMessageListAboutCandidateTime := []string{}

	if models.IsCandidateTimeEmpty(newCTList) == true {
		errorMessageListAboutCandidateTime = append(errorMessageListAboutCandidateTime, config.CandidateTimeIsEmpty)
	}

	if models.EmptyCandidateTimeExists(newCTList) == true {
		errorMessageListAboutCandidateTime = append(errorMessageListAboutCandidateTime, config.CandidateTimeIsEmpty)
	}

	if models.PastCandidateTimeExists(newCTList) == true {
		errorMessageListAboutCandidateTime = append(errorMessageListAboutCandidateTime, config.CandidateTimeIsPast)
	}

	if models.ErrorsExist(errorMessageListAboutCandidateTime) {
		return c.JSON(http.StatusBadRequest, errorMessageListAboutCandidateTime)
	}

	ctList := []models.CandidateTime{}

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
		err := tx.Model(&oldct).Updates(newct)
		if err != nil {
			return c.JSON(http.StatusBadRequest, config.ErrFailedToUpdateCandidateTime)
		}
		ctList = append(ctList, oldct)
	}

	if oldlen < newlen {
		newct := newCTList[oldlen:newlen]
		err := db.Create(&newct)
		if err != nil {
			return c.JSON(http.StatusBadRequest, config.ErrFailedToCreateCandidateTime)
		}
	}

	if len(oldCTList) > len(newCTList) {
		for i := len(newCTList); i < len(oldCTList); i++ {
			oldp := oldCTList[i]
			err := tx.Delete(&oldp)
			if err != nil {
				return c.JSON(http.StatusBadRequest, config.ErrFailedToDeleteCandidateTime)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToExecuteTransaction)
	}

	return c.JSON(http.StatusOK, ctList)

}

func DeleteCandidateTime(c echo.Context) error {

	CandidateTime := models.CandidateTime{}
	
	err := c.Bind(&CandidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTime)
	}

	err = db.Delete(&CandidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToDeleteCandidateTime)
	}

	return c.JSON(http.StatusNoContent, CandidateTime)

}

func GetAvailableTimeByMeetingId(c echo.Context) error {

	pmi := c.Param("meeting_id")
	mi, err := strconv.Atoi(pmi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}
	availableTimeList := models.GetAvailableTimeByMeetingId(db, mi)
	if models.AvailableTimeIsNotFound(availableTimeList) {
		return c.JSON(http.StatusBadRequest, config.AvailableTimeNotFound)
	}
	return c.JSON(http.StatusOK, availableTimeList)
}

func UpdateAvailableTimeByMeetingId(c echo.Context) error {
	paramId := c.Param("meeting_id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	oldMeeting := models.Meeting{}
	if err := db.First(&oldMeeting, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, config.MeetingNotFound)
	}

	availableTime := models.AvailableTime{}
	if err := c.Bind(&availableTime); err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTime)
	}

	errorMessageListAboutAvailableTime := []string{}

	if models.IsAvailableTimeEmpty(availableTime) == true {
		errorMessageListAboutAvailableTime = append(errorMessageListAboutAvailableTime, config.AvailableTimeNotFound)
	}


	if IsAvailableTimeMoreThanExpected(availableTime, id) {
		errorMessageListAboutAvailableTime = append(errorMessageListAboutAvailableTime, config.AvailableTimeTooLong)
	}

	if IsAvailableTimeWithinLimit(availableTime, id) {
		errorMessageListAboutAvailableTime = append(errorMessageListAboutAvailableTime, config.AvailableTimeOutOfLimit)
	}

	if models.ErrorsExist(errorMessageListAboutAvailableTime) {
		return c.JSON(http.StatusBadRequest, errorMessageListAboutAvailableTime)
	}

	// Updating only the ActualStartTime and ActualEndTime fields
	db.Model(&oldMeeting).Updates(models.AvailableTime{
		ActualStartTime: availableTime.ActualStartTime,
		ActualEndTime:   availableTime.ActualEndTime,
	})

	return c.JSON(http.StatusOK, oldMeeting)
}

func IsAvailableTimeMoreThanExpected(availableTime models.AvailableTime, id int) bool {

	meeting := models.Meeting{}
	err := db.First(&meeting, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrMeetingNotFound)
	}

	startTime := availableTime.ActualStartTime
	endTime := availableTime.ActualEndTime

	duration := endTime.Sub(startTime)
	hours := duration.Hours()

	return hours > float64(meeting.Hour)

}

func IsAvailableTimeWithinLimit(availableTime models.AvailableTime, id int) bool {
	availableTimeList := models.GetAvailableTimeByMeetingId(db, id)
	if models.AvailableTimeIsNotFound(availableTimeList) {
		return false
	}

	for _, at := range availableTimeList {
		if at.StartTime.Before(availableTime.ActualEndTime) && at.EndTime.Before(availableTime.ActualStartTime) {
			return true
		}
	}
	return false
}


