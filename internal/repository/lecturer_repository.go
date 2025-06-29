package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LecturerRepository struct {
	Repository[entity.Lecturer]
	Log *logrus.Logger
}

func NewLecturerRepository(log *logrus.Logger) *LecturerRepository {

	return &LecturerRepository{Log: log}

}

func (r *LecturerRepository) FindAll(db *gorm.DB) ([]entity.Lecturer, error) {

	var lecturers []entity.Lecturer
	if err := db.Preload("User").Find(&lecturers).Error; err != nil {
		return nil, err
	}

	return lecturers, nil
}

func (r *LecturerRepository) FindByNIDN(db *gorm.DB, user *entity.Lecturer, nidn uint) error {
	return db.Preload("User").Where("nidn = ?", nidn).First(user).Error
}
