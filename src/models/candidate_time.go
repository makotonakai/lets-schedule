package models

import (
	"time"
)

type CandidateTime struct {
	Id int `gorm:"primaryKey:autoIncrement" json:"id"`
	StartTime string `gorm:"not null" json:"start_time"`
	EndTime string `gorm:"not null" json:"end_time"`
	CreatedAt time.Time `gorm:"autoCreateTime:int" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:int" json:"updated_at"`
}
