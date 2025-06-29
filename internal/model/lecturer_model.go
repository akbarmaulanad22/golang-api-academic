package model

type LecturerResponse struct {
	Username string `json:"username" validate:"required"`

	NIDN       uint   `json:"nidn"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Degree     string `json:"degree"`
	IsFullTime bool   `json:"is_full_time"`
}

type CreateLecturerRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`

	NIDN       uint   `json:"nidn" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Gender     string `json:"gender" validate:"required"`
	Degree     string `json:"degree" validate:"required"`
	IsFullTime bool   `json:"is_full_time"`
}

type UpdateLecturerRequest struct {
	Password string `json:"password" validate:"required"`

	NIDN       uint   `json:"-"`
	Name       string `json:"name" validate:"required"`
	Gender     string `json:"gender" validate:"required"`
	Degree     string `json:"degree" validate:"required"`
	IsFullTime bool   `json:"is_full_time" validate:"required"`
}

type GetLecturerRequest struct {
	NIDN uint `json:"-" validate:"required"`
}

type DeleteLecturerRequest struct {
	NIDN uint `json:"-" validate:"required"`
}
