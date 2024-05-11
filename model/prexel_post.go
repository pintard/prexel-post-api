package model

import "time"

type PrexelPost struct {
	ID     int64     `json:"id"`
	UserId int64     `json:"user_id" binding:"required"`
	Code   string    `json:"code" binding:"required"`
	Title  string    `json:"title" binding:"required"`
	Tags   []string  `json:"tags"`
	Image  string    `json:"image"`
	Date   time.Time `json:"date"`
}
