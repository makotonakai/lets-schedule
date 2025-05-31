package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/gorilla/sessions"
  "github.com/labstack/echo-contrib/session"

	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/database"
	"github.com/MakotoNakai/lets-schedule/models"
)

var errorMessageList = []string{}
var db = database.Connect()

type JWTCustomClaims struct {
	Id       int    `json:"uid"`
	UserName string `json:"name"`
	jwt.StandardClaims
}

type Credential struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Token string `json:"token"`
}

// Login generates a authorization token if the login becomes successful
// @Summary Generate a authorization token if the login becomes successful
// @Description Generate a authorization token if the login becomes successful
// @Tags handler
// @Accept json
// @Produce json
// @Success 200 {object} Credential
// @Failure 401 {object} string "Error message"
// @Router /api/login [get]
func Login(c echo.Context) error {

	u := models.User{}
	err := c.Bind(&u)
	if err != nil {
		log.Fatal(err)
	}

	if models.IsUserNameEmptyOrNull(u) == true {
		errorMessageList = append(errorMessageList, config.UserNameIsEmpty.Error())
	}

	if models.IsPasswordEmptyOrNull(u) == true {
		errorMessageList = append(errorMessageList, config.PasswordIsEmpty.Error())
	}

	if models.ErrorsExist(errorMessageList) {
		return c.JSON(http.StatusBadRequest, errorMessageList)
	}

	user := models.User{}
	db.Where("user_name = ?", u.UserName).Find(&user)

	// if user doesn't exist
	if user.Id == 0 || u.Password != user.Password {
		return c.JSON(http.StatusUnauthorized, config.LoginFailed)
	}

	claims := JWTCustomClaims{
		Id:       user.Id,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingKey := []byte("secret")
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	session, _ := session.Get("session", c)
	session.Options = &sessions.Options{
		Path:     "/login",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	session.Values["id"] = strconv.Itoa(user.Id)
	session.Values["user_name"] = user.UserName
	session.Values["token"] = t
	session.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, Credential{
		Id:        user.Id,
		UserName: user.UserName,
		Token:     t,
	})

}

// Accessible returns "Accessible" when the api server is alive
// @Summary Return "Accessible" when the api server is alive
// @Description Return "Accessible" when the api server is alive
// @Tags handler
// @Success 200 {string} Accessible
// @Router /api/ [get]
func Accessible(c echo.Context) error {
	return c.JSON(http.StatusOK, "Accessible")
}

// Restricted returns a greeting if the given token is valid
// @Summary Return a greeting if the given token is valid
// @Description Return a greeting if the given token is valid
// @Tags handler
// @Success 200 {string} Accessible
// @Router /api/restricted [get]
func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)
	name := claims.UserName
	return c.JSON(http.StatusOK, "Welcome "+name+"!")
}
