package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func ClassroomToResponse(faculty *entity.Classroom) *model.ClassroomResponse {
	return &model.ClassroomResponse{
		ID:       faculty.ID,
		Name:     faculty.Name,
		Capacity: faculty.Capacity,
		Location: faculty.Location,
	}
}
