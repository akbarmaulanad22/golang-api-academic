package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func ScheduleToResponse(schedule entity.Schedule) model.ScheduleResponse {
	return model.ScheduleResponse{
		Course:    schedule.Course.Name,
		Lecturer:  schedule.Lecturer.Name,
		Classroom: schedule.Classroom.Name,
		StartAt:   schedule.StartAt,
		EndAt:     schedule.EndAt,
		Date:      schedule.Date,
	}
}

func ScheduleToResponses(schedules []entity.Schedule) []model.ScheduleResponse {
	scheduleResponses := []model.ScheduleResponse{}

	for _, schedule := range schedules {
		scheduleResponses = append(scheduleResponses, ScheduleToResponse(schedule))
	}

	return scheduleResponses

}
