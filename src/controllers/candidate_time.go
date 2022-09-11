package controllers

import (

	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
)

//----------
// Handlers
//----------

func CreateCandidateTime(c echo.Context) error {
	
	newCandidateTime := models.CandidateTime{}
	
	err := c.Bind(&newCandidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Create(&newCandidateTime)
	return c.JSON(http.StatusCreated, newCandidateTime)
	
}

func GetAllCandidateTime(c echo.Context) error {

	CandidateTimeList:= []models.CandidateTime{}

	db.Find(&CandidateTimeList)
	return c.JSON(http.StatusOK, CandidateTimeList)

}

func GetCandidateTimeById(c echo.Context) error {

	CandidateTime := models.CandidateTime{}
	err := c.Bind(&CandidateTime)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.First(&CandidateTime)
	return c.JSON(http.StatusOK, CandidateTime)

}

func UpdateCandidateTime(c echo.Context) error {

	CandidateTime := models.CandidateTime{}

	err := c.Bind(&CandidateTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Save(&CandidateTime)
	return c.JSON(http.StatusOK, CandidateTime)

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
