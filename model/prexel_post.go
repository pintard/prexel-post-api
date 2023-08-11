package model

type PrexelPost struct {
	UUID     int64  `json:"uuid"`
	Username string `json:"username" binding:"required"`
	Contact  string `json:"contact"`
	Content  string `json:"content" binding:"required"`
}