package model

type FacultyResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Dekan   string `json:"dekan"`
	Address string `json:"address"`
}

type CreateFacultyRequest struct {
	Name    string `json:"name" validate:"required"`
	Code    string `json:"code" validate:"required"`
	Dekan   string `json:"dekan" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type UpdateFacultyRequest struct {
	ID      uint   `json:"-"`
	Name    string `json:"name" validate:"required"`
	Code    string `json:"code" validate:"required"`
	Dekan   string `json:"dekan" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type GetFacultyRequest struct {
	ID uint `json:"-" validate:"required"`
}

type DeleteFacultyRequest struct {
	ID uint `json:"-" validate:"required"`
}
