package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func GradeInLecturerToResponse(grade *entity.Grade) *model.GradeInLecturerResponse {
	return &model.GradeInLecturerResponse{
		ID:    grade.ID,
		Type:  grade.GradeComponent.Name,
		Score: grade.Score,
	}
}
