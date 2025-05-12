package entity

import "time"

type Schedule struct {
	Model
	Date    time.Time `gorm:"column:date"`
	StartAt time.Time `gorm:"column:start_at"`
	EndAt   time.Time `gorm:"column:end_at"`

	// foreign key
	CourseId    int `gorm:"column:course_id"`
	LecturerId  int `gorm:"column:lecturer_id"`
	ClassroomId int `gorm:"column:classroom_id"`
}
