package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/database"
	"github.com/MakotoNakai/lets-schedule/models"
)

// ----------
// Handlers
// ----------
var db = database.Connect()
var errorMessageList = []string{}

func CreateUser(c echo.Context) error {

	newUser := models.User{}

	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if models.IsEmailAddressEmptyOrNull(newUser) == true {
		errorMessageList = append(errorMessageList, config.EmailAddressIsEmpty)
	}

	if models.IsUserNameEmptyOrNull(newUser) == true {
		errorMessageList = append(errorMessageList, config.UserNameIsEmpty)
	}

	if models.IsPasswordEmptyOrNull(newUser) == true {
		errorMessageList = append(errorMessageList, config.PasswordIsEmpty)
	}

	userExist, _, _ := models.AlreadyExists(db, newUser)
	if userExist == true {
		errorMessageList = append(errorMessageList, config.UserAlreadyExists)
	}

	if models.ErrorsExist(errorMessageList) {
		return c.JSON(http.StatusBadRequest, errorMessageList)
	}

	db.Create(&newUser)
	return c.JSON(http.StatusCreated, newUser)

}

func GetAllUser(c echo.Context) error {

	userList := []models.User{}

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

	err = models.ResetPassword(db, id, np.NewPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, np.NewPassword)

}
