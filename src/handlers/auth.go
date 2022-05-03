package handlers

import (

	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
	"github.com/MakotoNakai/lets-schedule/database"
)

func BasicAuth (username, password string, c echo.Context) (bool, error) {

	db := database.Connect()

	user := models.User{}
	db.Where("user_name = ? AND password = ?", username, password).Find(&user)

	// if user doesn't exist
	if user.Id == 0 {
		return false, nil
	}

	return true, nil
}