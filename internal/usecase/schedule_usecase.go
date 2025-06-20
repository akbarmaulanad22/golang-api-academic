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

type ScheduleUseCase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	ScheduleRepository *repository.ScheduleRepository
}

func NewScheduleUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	scheduleRepository *repository.ScheduleRepository,

) *ScheduleUseCase {

	return &ScheduleUseCase{
		DB:                 db,
		Log:                log,
		Validate:           validate,
		ScheduleRepository: scheduleRepository,
	}

}

func (c *ScheduleUseCase) ListScheduleByStudentUserID(ctx context.Context, request *model.ListScheduleRequest) ([]model.ScheduleResponse, error) {
	// start transaction
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	schedules, err := c.ScheduleRepository.FindAllScheduleByStudentUserID(tx, request.UserID)

	if err != nil {
		return nil, err
	}

	// commit db transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	return converter.ScheduleToResponses(schedules), nil
}

func (c *ScheduleUseCase) ListScheduleByLecturerUserID(ctx context.Context, request *model.ListScheduleRequest) ([]model.ScheduleResponse, error) {
	// start transaction
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	schedules, err := c.ScheduleRepository.FindAllSchedulesByLecturerUserID(tx, request.UserID)

	if err != nil {
		return nil, err
	}

	// commit db transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	return converter.ScheduleToResponses(schedules), nil
}
