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

type CreateCourseRequest struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	Sks           int    `json:"sks"`
	Semester      int    `json:"semester"`
	TotalMeetings int    `json:"total_meetings"`
	LecturerNIDN  uint   `json:"lecturer_nidn"`
}

type UpdateCourseRequest struct {
	Code          string `json:"-"`
	Name          string `json:"name"`
	Sks           int    `json:"sks"`
	Semester      int    `json:"semester"`
	TotalMeetings int    `json:"total_meetings"`
	LecturerNIDN  uint   `json:"lecturer_nidn"`
}

type GetCourseRequest struct {
	Code string `json:"-" validate:"required"`
}

type DeleteCourseRequest struct {
	Code string `json:"-" validate:"required"`
}
