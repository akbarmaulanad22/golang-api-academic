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

type EnrollmentUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	EnrollmentRepository *repository.EnrollmentRepository
}

func NewEnrollmentUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	enrollmentRepository *repository.EnrollmentRepository,

) *EnrollmentUseCase {

	return &EnrollmentUseCase{
		DB:                   db,
		Log:                  log,
		Validate:             validate,
		EnrollmentRepository: enrollmentRepository,
	}

}

func (c *EnrollmentUseCase) GetEnrollmentByStudentUserID(ctx context.Context, userID uint) ([]model.EnrollmentResponse, error) {
	// start transaction
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	schedules, err := c.EnrollmentRepository.GetEnrollmentByStudentUserID(tx, userID)

	if err != nil {
		return nil, err
	}

	// commit db transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	return converter.EnrollmentToResponses(schedules), nil
}
