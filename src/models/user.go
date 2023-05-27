package models

import (
	"time"
	"regexp"
	"strings"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
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

func IsEmailAddressValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(e)
}

func IsEmailAddressEmptyOrNull(u User) bool {
	if govalidator.IsNull(u.EmailAddress) {
		return true
	}else {
		u.EmailAddress = strings.ReplaceAll(u.EmailAddress, " ", "")
		return u.EmailAddress == ""
	}
}

func IsUserNameEmptyOrNull(u User) bool {
	if govalidator.IsNull(u.UserName) {
		return true
	}else {
		u.EmailAddress = strings.ReplaceAll(u.UserName, " ", "")
		return u.UserName == ""
	}
}

func IsPasswordEmptyOrNull(u User) bool {
	if govalidator.IsNull(u.Password) {
		return true
	}else {
		u.EmailAddress = strings.ReplaceAll(u.Password, " ", "")
		return u.Password == ""
	}
}

func ErrorsExist(errorMessageList []string) bool {
	return len(errorMessageList) != 0
}


func AlreadyExists(db *gorm.DB, u User) bool {

	var sameEmailAddress User
	var sameUserName User

	db.Table("users").Select("*").Where("users.email_address = ?", u.EmailAddress).Find(&sameEmailAddress)
	db.Table("users").Select("*").Where("users.user_name = ?", u.UserName).Find(&sameUserName)

	// If the user with either the given email addresss or the given username exists, returns true
	if sameEmailAddress.Id == 0 && sameUserName.Id == 0 {
		return false
	}
	return true
}

func GetUserIdFromUserName(db *gorm.DB, UserName string) int {
	user := User{}
	db.Table("users").
		Select("users.id").
		Where("users.user_name = ?", UserName).
		Find(&user)
	return user.Id
}

func GetUserNameFromUserId(db *gorm.DB, UserId int) string {
	user := User{}
	db.Table("users").
		Select("users.user_name").
		Where("users.id = ?", UserId).
		Find(&user)
	return user.UserName
}

func GetUserIdFromEmailAddress(db *gorm.DB, EmailAddress string) int {
	User := User{}
	db.Table("users").
		Select("users.id").
		Where("users.email_address = ?", EmailAddress).
		Find(&User)
	return User.Id
}

func ResetPassword(db *gorm.DB, Id int, NewPassword string) error {

	err := db.Model(&User{}).Where("id = ?", Id).Update("password", NewPassword).Error
	return err
}



