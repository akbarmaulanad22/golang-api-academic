package entity

import "time"

type Schedule struct {
	Entity
	Date    time.Time `gorm:"column:date"`
	StartAt string    `gorm:"column:start_at"`
	EndAt   string    `gorm:"column:end_at"`

	// foreign key
	CourseCode   int `gorm:"column:course_code"`
	LecturerNIDN int `gorm:"column:lecturer_nidn"`
	ClassroomId  int `gorm:"column:classroom_id"`

	// relationship
	Course    Course    `gorm:"foreignKey:course_code;references:code"`
	Lecturer  Lecturer  `gorm:"foreignKey:lecturer_nidn:references:nidn"`
	Classroom Classroom `gorm:"foreignKey:classroom_id;references:id"`
}
