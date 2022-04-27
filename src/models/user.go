package models

import (
	"time"
)

type User struct {
	ID int								`json:"id"`
	UserName string 			`json:"userName"`
	EmailAddress string 	`json:"emailAddress"`
	Password string 			`json:"password"`
	IsAdmin bool 					`json:"isAdmin"`
	CanLogin bool 				`json:"canLogin"`
	CreatedAt time.Time		`json:"CreatedAt"`
	UpdatedAt time.Time		`json:"UpdatedAt"`
}