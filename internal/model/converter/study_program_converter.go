package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func StudyProgramToResponse(studyProgram *entity.StudyProgram) *model.StudyProgramResponse {
	return &model.StudyProgramResponse{
		ID:               studyProgram.ID,
		Name:             studyProgram.Name,
		Level:            studyProgram.Level,
		Accreditation:    studyProgram.Accreditation,
		DurationSemester: studyProgram.DurationSemester,
		FacultyID:        studyProgram.FacultyId,
	}
}
