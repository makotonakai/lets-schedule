package controllers

import (

	"time"
	"strconv"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
)

func CreateCandidateTimeList(c echo.Context) error {
	
	newCandidateTimeList := []models.CandidateTime{}
	err := c.Bind(&newCandidateTimeList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	for _, newCandidateTime := range newCandidateTimeList {
		newCandidateTime.CreatedAt = time.Now()
		newCandidateTime.UpdatedAt = time.Now()
	}
	

	db.Create(&newCandidateTimeList)
	return c.JSON(http.StatusCreated, newCandidateTimeList)
	
}

func GetCandidateTimeByUserIdAndMeetingId(c echo.Context) error {

	userId, _ := strconv.Atoi(c.Param("user_id"))
	meetingId, _ := strconv.Atoi(c.Param("meeting_id"))
	candidateTimeList:= []models.CandidateTime{}
	db.Where("meeting_id=?", meetingId).Where("user_id=?", userId).Find(&candidateTimeList)
	
	return c.JSON(http.StatusOK, candidateTimeList)

}

func UpdateCandidateTime(c echo.Context) error {

	candidateTime := models.CandidateTime{}
	err := c.Bind(&candidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	candidateTime.UpdatedAt = time.Now()
	db.Save(&candidateTime)

	return c.JSON(http.StatusOK, candidateTime)
}

func DeleteCandidateTime(c echo.Context) error {

	candidateTime:= models.CandidateTime{}
	err := c.Bind(&candidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Delete(&candidateTime)

	return c.JSON(http.StatusNoContent, candidateTime)
}

