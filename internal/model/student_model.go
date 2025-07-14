package model

import "time"

type ListStudentRequest struct {
	CourseCode string `json:"course_code"`
}

type StudentResponse struct {
	Npm  uint   `json:"npm"`
	Name string `json:"name"`
}

type StudentAdminResponse struct {
	Username string `json:"username" validate:"required"`

	Npm              uint      `json:"npm" validate:"required"`
	Class            string    `json:"class" validate:"required"`
	RegistrationWave string    `json:"registration_wave" validate:"required"`
	RegistrationDate time.Time `json:"registration_date" validate:"required"`

	Name      string    `json:"name" validate:"required"`
	DateBirth time.Time `json:"date_birth" validate:"required"`
	Address   string    `json:"address" validate:"required"`
	Gender    string    `json:"gender" validate:"required"`

	StudyProgramId uint `json:"study_program_id" validate:"required"`
}

type CreateStudentRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`

	Npm              uint      `json:"npm" validate:"required"`
	Class            string    `json:"class" validate:"required"`
	RegistrationWave string    `json:"registration_wave" validate:"required"`
	RegistrationDate time.Time `json:"registration_date" validate:"required"`

	Name      string    `json:"name" validate:"required"`
	DateBirth time.Time `json:"date_birth" validate:"required"`
	Address   string    `json:"address" validate:"required"`
	Gender    string    `json:"gender" validate:"required"`

	// UserId         uint `json:"user_id" validate:"required"`
	StudyProgramId uint `json:"study_program_id" validate:"required"`
}

type UpdateStudentRequest struct {
	Password string `json:"password"`

	Npm              uint      `json:"npm" validate:"required"`
	Class            string    `json:"class" validate:"required"`
	RegistrationWave string    `json:"registration_wave" validate:"required"`
	RegistrationDate time.Time `json:"registration_date" validate:"required"`

	Name      string    `json:"name" validate:"required"`
	DateBirth time.Time `json:"date_birth" validate:"required"`
	Address   string    `json:"address" validate:"required"`
	Gender    string    `json:"gender" validate:"required"`

	// UserId         uint `json:"user_id" validate:"required"`
	StudyProgramId uint `json:"study_program_id" validate:"required"`
}

type GetStudentRequest struct {
	Npm uint `json:"-" validate:"required"`
}

type DeleteStudentRequest struct {
	Npm uint `json:"-" validate:"required"`
}
