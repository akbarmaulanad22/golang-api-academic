package model

import "time"

type ListScheduleRequest struct {
	UserID uint `json:"user_id"`
}

type ScheduleResponse struct {
	Course    string    `json:"course"`
	Lecturer  string    `json:"lecturer"`
	Classroom string    `json:"classroom"`
	StartAt   string    `json:"start_at"`
	EndAt     string    `json:"end_at"`
	Date      time.Time `json:"date"`
}

type ScheduleStudentResponse struct {
	Course         string `json:"course"`
	Lecturer       string `json:"lecturer"`
	Classroom      string `json:"classroom"`
	LecturerStatus string `json:"lecturer_status"`
}

type ScheduleAdminResponse struct {
	ID           uint      `json:"id"`
	Date         time.Time `json:"date"`
	StartAt      string    `json:"start_at"`
	EndAt        string    `json:"end_at"`
	CourseCode   string    `json:"course_code"`
	LecturerNIDN uint      `json:"lecturer_nidn"`
	ClassroomID  uint      `json:"classroom_id"`
}

type CreateScheduleRequest struct {
	Date         time.Time `json:"date" validate:"required"`
	StartAt      string    `json:"start_at" validate:"required"`
	EndAt        string    `json:"end_at" validate:"required"`
	CourseCode   string    `json:"course_code" validate:"required"`
	LecturerNIDN uint      `json:"lecturer_nidn" validate:"required"`
	ClassroomID  uint      `json:"classroom_id" validate:"required"`
}

type UpdateScheduleRequest struct {
	ID           uint      `json:"-"`
	Date         time.Time `json:"date" validate:"required"`
	StartAt      string    `json:"start_at" validate:"required"`
	EndAt        string    `json:"end_at" validate:"required"`
	CourseCode   string    `json:"course_code" validate:"required"`
	LecturerNIDN uint      `json:"lecturer_nidn" validate:"required"`
	ClassroomID  uint      `json:"classroom_id" validate:"required"`
}

type GetScheduleRequest struct {
	ID uint `json:"-" validate:"required"`
}

type DeleteScheduleRequest struct {
	ID uint `json:"-" validate:"required"`
}
