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

// CreateMeeting creates a new candidate time
// @Summary Create a new candidate time
// @Description Create a new candidate time
// @Accept json
// @Produce json
// @Param meeting body models.CandidateTime true "Details of candidate time"
// @Success 201 {object} models.CandidateTime
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/candidate_times/new [post]
func CreateCandidateTime(c echo.Context) error {
	
	newCandidateTime := []models.CandidateTime{}
	err := c.Bind(&newCandidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTime)
	}

	errorMessageListAboutCandidateTime := []string{}


	emptyCandidateTimeExists, err := models.EmptyCandidateTimeExists(&newCandidateTime);
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTime)
	}

	pastCandidateTimeExists, err := models.PastCandidateTimeExists(&newCandidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTime)
	}

	if emptyCandidateTimeExists {
		errorMessageListAboutCandidateTime = append(errorMessageListAboutCandidateTime, config.CandidateTimeIsEmpty.Error())
	}

	if pastCandidateTimeExists {
		errorMessageListAboutCandidateTime = append(errorMessageListAboutCandidateTime, config.CandidateTimeIsPast.Error())
	}

	errorsExist := models.ErrorsExist(errorMessageListAboutCandidateTime);

	if errorsExist {
		return c.JSON(http.StatusBadRequest, errorMessageListAboutCandidateTime)
	}

	err = db.Create(&newCandidateTime).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToCreateCandidateTime)
	}

	return c.JSON(http.StatusCreated, newCandidateTime)
	
}

// GetCandidateTimeWithUserNameByMeetingId gets a CandidateTimeWithUserName by meeting id
// @Summary Get a CandidateTimeWithUserName by meeting id
// @Description Get a CandidateTimeWithUserName by meeting id
// @Accept json
// @Produce json
// @Param meeting_id path int true "Meeting ID"
// @Success 200 {object} models.CandidateTimeWithUserName 
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/candidate_times/meeting/{meeting_id} [get]
func GetCandidateTimeWithUserNameByMeetingId(c echo.Context) error {

	MeetingIdString := c.Param("meeting_id")
	MeetingId, err := strconv.Atoi(MeetingIdString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	newCandidateTimeList := []models.CandidateTime{}
	err = db.Table("candidate_times").
		Select("candidate_times.*").
		Where("candidate_times.meeting_id = ?", MeetingId).
		Find(&newCandidateTimeList).Error

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

// GetCandidateTimeByUserIdAndMeetingId gets a candidate time by user id and meeting id
// @Summary Get a candidate time by user id and meeting id
// @Description Get a candidate time by user id and meeting id
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param meeting_id path int true "Meeting ID"
// @Success 200 {object} models.CandidateTime
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/candidate_times/user/{user_id}/meeting/{meeting_id} [get]
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
	CandidateTimeList, err = models.GetCandidateTimeByMeetingIdAndUserId(db, MeetingId, UserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrUserNotFound)
	}

	return c.JSON(http.StatusOK, CandidateTimeList)

}

// UpdateCandidateTimeByUserIdAndMeetingId updates a candidate time by user id and meeting id
// @Summary Update a candidate time by user id and meeting id
// @Description Update a candidate time by user id and meeting id
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param meeting_id path int true "Meeting ID"
// @Success 200 {object} models.CandidateTime
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/candidate_times/user/{user_id}/meeting/{meeting_id} [put]
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
	oldCTList, err = models.GetCandidateTimeByMeetingIdAndUserId(db, mi, ui)
	if err != nil {
			return c.JSON(http.StatusBadRequest, config.ErrUserNotFound)
	}

	newCTList := []models.CandidateTime{}
	err = c.Bind(&newCTList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTimeList)
	}

	errorMessageListAboutCandidateTime := []string{}

	IsCandidateTimeEmpty, err := models.EmptyCandidateTimeExists(&newCTList);
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTime)
	}

	pastCandidateTimeExists, err := models.PastCandidateTimeExists(&newCTList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTime)
	}

	if IsCandidateTimeEmpty {
		errorMessageListAboutCandidateTime = append(errorMessageListAboutCandidateTime, config.CandidateTimeIsEmpty.Error())
	}

	if pastCandidateTimeExists {
		errorMessageListAboutCandidateTime = append(errorMessageListAboutCandidateTime, config.CandidateTimeIsPast.Error())
	}

	errorsExist := models.ErrorsExist(errorMessageListAboutCandidateTime)

	if errorsExist {
		return c.JSON(http.StatusBadRequest, errorMessageListAboutCandidateTime)
	}

	ctList := []models.CandidateTime{}

	db.Clauses(
		clause.Locking{Strength: "UPDATE"},
	).Find(&models.CandidateTime{})

	oldlen := len(oldCTList)
	newlen := len(newCTList)
	minlen := models.Min(oldlen, newlen)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIntegerIsNil)
	}

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

	err = tx.Commit().Error
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

	err = db.Delete(&CandidateTime).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToDeleteCandidateTime)
	}

	return c.JSON(http.StatusNoContent, CandidateTime)

}

// GetAvailableTimeByMeetingId gets an available time by meeting id
// @Summary Get an available time by meeting id
// @Description Get an available time by meeting id
// @Accept json
// @Produce json
// @Param meeting_id path int true "Meeting ID"
// @Success 200 {object} models.CandidateTime
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/candidate_times/available-time/{meeting_id} [get]
func GetAvailableTimeByMeetingId(c echo.Context) error {

	mistr := c.Param("meeting_id")
	mi, err := strconv.Atoi(mistr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	availableTimeList, err := models.GetAvailableTimeByMeetingId(db, mi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.MeetingNotFound)
	}

	availableTimeNotFound, _ := models.AvailableTimeIsNotFound(&availableTimeList)
	if availableTimeNotFound {
		return c.JSON(http.StatusBadRequest, config.AvailableTimeNotFound)
	}
	return c.JSON(http.StatusOK, availableTimeList)
}

// UpdateAvailableTimeByMeetingId gets an available time by meeting id
// @Summary Update an available time by meeting id
// @Description Update an available time by meeting id
// @Accept json
// @Produce json
// @Param meeting_id path int true "Meeting ID"
// @Success 200 {object} models.CandidateTime
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/candidate_times/available-time/{meeting_id} [put]
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
		return c.JSON(http.StatusBadRequest, config.MeetingNotFound)
	}

	errorMessageListAboutAvailableTime := []string{}

	availableTimeEmpty, err := models.IsAvailableTimeEmpty(&availableTime);
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.AvailableTimeNotFound)
	}

	availableTimeMoreThanExpected, err := IsAvailableTimeMoreThanExpected(availableTime, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTime)
	}

	availableTimeNotWithinLimit, err := IsAvailableTimeWithinLimit(availableTime, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindCandidateTime)
	}

	if availableTimeEmpty {
		errorMessageListAboutAvailableTime = append(errorMessageListAboutAvailableTime, config.AvailableTimeNotFound.Error())
	}

	if availableTimeMoreThanExpected {
		errorMessageListAboutAvailableTime = append(errorMessageListAboutAvailableTime, config.AvailableTimeTooLong.Error())
	}

	if availableTimeNotWithinLimit {
		errorMessageListAboutAvailableTime = append(errorMessageListAboutAvailableTime, config.AvailableTimeOutOfLimit.Error())
	}

	errorsExist := models.ErrorsExist(errorMessageListAboutAvailableTime)

	if errorsExist {
		return c.JSON(http.StatusBadRequest, errorMessageListAboutAvailableTime)
	}

	// Updating only the ActualStartTime and ActualEndTime fields
	db.Model(&oldMeeting).Updates(models.AvailableTime{
		ActualStartTime: availableTime.ActualStartTime,
		ActualEndTime:   availableTime.ActualEndTime,
	})

	return c.JSON(http.StatusOK, oldMeeting)
}

func IsAvailableTimeMoreThanExpected(availableTime models.AvailableTime, id int) (bool, error) {

	meeting := models.Meeting{}
	err := db.First(&meeting, id).Error
	if err != nil {
		return false, err
	}

	startTime := availableTime.ActualStartTime
	endTime := availableTime.ActualEndTime

	duration := endTime.Sub(startTime)
	hours := duration.Hours()

	return hours > float64(meeting.Hour), nil

}

func IsAvailableTimeWithinLimit(availableTime models.AvailableTime, id int) (bool, error) {
	availableTimeList, err := models.GetAvailableTimeByMeetingId(db, id)
	if err != nil {
		return false, err
	}

	_, err = models.AvailableTimeIsNotFound(&availableTimeList)
	if err != nil {
		return false, nil
	}

	for _, at := range availableTimeList {
		if at.StartTime.Before(availableTime.ActualEndTime) && at.EndTime.Before(availableTime.ActualStartTime) {
			return true, nil
		}
	}
	return false, nil
}


