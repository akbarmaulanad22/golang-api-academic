package entity

type Grade struct {
	Model
	Score int `gorm:"column:score"`

	// foreign key
	EnrollmentId     int `gorm:"column:enrollment_id"`
	GradeComponentId int `gorm:"column:grade_component_id"`
}
