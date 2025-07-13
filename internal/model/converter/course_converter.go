package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func CourseToResponse(course *entity.Course) *model.CourseResponse {
	return &model.CourseResponse{
		Code:         course.Code,
		Name:         course.Name,
		Sks:          course.Sks,
		Semester:     course.Semester,
		LecturerNIDN: course.LecturerNIDN,
		// TotalMeetings: course.TotalMeetings,
	}
}
