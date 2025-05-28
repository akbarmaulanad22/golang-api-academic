package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	Repository[entity.Schedule]
	Log *logrus.Logger
}

func NewScheduleRepository(log *logrus.Logger) *ScheduleRepository {

	return &ScheduleRepository{Log: log}

}

// Cek apakah jadwal aktif hari ini (30 menit sebelum start_at)
func (r *ScheduleRepository) GetTodayScheduleIDByUserID(db *gorm.DB, userID uint) (int64, error) {
	var result int64

	err := db.Raw(`
        SELECT s.id AS schedule_id
        FROM students st
        JOIN enrollments e ON st.npm = e.student_npm
        JOIN schedules s ON e.schedule_id = s.id
        WHERE st.user_id = ?
            AND s.date = CURDATE()
            AND NOW() BETWEEN DATE_SUB(CONCAT(s.date, ' ', s.start_at), INTERVAL 30 MINUTE) 
            AND CONCAT(s.date, ' ', s.end_at)
        ORDER BY s.start_at
        LIMIT 1;
    `, userID).Scan(&result).Error

	return result, err

}
