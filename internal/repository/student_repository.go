package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StudentRepository struct {
	Repository[entity.Student]
	Log *logrus.Logger
}

func NewStudentRepository(log *logrus.Logger) *StudentRepository {

	return &StudentRepository{Log: log}

}

func (r *StudentRepository) FindByUserId(db *gorm.DB, entity *entity.Student, userId uint) error {
	return db.Where("user_id = ?", userId).Take(entity).Error
}
