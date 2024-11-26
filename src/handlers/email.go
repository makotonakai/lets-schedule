package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	gomail "gopkg.in/gomail.v2"

	"github.com/MakotoNakai/lets-schedule/models"
)

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user
// @Accept json
// @Produce json
// @Param user body models.EmailAddress true "Details of email address"
// @Success 201 {object} models.EmailAddress
// @Failure 400 {object} string "Error message"
// @Router /api/send-email [post]
func SendEmail(c echo.Context) error {

	ea := models.EmailAddress{}
	err := c.Bind(&ea)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ui := models.GetUserIdFromEmailAddress(db, ea.EmailAddress)

	m := gomail.NewMessage()
	m.SetHeader("From", "from@email.com")
	m.SetHeader("To", ea.EmailAddress)
	m.SetHeader("Subject", "Please reset your password for Let's Schedule")
	m.SetBody("text/plain", fmt.Sprintf("http://localhost:3000/%d/reset-password", ui))

	d := gomail.Dialer{Host: "localhost", Port: 1025}
	err = d.DialAndSend(m)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, "There is some problem with sending your email")
	}

	return c.JSON(http.StatusOK, "Email was sent successfully!")
}
