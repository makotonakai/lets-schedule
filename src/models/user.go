package models

import (
	"time"
	"regexp"
	"strings"
	"gorm.io/gorm"
	"github.com/MakotoNakai/lets-schedule/config"
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

	emailRegex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return emailRegex.MatchString(e)

}

func IsEmailAddressEmptyOrNull(u User) bool {

	u.EmailAddress = strings.ReplaceAll(u.EmailAddress, " ", "")
	return u.EmailAddress == ""

}

func IsUserNameEmptyOrNull(u User) bool {

	u.UserName = strings.ReplaceAll(u.UserName, " ", "")
	return u.UserName == ""

}

func IsPasswordEmptyOrNull(u User) bool {

	u.Password = strings.ReplaceAll(u.Password, " ", "")
	return u.Password == ""

}

func ErrorsExist(errorMessageList []string) bool {
	return len((errorMessageList)) != 0
}


func AlreadyExists(db *gorm.DB, u User) (bool, error, error) {

	var sameEmailAddress User
	var sameUserName User

	emailAddressErr := db.Table("users").Select("*").Where("users.email_address = ?", u.EmailAddress).Find(&sameEmailAddress).Error
	userNameErr := db.Table("users").Select("*").Where("users.user_name = ?", u.UserName).Find(&sameUserName).Error

	if emailAddressErr != nil && userNameErr != nil {
		return false, config.ErrEmailAddressNotFound, config.ErrUserNameNotFound
	} else if emailAddressErr != nil {
		return true, config.ErrEmailAddressNotFound,  nil
	} else if userNameErr != nil {
		return true, nil, config.ErrUserNameNotFound
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
			return -1, config.ErrUserWithUserNameNotFound
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
		return "", config.ErrUserWithUserIdNotFound
	}
	return user.UserName, nil
}

func GetUserIdFromEmailAddress(db *gorm.DB, EmailAddress string) (int, error) {
	User := User{}
	err := db.Table("users").
		Select("users.id").
		Where("users.email_address = ?", EmailAddress).
		Find(&User).Error

	if err != nil {
		return -1, config.ErrUserIdWithEmailAddressNotFound
	}

	return User.Id, nil
}

func ResetPassword(db *gorm.DB, Id int, NewPassword string) error {

	err := db.Model(&User{}).Where("id = ?", Id).Update("password", NewPassword).Error
	if err != nil {
		return config.ErrFailedToResetPassword
	}
	
	return nil
}



