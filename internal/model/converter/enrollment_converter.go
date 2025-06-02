package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func EnrollmentToResponse(enrollment entity.Enrollment) model.EnrollmentResponse {
	return model.EnrollmentResponse{
		Status:       enrollment.Status,
		AcademicYear: enrollment.AcademicYear,
		Name:         enrollment.Course.Name,
		Sks:          enrollment.Course.Sks,
		Semester:     enrollment.Course.Semester,
	}
}

func EnrollmentToResponses(enrollments []entity.Enrollment) []model.EnrollmentResponse {
	enrollmentResponses := []model.EnrollmentResponse{}

	for _, enrollment := range enrollments {
		enrollmentResponses = append(enrollmentResponses, EnrollmentToResponse(enrollment))
	}

	return enrollmentResponses

}
