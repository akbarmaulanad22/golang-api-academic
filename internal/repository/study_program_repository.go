package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StudyProgramRepository struct {
	Repository[entity.StudyProgram]
	Log *logrus.Logger
}

func NewStudyProgramRepository(log *logrus.Logger) *StudyProgramRepository {

	return &StudyProgramRepository{Log: log}

}

func (r *StudyProgramRepository) FindAll(db *gorm.DB) ([]entity.StudyProgram, error) {

	var studyProgram []entity.StudyProgram
	if err := db.Preload("Faculty").Find(&studyProgram).Error; err != nil {
		return nil, err
	}

	return studyProgram, nil
}
