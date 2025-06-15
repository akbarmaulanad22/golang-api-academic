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

func (r *CourseRepository) FindAllByNIDNUserID(db *gorm.DB, userID uint) ([]entity.Course, error) {
	var courses []entity.Course
	if err := db.
		Joins("JOIN lecturers ON courses.lecturer_nidn = lecturers.nidn").
		Where("lecturers.user_id = ?", userID).
		Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}
