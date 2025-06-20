package model

type ListGradeRequest struct {
	UserID uint `json:"user_id"`
}

type GradeResponse struct {
	CourseName string                        `json:"course_name"`
	TotalScore float64                       `json:"total_score"`
	Components []GradeComponentScoreResponse `json:"components"`
}

type GradeComponentScoreResponse struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}
