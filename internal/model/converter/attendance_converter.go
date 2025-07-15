package converter

import (
	"time"
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func AttendanceToResponse(attendance *entity.Attendance) *model.AttendanceResponse {
	return &model.AttendanceResponse{
		ID:     attendance.ID,
		Status: attendance.Status,
		Time:   attendance.Time,
	}
}

func AttendanceGroupedToResponse(attendance map[string]any) model.AttendanceGroupedResponse {
	return model.AttendanceGroupedResponse{
		// Course:      attendance["sss"],
		Attendances: []model.AttendanceResponse{
			{
				Status: attendance["attendance_status"].(string),
				Time:   attendance["attendance_time"].(time.Time),
			},
		},
		Course: attendance["course_name"].(string),
	}
}
