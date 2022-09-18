package models

import (
	"time"
)

type User struct {	
	Id int `gorm:"primaryKey:not null:autoIncrement" json:"id"`	
	UserName string `gorm:"not null:unique" json:"user_name"`
	EmailAddress string `gorm:"not null:unique" json:"email_address"`
	Password string `gorm:"not null" json:"password"`
	IsAdmin	bool `gorm:"not null" json:"is_admin"`
	CanLogin bool `gorm:"not null" json:"can_login"`	
	CreatedAt time.Time `gorm:"autoCreateTime:int" json:"created_at"`	
	UpdatedAt time.Time `gorm:"autoUpdateTime:int" json:"updated_at"`	
}

func GetUserIdFromUserName(UserName string) int {
	User := User{}
	db.Table("users").
		Select("users.id").
		Where("users.user_name = ?", UserName).
		Find(&User)
	return User.Id
}

func GetUserNameFromUserId(UserId int) string {
	User := User{}
	db.Table("users").
		Select("users.user_name").
		Where("users.id = ?", UserId).
		Find(&User)
	return User.UserName
}


