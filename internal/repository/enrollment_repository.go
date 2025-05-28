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

// Cek apakah mahasiswa terdaftar di jadwal tertentu
func (r *EnrollmentRepository) IsStudentEnrolled(db *gorm.DB, studentNPM string, scheduleID uint) bool {
	var count int64
	db.Model(&entity.Enrollment{}).
		Where("student_npm = ? AND schedule_id = ?", studentNPM, scheduleID).
		Count(&count)
	return count > 0
}

// Ambil semua enrollment per mahasiswa
func (r *EnrollmentRepository) GetByStudentNPM(db *gorm.DB, npm string) ([]entity.Enrollment, error) {
	var enrollments []entity.Enrollment
	err := db.Where("student_npm = ?", npm).Find(&enrollments).Error
	return enrollments, err
}
