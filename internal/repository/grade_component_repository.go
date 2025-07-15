package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GradeComponentRepository struct {
	Repository[entity.GradeComponent]
	Log *logrus.Logger
}

func NewGradeComponentRepository(log *logrus.Logger) *GradeComponentRepository {

	return &GradeComponentRepository{Log: log}

}

func (r *GradeComponentRepository) FindAll(db *gorm.DB) ([]entity.GradeComponent, error) {

	var gradeComponent []entity.GradeComponent
	if err := db.Find(&gradeComponent).Error; err != nil {
		return nil, err
	}

	return gradeComponent, nil
}

func (r *GradeComponentRepository) FindAvailableByCourseCodeAndNpm(db *gorm.DB, courseCode string, npm uint) ([]entity.GradeComponent, error) {

	var gradeComponent []entity.GradeComponent
	if err := db.Where(`
		NOT EXISTS (
			SELECT 1
			FROM grades g
			JOIN enrollments e ON g.enrollment_id = e.id
			WHERE g.grade_component_id = grade_components.id
			AND e.course_code = ?
			AND e.student_npm = ?
		)
	`, courseCode, npm).
		Find(&gradeComponent).Error; err != nil {
		return nil, err
	}

	return gradeComponent, nil
}
