package models

import (
	"fmt"
	"time"
	"errors"
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
	} else {
		u.EmailAddress = strings.ReplaceAll(u.EmailAddress, " ", "")
		return u.EmailAddress == ""
	}
}

func IsUserNameEmptyOrNull(u User) bool {
	if govalidator.IsNull(u.UserName) {
		return true
	} else {
		u.UserName = strings.ReplaceAll(u.UserName, " ", "")
		return u.UserName == ""
	}
}

func IsPasswordEmptyOrNull(u User) bool {
	if govalidator.IsNull(u.Password) {
		return true
	} else {
		u.Password = strings.ReplaceAll(u.Password, " ", "")
		return u.Password == ""
	}
}

func ErrorsExist(errorMessageList []string) bool {
	return len(errorMessageList) != 0
}


func AlreadyExists(db *gorm.DB, u User) (bool, error, error) {

	var sameEmailAddress User
	var sameUserName User

	emailAddressErr := db.Table("users").Select("*").Where("users.email_address = ?", u.EmailAddress).Find(&sameEmailAddress).Error
	userNameErr := db.Table("users").Select("*").Where("users.user_name = ?", u.UserName).Find(&sameUserName).Error

	if emailAddressErr != nil && userNameErr != nil {
		return false, errors.New("Email address not found"), errors.New("Username not found")
	} else if emailAddressErr != nil {
		return true, errors.New("Email address not found"),  nil
	} else if userNameErr != nil {
		return true, nil, errors.New("Username not found")
	} else {
		return true, nil, nil
	}
}

func GetUserIdFromUserName(db *gorm.DB, UserName string) (int, error) {
	user := User{}
	err := db.Table("users").
			Select("users.id").
			Where("users.user_name = ?", UserName).
			Find(&user).Error

	// Check if user was found
	if err != nil {
			return -1, errors.New(fmt.Sprintf("user with username '%s' not found", UserName))
	}

	return user.Id, nil
}

func GetUserNameFromUserId(db *gorm.DB, UserId int) (string, error) {
	user := User{}
	err := db.Table("users").
		Select("users.user_name").
		Where("users.id = ?", UserId).
		Find(&user).Error
	if err != nil {
		return "", errors.New(fmt.Sprintf("user with id '%d' not found", UserId))
	}
	return user.UserName, nil
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



