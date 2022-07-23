package models

import (
	"time"
)

type Meeting struct {
	Id int `gorm:"primaryKey:not null:autoIncrement" json:"id"`	
	Title string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Type string `json:"type"`
	Place string `json:"place"`
	Url string `json:"url"`
	IsConfirmed bool`gorm:"not null" json:"is_confirmed"`
	CreatedAt time.Time `gorm:"autoCreateTime:int" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:int" json:"updated_at"`
}
