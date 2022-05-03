package controllers

import (

	"time"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
)

func CreateCandidateTime(c echo.Context) error {
	
	newCandidateTime := models.CandidateTime{}
	err := c.Bind(&newCandidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newCandidateTime.CreatedAt = time.Now()
	newCandidateTime.UpdatedAt = time.Now()

	db.Create(&newCandidateTime)
	return c.JSON(http.StatusCreated, newCandidateTime)
	
}

func GetCandidateTime(c echo.Context) error {

	candidateTime := models.CandidateTime{}
	err := c.Bind(&candidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db.First(&candidateTime)

	return c.JSON(http.StatusOK, candidateTime)
}

func GetCandidateTimes(c echo.Context) error {

	candidateTimeList:= []models.CandidateTime{}
	db.Find(&candidateTimeList)
	
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

