package model

import "time"

type AttendanceResponse struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

type AttendanceCreateResponse struct {
	UserId uint `json:"user_id" validate:"required"`
}

type ListAttendanceRequest struct {
	Npm uint `json:"npm"`
}

type ListInLecturerAttendanceRequest struct {
	Npm        uint   `json:"npm"`
	CourseCode string `json:"course_code"`
}
