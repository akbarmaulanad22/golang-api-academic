package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func StudyProgramToResponse(attendance *entity.StudyProgram) *model.StudyProgramResponse {
	return &model.StudyProgramResponse{
		ID:               attendance.ID,
		Name:             attendance.Name,
		Level:            attendance.Level,
		Accreditation:    attendance.Accreditation,
		DurationSemester: attendance.DurationSemester,
	}
}
