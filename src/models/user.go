package models

import (
	"time"
)

type User struct {	
	Id int `gorm:"primaryKey:autoIncrement" json:"id" param:"id"`	
	UserName string `gorm:"unique" json:"user_name"`
	EmailAddress string `gorm:"unique" json:"email_address"`
	Password string `gorm:"not null" json:"password"`
	IsAdmin	bool `gorm:"not null" json:"is_admin"`
	CanLogin bool `gorm:"not null" json:"can_login"`	
	CreatedAt time.Time `gorm:"autoCreateTime:int" json:"created_at"`	
	UpdatedAt time.Time `gorm:"autoUpdateTime:int" json:"updated_at"`	
}