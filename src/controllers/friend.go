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
	userId, err := models.GetUserIdFromUserName(db, UserName)
	newFriend.UserId = userId
	newFriendId, err := models.GetUserIdFromUserName(db, FriendUserName)
	newFriend.FriendUserId = newFriendId

	db.Create(&newFriend)
	return c.JSON(http.StatusCreated, newFriend)
	
}

