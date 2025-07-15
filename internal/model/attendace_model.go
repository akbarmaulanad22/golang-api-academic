package model

import "time"

type AttendanceResponse struct {
	ID     uint      `json:"id"`
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

type AttendanceCreateLecturerRequest struct {
	Status     string `json:"status" validate:"required"`
	ScheduleID uint   `json:"schedule_id" validate:"required"`
	Npm        uint   `json:"-" validate:"required"`
}

type AttendanceCreateRequest struct {
	UserId uint   `json:"-" validate:"required"`
	Status string `json:"status" validate:"required"`
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

type ListAvailableScheduleAttendanceStudentRequest struct {
	Npm        uint   `json:"-"`
	CourseCode string `json:"-"`
}

type AttendanceGroupedResponse struct {
	Course      string               `json:"course"`
	Attendances []AttendanceResponse `json:"attendances"`
}

type DeleteAttendanceRequest struct {
	ID uint `json:"-" validate:"required"`
}
