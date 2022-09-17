package models

import (
	"time"
)

type Friend struct {	
	Id int `gorm:"primaryKey:not null:autoIncrement" json:"id"`	
	UserId int `gorm:"not null" json:"user_id"`
	FriendUserId int `gorm:"not null" json:"friend_user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime:int" json:"created_at"`	
	UpdatedAt time.Time `gorm:"autoUpdateTime:int" json:"updated_at"`	
}
