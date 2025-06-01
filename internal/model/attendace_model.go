package model

import "time"

type AttendanceResponse struct {
	Status string    `json:"status"`
	Time   time.Time `json:"date"`
}

type AttendanceCreateResponse struct {
	UserId uint `json:"username" validate:"required"`
}
