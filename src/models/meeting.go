package models

import (
	"time"
)

type Meeting struct {
	Id int `json:"id" param:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Type string `json:"type"`
	Place string `json:"place"`
	Url string `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}