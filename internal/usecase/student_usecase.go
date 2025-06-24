package usecase

import (
	"context"
	"tugasakhir/internal/model"
	"tugasakhir/internal/model/converter"
	"tugasakhir/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StudentUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	StudentRepository *repository.StudentRepository
}

func NewStudentUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	studentRepository *repository.StudentRepository,

) *StudentUseCase {

	return &StudentUseCase{
		DB:                db,
		Log:               log,
		Validate:          validate,
		StudentRepository: studentRepository,
	}

}

func (c *StudentUseCase) ListByCourseCode(ctx context.Context, request *model.ListStudentRequest) ([]model.StudentResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	students, err := c.StudentRepository.FindAllStudentByCouseCode(tx, request.CourseCode)
	if err != nil {
		c.Log.WithError(err).Error("failed to find students")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.StudentToResponses(students), nil
}
