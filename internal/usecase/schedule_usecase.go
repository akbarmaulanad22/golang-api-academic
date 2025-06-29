package usecase

import (
	"context"
	"tugasakhir/internal/entity"
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

	responses := make([]model.ScheduleResponse, len(schedules))
	for i, schedule := range schedules {
		responses[i] = *converter.ScheduleToResponse(&schedule)
	}

	return responses, nil

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

	responses := make([]model.ScheduleResponse, len(schedules))
	for i, schedule := range schedules {
		responses[i] = *converter.ScheduleToResponse(&schedule)
	}

	return responses, nil
}

func (c *ScheduleUseCase) Create(ctx context.Context, request *model.CreateScheduleRequest) (*model.ScheduleAdminResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	schedule := &entity.Schedule{
		Date:         request.Date,
		StartAt:      request.StartAt,
		EndAt:        request.EndAt,
		CourseCode:   request.CourseCode,
		LecturerNIDN: request.LecturerNIDN,
		ClassroomID:  request.ClassroomID,
	}

	if err := c.ScheduleRepository.Create(tx, schedule); err != nil {
		c.Log.WithError(err).Error("error creating schedule")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating schedule")
		return nil, err
	}

	return converter.ScheduleToAdminResponse(schedule), nil

}

func (c *ScheduleUseCase) List(ctx context.Context) ([]model.ScheduleAdminResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	schedules, err := c.ScheduleRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error creating schedule")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating schedule")
		return nil, err
	}

	responses := make([]model.ScheduleAdminResponse, len(schedules))
	for i, schedule := range schedules {
		responses[i] = *converter.ScheduleToAdminResponse(&schedule)
	}

	return responses, nil
}

func (c *ScheduleUseCase) Update(ctx context.Context, request *model.UpdateScheduleRequest) (*model.ScheduleAdminResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	schedule := new(entity.Schedule)
	if err := c.ScheduleRepository.FindById(tx, schedule, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find schedule")
		return nil, err
	}

	schedule.Date = request.Date
	schedule.StartAt = request.StartAt
	schedule.EndAt = request.EndAt
	schedule.CourseCode = request.CourseCode
	schedule.LecturerNIDN = request.LecturerNIDN
	schedule.ClassroomID = request.ClassroomID

	if err := c.ScheduleRepository.Update(tx, schedule); err != nil {
		c.Log.WithError(err).Error("failed to update schedule")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.ScheduleToAdminResponse(schedule), nil
}

func (c *ScheduleUseCase) Delete(ctx context.Context, request *model.DeleteScheduleRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	schedule := new(entity.Schedule)
	if err := c.ScheduleRepository.FindById(tx, schedule, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find schedule")
		return err
	}

	if err := c.ScheduleRepository.Delete(tx, schedule); err != nil {
		c.Log.WithError(err).Error("failed to update schedule")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return err
	}

	return nil
}
