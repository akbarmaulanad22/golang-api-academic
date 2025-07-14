package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func ScheduleToResponse(schedule *entity.Schedule) *model.ScheduleResponse {
	return &model.ScheduleResponse{
		Course:    schedule.Course.Name,
		Lecturer:  schedule.Lecturer.Name,
		Classroom: schedule.Classroom.Name,
		StartAt:   schedule.StartAt,
		EndAt:     schedule.EndAt,
		Date:      schedule.Date,
	}
}

func ScheduleToStudentResponse(schedule *entity.Schedule, status string) *model.ScheduleStudentResponse {
	return &model.ScheduleStudentResponse{
		Course:         schedule.Course.Name,
		Lecturer:       schedule.Lecturer.Name,
		Classroom:      schedule.Classroom.Name,
		LecturerStatus: status,
	}
}

func ScheduleToAdminResponse(schedule *entity.Schedule) *model.ScheduleAdminResponse {
	return &model.ScheduleAdminResponse{
		StartAt:      schedule.StartAt,
		EndAt:        schedule.EndAt,
		Date:         schedule.Date,
		ID:           schedule.ID,
		CourseCode:   schedule.CourseCode,
		LecturerNIDN: schedule.LecturerNIDN,
		ClassroomID:  schedule.ClassroomID,
	}
}
