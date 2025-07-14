package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func GradeComponentToResponse(gradeComponent *entity.GradeComponent) *model.GradeComponentAdminResponse {
	return &model.GradeComponentAdminResponse{
		ID:     gradeComponent.ID,
		Name:   gradeComponent.Name,
		Weight: gradeComponent.Weight,
	}
}
