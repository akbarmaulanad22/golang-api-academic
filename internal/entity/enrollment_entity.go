package entity

import "time"

type Enrollment struct {
	Entity
	Status           string `gorm:"column:status"`
	AcademicYear     string `gorm:"column:academic_year"`
	DateRegistration time.Time

	// foreign key
	StudentNpm int `gorm:"column:student_npm"`
	ScheduleId int `gorm:"column:schedule_id"`

	// relationship
	Student  Student  `gorm:"foreignKey:student_npm;references:npm"`
	Schedule Schedule `gorm:"foreignKey:schedule_id;references:id"`
}
