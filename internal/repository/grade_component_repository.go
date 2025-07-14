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
