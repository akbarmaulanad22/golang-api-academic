package entity

import "time"

type Student struct {
	Nim             int       `gorm:"column:nim;primaryKey"`
	Class           string    `gorm:"column:class"`
	RegistarionWave string    `gorm:"column:applicantWave"`
	RegistarionDate time.Time `gorm:"column:batch_year"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	// embedded data
	Biodata StudentBio `gorm:"embedded"`

	// foreign key
	UserId         int `gorm:"column:user_id"`
	StudyProgramId int `gorm:"column:study_program_id"`

	// relationship belongs to (one to one)
	User         *User         `gorm:"foreignKey:user_id;references:id"`
	StudyProgram *StudyProgram `gorm:"foreignKey:study_program_id;references:id"`
}

type StudentBio struct {
	Name      string    `gorm:"column:name"`
	DateBirth time.Time `gorm:"column:date_birth"`
	Address   string    `gorm:"column:address"`
	Gender    string    `gorm:"column:gender"`
}
