package entity

import "time"

type Enrollment struct {
	Model
	Status           string `gorm:"column:status"`
	AcademicYear     string `gorm:"column:academic_year"`
	DateRegistration time.Time

	StudentId  int `gorm:"column:student_id"`
	ScheduleId int `gorm:"column:schedule_id"`
}
