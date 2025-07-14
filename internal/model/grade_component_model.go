package model

type GradeComponentResponse struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
	Grade float64 `json:"grade"`
}

type GradeComponentAdminResponse struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

type CreateGradeComponentRequest struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

type UpdateGradeComponentRequest struct {
	ID     uint    `json:"-"`
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

type GetGradeComponentRequest struct {
	ID uint `json:"-" validate:"required"`
}

type DeleteGradeComponentRequest struct {
	ID uint `json:"-" validate:"required"`
}
