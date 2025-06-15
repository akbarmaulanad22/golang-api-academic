package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CourseRepository struct {
	Repository[entity.Course]
	Log *logrus.Logger
}

func NewCourseRepository(log *logrus.Logger) *CourseRepository {

	return &CourseRepository{Log: log}

}

func (r *CourseRepository) GetCoursesByNPM(db *gorm.DB, npm uint) ([]entity.CourseWithMeetings, error) {
	var courses []entity.CourseWithMeetings

	err := db.Raw(`
        SELECT 
            c.code AS code,
            c.name AS name,
            COUNT(s.id) AS total_meetings
        FROM enrollments e
        JOIN schedules s ON e.schedule_id = s.id
        JOIN courses c ON s.course_code = c.code
        WHERE e.student_npm = ?
        GROUP BY c.code;
    `, npm).Scan(&courses).Error

	return courses, err
}
