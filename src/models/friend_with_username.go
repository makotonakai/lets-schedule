package models

import (
	"time"
)

type FriendWithUserName struct {	
	Id int `json:"id"`	
	UserName string `json:"user_name"`
	FriendUserName string `json:"friend_user_name"`
	CreatedAt time.Time `json:"created_at"`	
	UpdatedAt time.Time `json:"updated_at"`	
}
