package model

import "time"

type PrexelPost struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username" binding:"required"`
	Contact    string    `json:"contact"`
	ContactURL string    `json:"contact_url"`
	Code       string    `json:"code" binding:"required"`
	Date       time.Time `json:"date"`
}
