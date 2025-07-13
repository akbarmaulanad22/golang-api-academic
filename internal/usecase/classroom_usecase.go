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

type ClassroomUseCase struct {
	DB                  *gorm.DB
	Log                 *logrus.Logger
	Validate            *validator.Validate
	ClassroomRepository *repository.ClassroomRepository
}

func NewClassroomUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	classroomRepository *repository.ClassroomRepository,

) *ClassroomUseCase {

	return &ClassroomUseCase{
		DB:                  db,
		Log:                 log,
		Validate:            validate,
		ClassroomRepository: classroomRepository,
	}

}

func (c *ClassroomUseCase) Create(ctx context.Context, request *model.CreateClassroomRequest) (*model.ClassroomResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	classroom := &entity.Classroom{
		Name:     request.Name,
		Capacity: request.Capacity,
		Location: request.Location,
	}

	if err := c.ClassroomRepository.Create(tx, classroom); err != nil {
		c.Log.WithError(err).Error("error creating classroom")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating classroom")
		return nil, err
	}

	return converter.ClassroomToResponse(classroom), nil

}

func (c *ClassroomUseCase) List(ctx context.Context) ([]model.ClassroomResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	classrooms, err := c.ClassroomRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error creating classroom")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating classroom")
		return nil, err
	}

	responses := make([]model.ClassroomResponse, len(classrooms))
	for i, contact := range classrooms {
		responses[i] = *converter.ClassroomToResponse(&contact)
	}

	return responses, nil
}

func (c *ClassroomUseCase) Update(ctx context.Context, request *model.UpdateClassroomRequest) (*model.ClassroomResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	classroom := new(entity.Classroom)
	if err := c.ClassroomRepository.FindById(tx, classroom, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find classroom")
		return nil, err
	}

	classroom.Name = request.Name
	classroom.Capacity = request.Capacity
	classroom.Location = request.Location

	if err := c.ClassroomRepository.Update(tx, classroom); err != nil {
		c.Log.WithError(err).Error("failed to update classroom")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.ClassroomToResponse(classroom), nil
}

func (c *ClassroomUseCase) Delete(ctx context.Context, request *model.DeleteClassroomRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	classroom := new(entity.Classroom)
	if err := c.ClassroomRepository.FindById(tx, classroom, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find classroom")
		return err
	}

	if err := c.ClassroomRepository.Delete(tx, classroom); err != nil {
		c.Log.WithError(err).Error("failed to update classroom")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return err
	}

	return nil
}
