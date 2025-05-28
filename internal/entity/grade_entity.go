package entity

type Grade struct {
	Entity
	Score int `gorm:"column:score"`

	// foreign key
	EnrollmentId     int `gorm:"column:enrollment_id"`
	GradeComponentId int `gorm:"column:grade_component_id"`

	// relationship
	Enrollment     Enrollment     `gorm:"foreignKey:enrollment_id;references:id"`
	GradeComponent GradeComponent `gorm:"foreignKey:grade_component_id;references:id"`
}
