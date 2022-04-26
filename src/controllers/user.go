package controllers

import (

	"strconv"
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
	err := c.Bind(newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Create(&newUser)
	return c.JSON(http.StatusCreated, newUser)
	
}

func GetUser(c echo.Context) error {

	user := models.User{}
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&user, id)

	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {

	userList:= []models.User{}
	users := db.Find(&userList)

	return c.JSON(http.StatusOK, users)

}

func UpdateUser(c echo.Context) error {

	user := models.User{}
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&user, id)
	db.Save(&user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {

	user := models.User{}
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, _ := strconv.Atoi(c.Param("id"))
	db.Delete(&user, id)
	return c.JSON(http.StatusNoContent, user)
}

