package controllers

import (

	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
)

//----------
// Handlers
//----------

func CreateFriend(c echo.Context) error {
	
	newFriendWithUserName := models.FriendWithUserName{}
	err := c.Bind(&newFriendWithUserName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	UserName := newFriendWithUserName.UserName
	FriendUserName := newFriendWithUserName.FriendUserName

	newFriend := models.Friend{}
	newFriend.UserId = models.GetUserIdFromUserName(UserName)
	newFriend.FriendUserId = models.GetUserIdFromUserName(FriendUserName)

	db.Create(&newFriend)
	return c.JSON(http.StatusCreated, newFriend)
	
}

