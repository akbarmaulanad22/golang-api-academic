package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func FacultyToResponse(faculty *entity.Faculty) *model.FacultyResponse {
	return &model.FacultyResponse{
		ID:      faculty.ID,
		Name:    faculty.Name,
		Code:    faculty.Code,
		Dekan:   faculty.Dekan,
		Address: faculty.Address,
	}
}
