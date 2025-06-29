package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func StudentToResponse(student *entity.Student) *model.StudentResponse {
	return &model.StudentResponse{
		Npm:  student.Npm,
		Name: student.Biodata.Name,
	}
}

func StudentAdminToResponse(student *entity.Student) *model.StudentAdminResponse {
	return &model.StudentAdminResponse{
		Username:         student.User.Username,
		Npm:              student.Npm,
		Class:            student.Class,
		RegistrationWave: student.RegistrationWave,
		RegistrationDate: student.RegistrationDate,
		Name:             student.Biodata.Name,
		DateBirth:        student.Biodata.DateBirth,
		Address:          student.Biodata.Address,
		Gender:           student.Biodata.Gender,
		StudyProgramId:   student.StudyProgramId,
	}
}
