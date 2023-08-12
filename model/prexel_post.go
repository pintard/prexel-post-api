package model

import "time"

type PrexelPost struct {
	UUID     int64     `json:"uuid"`
	Username string    `json:"username" binding:"required"`
	Contact  string    `json:"contact"`
	Code     string    `json:"code" binding:"required"`
	Date     time.Time `json:"date"`
}
