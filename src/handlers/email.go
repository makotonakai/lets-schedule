package handlers

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
  gomail "gopkg.in/gomail.v2"

	"github.com/MakotoNakai/lets-schedule/models"
)


func SendEmail(c echo.Context) error {

	ea := models.EmailAddress{}
	err := c.Bind(&ea)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ui := models.GetUserIdFromEmailAddress(ea.EmailAddress)

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
