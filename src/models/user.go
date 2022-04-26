package models

import (
	"time"
)

type User struct {
	ID uint
	UserName string 			`gorm:"unique" gorm:"not null" json:"userName"`
	EmailAddress string 	`gorm:"unique" gorm:"not null" json:"emailAddress"`
	Password string 			`gorm:"not null" json:"password"`
	IsAdmin bool 					`gorm:"not null" json:"isAdmin"`
	CanLogin bool 				`gorm:"not null" json:"canLogin"`
	CreatedAt time.Time		`json:"CreatedAt"`
	UpdatedAt time.Time		`json:"UpdatedAt"`
}