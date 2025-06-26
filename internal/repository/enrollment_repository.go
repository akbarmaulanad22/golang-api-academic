package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EnrollmentRepository struct {
	Repository[entity.Enrollment]
	Log *logrus.Logger
}

func NewEnrollmentRepository(log *logrus.Logger) *EnrollmentRepository {

	return &EnrollmentRepository{Log: log}

}

// GetActiveEnrollmentID - ambil enrollment berdasarkan user_id
func (r *EnrollmentRepository) FindAllEnrollmentByStudentUserID(db *gorm.DB, userID uint) ([]entity.Enrollment, error) {

	var enrollments []entity.Enrollment

	err := db.
		Joins("JOIN students ON enrollments.student_npm = students.npm").
		Preload("Course").
		Where("students.user_id = ?", userID).
		Order("enrollments.id DESC").
		Find(&enrollments).Error

	if err != nil {
		return nil, err
	}

	return enrollments, nil
}

func (r *EnrollmentRepository) FindEnrollmentByNpmAndCourseCode(db *gorm.DB, enrollment *entity.Enrollment, npm uint, courseCode string) error {
	return db.Where("student_npm = ? AND course_code = ? AND YEAR(registration_date) = YEAR(CURDATE())", npm, courseCode).Take(enrollment).Error
}
