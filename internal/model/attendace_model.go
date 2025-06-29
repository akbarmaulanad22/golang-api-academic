package model

import "time"

type AttendanceResponse struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

type AttendanceCreateRequest struct {
	UserId uint   `json:"user_id" validate:"required"`
	Status string `json:"status"`
}

type ListAttendanceRequest struct {
	Npm uint `json:"npm"`
}

type ListInLecturerAttendanceRequest struct {
	Npm        uint   `json:"npm"`
	CourseCode string `json:"course_code"`
}

type AttendanceUpdateRequest struct {
	ID     uint   `json:"-" validate:"required"`
	Status string `json:"status"`
}

type ListAttendanceStudentRequest struct {
	UserID uint `json:"-"`
}

type AttendanceGroupedResponse struct {
	Course      string               `json:"course"`
	Attendances []AttendanceResponse `json:"attendances"`
}
