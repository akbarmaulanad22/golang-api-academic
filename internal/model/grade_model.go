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
	ID    uint    `json:"id"`
	Type  string  `json:"type"`
	Score float64 `json:"score"`
}

type ListInLecturerGradeRequest struct {
	Npm        uint   `json:"npm"`
	CourseCode string `json:"course_code"`
}

type CreateGradeRequest struct {
	Npm              uint    `json:"-" validate:"required"`
	CourseCode       string  `json:"-" validate:"required"`
	Score            float64 `json:"score" validate:"required"`
	GradeComponentID uint    `json:"grade_component_id" validate:"required"`
}

type UpdateGradeRequest struct {
	ID    uint    `json:"-"`
	Score float64 `json:"score" validate:"required"`
}

type GetGradeRequest struct {
	ID uint `json:"-" validate:"required"`
}

type DeleteGradeRequest struct {
	ID uint `json:"-" validate:"required"`
}
