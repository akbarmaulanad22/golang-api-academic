package model

type ListCourseRequest struct {
	UserID uint `json:"user_id"`
}

type CourseResponse struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	Sks           int    `json:"sks"`
	Semester      int    `json:"semester"`
	TotalMeetings int    `json:"total_meetings"`
}
