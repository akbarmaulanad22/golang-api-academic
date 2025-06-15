package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func CourseToResponse(course entity.Course) model.CourseResponse {
	return model.CourseResponse{
		Code:          course.Code,
		Name:          course.Name,
		Sks:           course.Sks,
		Semester:      course.Semester,
		TotalMeetings: course.TotalMeetings,
	}
}

func CourseToResponses(courses []entity.Course) []model.CourseResponse {
	courseResponses := []model.CourseResponse{}

	for _, course := range courses {
		courseResponses = append(courseResponses, CourseToResponse(course))
	}

	return courseResponses

}
