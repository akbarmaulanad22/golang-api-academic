package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AttendanceRepository struct {
	Repository[entity.Attendance]
	Log *logrus.Logger
}

func NewAttendanceRepository(log *logrus.Logger) *AttendanceRepository {

	return &AttendanceRepository{Log: log}

}

// Cek apakah dosen sudah absen di jadwal ini
func (r *AttendanceRepository) HasLecturer(db *gorm.DB, scheduleID uint) bool {
	var count int64
	db.Raw(`
        SELECT COUNT(*) FROM attendance a
        JOIN lecturers l ON a.user_id = l.user_id
        WHERE a.schedule_id = ?
          AND DATE(a.time) = CURDATE()
    `, scheduleID).Count(&count)
	return count > 0
}

// Cek apakah mahasiswa sudah absen
func (r *AttendanceRepository) HasAlreadyAttended(db *gorm.DB, userID uint, scheduleID uint) bool {
	var count int64
	db.Where("user_id = ? AND schedule_id = ?", userID, scheduleID).Model(&entity.Attendance{}).Count(&count)
	return count > 0
}
