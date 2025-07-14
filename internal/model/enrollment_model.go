package model

import "time"

type EnrollmentResponse struct {
	Status       string `json:"status"`
	AcademicYear string `json:"academic_year"`
	Name         string `json:"name"`
	Sks          int    `json:"sks"`
	Semester     int    `json:"semester"`
	Lecturer     string `json:"lecturer"`
}

type EnrollmentAdminResponse struct {
	ID               uint      `json:"id"`
	Status           string    `json:"status"`
	AcademicYear     string    `json:"academic_year"`
	RegistrationDate time.Time `json:"registration_date"`
	StudentNpm       uint      `json:"student_npm"`
	CourseCode       string    `json:"course_code"`
}

type ListEnrollmentRequest struct {
	UserID uint `json:"user_id"`
}

type CreateEnrollmentRequest struct {
	Status           string    `json:"status"`
	AcademicYear     string    `json:"academic_year"`
	RegistrationDate time.Time `json:"registration_date"`
	StudentNpm       uint      `json:"student_npm"`
	CourseCode       string    `json:"course_code"`
}

type UpdateEnrollmentRequest struct {
	ID               uint      `json:"-"`
	Status           string    `json:"status"`
	AcademicYear     string    `json:"academic_year"`
	RegistrationDate time.Time `json:"registration_date"`
	StudentNpm       uint      `json:"student_npm"`
	CourseCode       string    `json:"course_code"`
}

type GetEnrollmentRequest struct {
	ID uint `json:"-" validate:"required"`
}

type DeleteEnrollmentRequest struct {
	ID uint `json:"-" validate:"required"`
}
