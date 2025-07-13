package model

type ClassroomResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Location string `json:"location"`
}

type CreateClassroomRequest struct {
	Name     string `json:"name" validate:"required"`
	Capacity int    `json:"capacity" validate:"required"`
	Location string `json:"location" validate:"required"`
}

type UpdateClassroomRequest struct {
	ID       uint   `json:"-"`
	Name     string `json:"name" validate:"required"`
	Capacity int    `json:"capacity" validate:"required"`
	Location string `json:"location" validate:"required"`
}

type GetClassroomRequest struct {
	ID uint `json:"-" validate:"required"`
}

type DeleteClassroomRequest struct {
	ID uint `json:"-" validate:"required"`
}
