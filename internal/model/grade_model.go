package model

type ListGradeRequest struct {
	UserID uint `json:"user_id"`
}

type GradeResponse struct {
	CourseName string                   `json:"course_name"`
	TotalScore float64                  `json:"total_score"`
	Components []GradeComponentResponse `json:"components"`
}

type GradeInLecturerResponse struct {
	Type  string  `json:"type"`
	Score float64 `json:"score"`
}

type ListInLecturerGradeRequest struct {
	Npm        uint   `json:"npm"`
	CourseCode string `json:"course_code"`
}
