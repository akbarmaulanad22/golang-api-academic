package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ClassroomRepository struct {
	Repository[entity.Classroom]
	Log *logrus.Logger
}

func NewClassroomRepository(log *logrus.Logger) *ClassroomRepository {

	return &ClassroomRepository{Log: log}

}

func (r *ClassroomRepository) FindAll(db *gorm.DB) ([]entity.Classroom, error) {

	var classrooms []entity.Classroom
	if err := db.Find(&classrooms).Error; err != nil {
		return nil, err
	}

	return classrooms, nil
}
