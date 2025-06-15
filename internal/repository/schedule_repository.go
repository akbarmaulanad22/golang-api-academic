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

// GetActiveScheduleID - ambil schedule_id 30 menit kedepan berdasarkan user_id
func (r *ScheduleRepository) GetActiveScheduleIDByUserID(db *gorm.DB, userID uint) (uint, error) {

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

// GetActiveScheduleID - ambil schedule_id berdasarkan user_id
func (r *ScheduleRepository) GetScheduleTodayByStudentUserID(db *gorm.DB, userID uint) ([]entity.Schedule, error) {

	// var schedules []entity.Schedule

	// SELECT c.name AS course, l.name AS lecturer, cr.name AS classroom, s.start_at, s.end_at

	// err := db.Raw(`
	// 	SELECT c.name, l.name, cr.name, s.start_at, s.end_at
	// 	FROM students st
	// 	JOIN enrollments e ON st.npm = e.student_npm
	// 	JOIN schedules s ON e.schedule_id = s.id
	// 	JOIN courses c ON s.course_code = c.code
	// 	JOIN lecturers l ON s.lecturer_nidn = l.nidn
	// 	JOIN classrooms cr ON s.classroom_id = cr.id
	// 	WHERE st.user_id = 3
	// 		AND s.date = CURDATE()
	// 	ORDER BY s.start_at
	// `, userID).Scan(&schedules).Error

	var schedules []entity.Schedule

	err := db.Model(&entity.Schedule{}).
		Joins("JOIN enrollments ON enrollments.schedule_id = schedules.id").
		Joins("JOIN students ON students.npm = enrollments.student_npm").
		Preload("Lecturer").
		Preload("Course").
		Preload("Classroom").
		Where("students.user_id = ? AND schedules.date = CURDATE()", userID).
		Order("schedules.start_at").
		Find(&schedules).Error

	if err != nil {
		return nil, err
	}

	return schedules, nil
}

// Hitung jumlah pertemuan per course
func (r *ScheduleRepository) GetTotalMeetingsByCourseCode(db *gorm.DB, courseCode string) (int64, error) {
	var total int64
	err := db.Table("schedules").
		Where("course_code = ?", courseCode).
		Count(&total).Error
	return total, err
}
