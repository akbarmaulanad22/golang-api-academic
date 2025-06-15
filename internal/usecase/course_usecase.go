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

type CourseUseCase struct {
	DB               *gorm.DB
	Log              *logrus.Logger
	Validate         *validator.Validate
	CourseRepository *repository.CourseRepository
}

func NewCourseUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	enrollmentRepository *repository.CourseRepository,

) *CourseUseCase {

	return &CourseUseCase{
		DB:               db,
		Log:              log,
		Validate:         validate,
		CourseRepository: enrollmentRepository,
	}

}

func (c *CourseUseCase) ListByLecturerUserID(ctx context.Context, request *model.ListCourseRequest) ([]model.CourseResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	courses, err := c.CourseRepository.FindAllByNIDNUserID(tx, request.UserID)
	if err != nil {
		c.Log.WithError(err).Error("failed to find courses")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.CourseToResponses(courses), nil
}
