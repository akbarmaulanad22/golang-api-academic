package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func LecturerToResponse(lecturer *entity.Lecturer) *model.LecturerResponse {
	return &model.LecturerResponse{
		Username:   lecturer.User.Username,
		NIDN:       lecturer.Nidn,
		Name:       lecturer.Name,
		Gender:     lecturer.Gender,
		Degree:     lecturer.Degree,
		IsFullTime: lecturer.IsFullTime,
	}
}

// func LecturerToResponse(lecturer *entity.Lecturer) *model.LecturerResponse {
// 	return &model.LecturerResponse{
// 		NIDN:       lecturer.Nidn,
// 		Name:       lecturer.Name,
// 		Gender:     lecturer.Gender,
// 		Degree:     lecturer.Degree,
// 		IsFullTime: lecturer.IsFullTime,
// 	}
// }
