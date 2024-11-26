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

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user
// @Accept json
// @Produce json
// @Param user body models.User true "Details of candidate time"
// @Success 201 {object} models.User
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/signup [post]
func CreateUser(c echo.Context) error {

	newUser := models.User{}

	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindUser)
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

func GetUserById(c echo.Context) error {

	user := models.User{}
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindUser)
	}

	db.First(&user)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {

	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindUser)
	}

	db.Save(&user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {

	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindUser)
	}

	db.Delete(&user)
	return c.JSON(http.StatusNoContent, user)
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} string "New password"
// @Failure 400 {object} string "Error message"
// @Router /user/{id}/reset-password [post]
func ResetPassword(c echo.Context) error {

	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	np := models.NewPassword{}
	err = c.Bind(&np)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindUser)
	}

	err = models.ResetPassword(db, id, np.NewPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToResetPassword)
	}

	return c.JSON(http.StatusOK, np.NewPassword)

}
