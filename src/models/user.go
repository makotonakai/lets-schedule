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

type EmailAddress struct {
	EmailAddress string `json:"email_address"`
}

type NewPassword struct {
	NewPassword string `json:"new_password"`
}

func AlreadyExist(u User) bool {
	
	err := db.First(&u, "user_name = ?", u.UserName).Error

	if err != nil {
		return false
	}

	return true

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

func GetUserIdFromEmailAddress(EmailAddress string) int {
	User := User{}
	db.Table("users").
		Select("users.id").
		Where("users.email_address = ?", EmailAddress).
		Find(&User)
	return User.Id
}

func ResetPassword(Id int, NewPassword string) error {
	err := db.Model(&User{}).Where("id = ?", Id).Update("password", NewPassword).Error
	return err
}



