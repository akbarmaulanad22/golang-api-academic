package entity

import "time"

type Enrollment struct {
	Entity
	Status           string    `gorm:"column:status"`
	AcademicYear     string    `gorm:"column:academic_year"`
	RegistrationDate time.Time `gorm:"column:registration_date"`

	// foreign key
	StudentNpm string `gorm:"column:student_npm"`
	CourseCode string `gorm:"column:course_code"`

	// relationship
	Student Student `gorm:"foreignKey:student_npm;references:npm"`
	Course  Course  `gorm:"foreignKey:course_code;references:code"`

	Grade []Grade `gorm:"foreignKey:enrollment_id;references:id"`
}
