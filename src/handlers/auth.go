package handlers

import (

	"log"
	"time"
	"strconv"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
	"github.com/MakotoNakai/lets-schedule/database"
)

type JWTCustomClaims struct {
	Id int `json:"uid"`
	UserName string `json:"name"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {

	db := database.Connect()

	u := models.User{}
	err := c.Bind(&u)
	if err != nil {
		log.Fatal(err);
	}

	user := models.User{}
	db.Where("user_name = ?", u.UserName).Find(&user)

	// if user doesn't exist
	if user.Id == 0 || u.Password != user.Password {
		return c.JSON(http.StatusUnauthorized, user)
	}

	claims := JWTCustomClaims{
		Id: user.Id,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	signingKey := []byte("secret")
	t, err := token.SignedString(signingKey)
	if err != nil {
			return err
	}

	return c.JSON(http.StatusOK, map[string]string{
			"id": strconv.Itoa(user.Id),
			"user_name": user.UserName,
			"token": t,
	})

}

func Accessible(c echo.Context) error {
	return c.JSON(http.StatusOK, "Accessible")
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)
	name := claims.UserName
	return c.JSON(http.StatusOK, "Welcome " + name + "!")
}
