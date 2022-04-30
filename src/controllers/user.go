package controllers

import (

	"time"
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
	"github.com/MakotoNakai/lets-schedule/database"
)

//----------
// Handlers
//----------
var db = database.Connect()

func CreateUser(c echo.Context) error {
	
	newUser := models.User{}

	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	db.Create(&newUser)
	return c.JSON(http.StatusCreated, newUser)
	
}

func GetUser(c echo.Context) error {

	user := models.User{}
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.First(&user)

	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {

	userList:= []models.User{}

	db.Find(&userList)
	
	return c.JSON(http.StatusOK, userList)

}

func UpdateUser(c echo.Context) error {

	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.UpdatedAt = time.Now()
	db.Save(&user)

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {

	user := models.User{}
	
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Delete(&user)

	return c.JSON(http.StatusNoContent, user)
}

