package entity

import "time"

type Schedule struct {
	Entity
	Date    time.Time `gorm:"column:date"`
	StartAt string    `gorm:"column:start_at"`
	EndAt   string    `gorm:"column:end_at"`

	// Foreign Keys
	CourseCode   string `gorm:"column:course_code"`
	LecturerNIDN uint   `gorm:"column:lecturer_nidn"`
	ClassroomID  uint   `gorm:"column:classroom_id"`

	// Relationships
	Course    Course    `gorm:"foreignKey:CourseCode;references:Code"`
	Lecturer  Lecturer  `gorm:"foreignKey:LecturerNIDN;references:Nidn"`
	Classroom Classroom `gorm:"foreignKey:ClassroomID;references:ID"`
}
