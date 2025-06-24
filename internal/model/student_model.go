package model

type ListStudentRequest struct {
	CourseCode string `json:"course_code"`
}

type StudentResponse struct {
	Npm  uint   `json:"npm"`
	Name string `json:"name"`
}
