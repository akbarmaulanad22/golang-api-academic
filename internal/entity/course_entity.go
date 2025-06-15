package entity

import "time"

type Course struct {
	Code          string    `gorm:"column:code;primaryKey"`
	Name          string    `gorm:"column:name"`
	Sks           int       `gorm:"column:sks"`
	Semester      int       `gorm:"column:semester"`
	TotalMeetings int       `gorm:"column:total_meetings"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	// foreign key
	LecturerNIDN uint `gorm:"column:lecturer_nidn"`

	// relationship
	StudyProgram []StudyProgram `gorm:"many2many:study_program_course"`
}
