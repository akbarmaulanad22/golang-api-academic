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

// IsLecturerPresent - cek apakah dosen sudah hadir di jadwal ini
func (r *AttendanceRepository) IsLecturerPresent(db *gorm.DB, scheduleID uint) bool {
	var count int64
	db.Raw(`
        SELECT COUNT(*) 
        FROM schedules s
        JOIN lecturers l ON s.lecturer_nidn = l.nidn
        LEFT JOIN attendance a ON a.schedule_id = s.id AND DATE(a.time) = CURDATE() AND a.user_id = l.user_id
        WHERE s.id = ?
        AND a.status IS NOT NULL
    `, scheduleID).Count(&count)

	return count > 0
}

// HasAlreadyAttended - cek apakah mahasiswa sudah absen
func (r *AttendanceRepository) HasAlreadyAttended(db *gorm.DB, userID uint, scheduleID uint) bool {
	var count int64
	db.Raw(`
        SELECT COUNT(*) FROM attendance a
        JOIN students st ON a.user_id = st.user_id
        WHERE a.schedule_id = ? AND st.user_id = ? AND DATE(a.time) = CURDATE()
    `, scheduleID, userID).Count(&count)

	return count > 0
}

func (r *AttendanceRepository) Create(db *gorm.DB, entity *entity.Attendance) error {

	return db.Create(entity).Error
}
