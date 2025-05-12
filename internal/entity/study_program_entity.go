package entity

type StudyProgram struct {
	Model
	Name             string `gorm:"column:name"`
	Level            string `gorm:"column:level"`
	Accreditation    string `gorm:"column:accreditation"`
	DurationSemester int    `gorm:"column:duration_semester"`

	// foreign key
	FacultyId int `gorm:"column:faculty_id"`

	// relationship has one (one to one)
	Student Student `gorm:"foreignKey:id;references:study_program_id"`
}
