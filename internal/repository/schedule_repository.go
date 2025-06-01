package repository

import (
	"fmt"
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

// GetActiveScheduleID - ambil schedule_id berdasarkan user_id
func (r ScheduleRepository) GetActiveScheduleIDByUserID(db *gorm.DB, userID uint) (uint, error) {

	var scheduleID uint
	err := db.Raw(`
		SELECT s.id 
        FROM students st
        JOIN enrollments e ON st.npm = e.student_npm
        JOIN schedules s ON e.schedule_id = s.id
        WHERE st.user_id = ?
          AND s.date = CURDATE()
          AND NOW() BETWEEN DATE_SUB(CONCAT(s.date, ' ', s.start_at), INTERVAL 30 MINUTE) 
                       AND CONCAT(s.date, ' ', s.end_at)
        ORDER BY s.start_at
        LIMIT 1;
	`, userID).Scan(&scheduleID).Error

	if err != nil {
		return 0, err
	}

	if scheduleID == 0 {
		return 0, fmt.Errorf("schedule not found")
	}

	return scheduleID, nil
}
