package entity

import "time"

type Attendance struct {
	Model
	Time time.Time `gorm:"column:date"`

	// foreign key
	ScheduleId   int `gorm:"column:schedule_id"`
	EnrollmentId int `gorm:"column:enrollment_id"`
	UserId       int `gorm:"column:user_id"`
}
