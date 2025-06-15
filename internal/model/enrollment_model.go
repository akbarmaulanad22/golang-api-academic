package model

type EnrollmentResponse struct {
	Status       string `json:"status"`
	AcademicYear string `json:"academic_year"`
	Name         string `json:"name"`
	Sks          int    `json:"sks"`
	Semester     int    `json:"semester"`
}

type ListEnrollmentRequest struct {
	UserID uint `json:"user_id"`
}
