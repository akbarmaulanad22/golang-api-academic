package repository

import (
	"fmt"
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ScheduleRepository struct {
	Repository[entity.Schedule]
	Log *logrus.Logger
}

func NewScheduleRepository(log *logrus.Logger) *ScheduleRepository {

	return &ScheduleRepository{Log: log}

}

// GetActiveScheduleID - ambil schedule_id 30 menit kedepan berdasarkan user_id
func (r *ScheduleRepository) GetStudentActiveScheduleIDByUserID(db *gorm.DB, userID uint) (uint, error) {

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
func (r *ScheduleRepository) FindAllScheduleByStudentUserID(db *gorm.DB, userID uint) ([]entity.Schedule, error) {

	var schedules []entity.Schedule

	err := db.Model(&entity.Schedule{}).
		Joins("JOIN enrollments ON enrollments.schedule_id = schedules.id").
		Joins("JOIN students ON students.npm = enrollments.student_npm").
		Preload("Lecturer").
		Preload("Course").
		Preload("Classroom").
		Where("students.user_id = ? AND schedules.date >= CURDATE()", userID).
		Order(clause.OrderBy{Columns: []clause.OrderByColumn{
			{Column: clause.Column{Name: "schedules.date"}, Desc: true},
			{Column: clause.Column{Name: "schedules.start_at"}, Desc: false},
		}}).
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

func (r ScheduleRepository) GetLecturerActiveScheduleByUserID(db *gorm.DB, userID uint) (uint, error) {
	var scheduleID uint

	err := db.Raw(`
        SELECT s.id 
        FROM schedules s
        JOIN lecturers l ON s.lecturer_nidn = l.nidn
        WHERE l.user_id = ?
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

func (r ScheduleRepository) FindAllSchedulesByLecturerUserID(db *gorm.DB, userID uint) ([]entity.Schedule, error) {
	var schedules []entity.Schedule

	err := db.Model(&entity.Schedule{}).
		Joins("JOIN lecturers ON lecturers.nidn = schedules.lecturer_nidn").
		Preload("Course").
		Preload("Classroom").
		Where("lecturers.user_id = ? AND schedules.date >= CURDATE()", userID).
		Order(clause.OrderBy{Columns: []clause.OrderByColumn{
			{Column: clause.Column{Name: "schedules.date"}, Desc: true},
			{Column: clause.Column{Name: "schedules.start_at"}, Desc: false},
		}}).
		Find(&schedules).Error

	return schedules, err
}
