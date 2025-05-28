package entity

type StudyProgram struct {
	Entity
	Name             string `gorm:"column:name"`
	Level            string `gorm:"column:level"`
	Accreditation    string `gorm:"column:accreditation"`
	DurationSemester int    `gorm:"column:duration_semester"`

	// foreign key
	FacultyId int `gorm:"column:faculty_id"`

	// relationship
	Faculty Faculty `gorm:"foreignKey:faculty_id;references:id"`
}
