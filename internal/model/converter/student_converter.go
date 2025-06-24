package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func StudentToResponse(student entity.Student) model.StudentResponse {
	return model.StudentResponse{
		Npm:  student.Npm,
		Name: student.Biodata.Name,
	}
}

func StudentToResponses(students []entity.Student) []model.StudentResponse {
	studentResponses := []model.StudentResponse{}

	for _, student := range students {
		studentResponses = append(studentResponses, StudentToResponse(student))
	}

	return studentResponses

}
