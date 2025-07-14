package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func EnrollmentToResponse(enrollment *entity.Enrollment) *model.EnrollmentResponse {
	return &model.EnrollmentResponse{
		Status:       enrollment.Status,
		AcademicYear: enrollment.AcademicYear,
		Name:         enrollment.Course.Name,
		Sks:          enrollment.Course.Sks,
		Semester:     enrollment.Course.Semester,
		Lecturer:     enrollment.Course.Lecturer.Name,
	}
}

func EnrollmentToAdminResponse(enrollment *entity.Enrollment) *model.EnrollmentAdminResponse {
	return &model.EnrollmentAdminResponse{
		ID:               enrollment.ID,
		Status:           enrollment.Status,
		AcademicYear:     enrollment.AcademicYear,
		RegistrationDate: enrollment.RegistrationDate,
		StudentNpm:       enrollment.StudentNpm,
		CourseCode:       enrollment.CourseCode,
	}
}
