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
	
	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Create(&newUser)
	return c.JSON(http.StatusCreated, newUser)
	
}

func GetAllUser(c echo.Context) error {

	userList:= []models.User{}

	db.Find(&userList)
	return c.JSON(http.StatusOK, userList)

}

func GetUserById(c echo.Context) error {

	user := models.User{}
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.First(&user)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {

	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

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

func ResetPassword(c echo.Context) error {

	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	np := models.NewPassword{}
	err = c.Bind(&np)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = models.ResetPassword(id, np.NewPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, np.NewPassword)

}
