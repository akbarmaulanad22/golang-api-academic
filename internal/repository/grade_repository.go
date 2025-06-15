package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GradeRepository struct {
	Repository[entity.Grade]
	Log *logrus.Logger
}

func NewGradeRepository(log *logrus.Logger) *GradeRepository {

	return &GradeRepository{Log: log}

}

// Cek apakah nilai absensi sudah ada
func (r *GradeRepository) AttendanceGradeAlreadyExists(db *gorm.DB, enrollmentID uint) bool {
	var count int64
	db.Model(&entity.Grade{}).
		Where("enrollment_id = ? AND grade_component_id = ?", enrollmentID, 1).
		Count(&count)
	return count > 0
}

// Simpan nilai absensi jika belum ada
func (r *GradeRepository) SaveAttendanceGrade(db *gorm.DB, enrollmentID uint, score float64) error {
	newGrade := entity.Grade{
		EnrollmentId:     enrollmentID,
		GradeComponentId: 1, // ID komponen 'Absensi'
		Score:            score,
	}
	return db.Create(&newGrade).Error
}

// Ambil semua nilai per enrollment
func (r *GradeRepository) GetAllGradesByEnrollmentID(db *gorm.DB, enrollmentID uint) ([]entity.Grade, error) {
	var grades []entity.Grade
	err := db.Where("enrollment_id = ?", enrollmentID).
		Preload("GradeComponent").
		Find(&grades).Error
	return grades, err
}
