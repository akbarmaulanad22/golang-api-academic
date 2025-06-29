package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FacultyRepository struct {
	Repository[entity.Faculty]
	Log *logrus.Logger
}

func NewFacultyRepository(log *logrus.Logger) *FacultyRepository {

	return &FacultyRepository{Log: log}

}

func (r *FacultyRepository) FindAll(db *gorm.DB) ([]entity.Faculty, error) {

	var studyProgram []entity.Faculty
	if err := db.Find(&studyProgram).Error; err != nil {
		return nil, err
	}

	return studyProgram, nil
}
