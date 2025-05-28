package model

import "time"

type AttendanceResponse struct {
	Status string    `gorm:"column:status"`
	Time   time.Time `gorm:"column:date"`
}

type AttendanceCreateResponse struct {
	ScheduleId uint `gorm:"column:schedule_id"`
	UserId     uint `gorm:"column:user_id"`
}
