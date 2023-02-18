package models

import (
	// "fmt"
	"time"
	"github.com/MakotoNakai/lets-schedule/database"
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

var db = database.Connect()

func IsEmailAddressEmpty(u User) bool {
	return u.EmailAddress == ""
}

func IsUserNameEmpty(u User) bool {
	return u.UserName == ""
}

func IsPasswordEmpty(u User) bool {
	return u.Password == ""
}

func ErrorsExist(errorMessageList []string) bool {
	return len(errorMessageList) != 0
}


func AlreadyExists(u User) bool {

	var sameEmailAddress User
	var sameUserName User

	db.Table("users").Select("*").Where("users.email_address = ?", u.EmailAddress).Find(&sameEmailAddress)
	db.Table("users").Select("*").Where("users.email_address = ?", u.UserName).Find(&sameUserName)

	// If the user with either the given email addresss or the given username exists, returns true
	if sameEmailAddress.Id == 0 && sameUserName.Id == 0 {
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



