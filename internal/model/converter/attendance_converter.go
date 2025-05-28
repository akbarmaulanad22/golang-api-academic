package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func AttendanceToResponse(attendance *entity.Attendance) *model.AttendanceResponse {
	return &model.AttendanceResponse{
		Status: attendance.Status,
		Time:   attendance.Time,
	}
}
