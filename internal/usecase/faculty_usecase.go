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

type FacultyUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	FacultyRepository *repository.FacultyRepository
}

func NewFacultyUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	studyProgramRepository *repository.FacultyRepository,

) *FacultyUseCase {

	return &FacultyUseCase{
		DB:                db,
		Log:               log,
		Validate:          validate,
		FacultyRepository: studyProgramRepository,
	}

}

func (c *FacultyUseCase) Create(ctx context.Context, request *model.CreateFacultyRequest) (*model.FacultyResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	studyProgram := &entity.Faculty{
		Code:    request.Code,
		Name:    request.Name,
		Dekan:   request.Dekan,
		Address: request.Address,
	}

	if err := c.FacultyRepository.Create(tx, studyProgram); err != nil {
		c.Log.WithError(err).Error("error creating faculty")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating faculty")
		return nil, err
	}

	return converter.FacultyToResponse(studyProgram), nil

}

func (c *FacultyUseCase) List(ctx context.Context) ([]model.FacultyResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	studyPrograms, err := c.FacultyRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error creating faculty")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating faculty")
		return nil, err
	}

	responses := make([]model.FacultyResponse, len(studyPrograms))
	for i, contact := range studyPrograms {
		responses[i] = *converter.FacultyToResponse(&contact)
	}

	return responses, nil
}

func (c *FacultyUseCase) Update(ctx context.Context, request *model.UpdateFacultyRequest) (*model.FacultyResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	studyProgram := new(entity.Faculty)
	if err := c.FacultyRepository.FindById(tx, studyProgram, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find faculty")
		return nil, err
	}

	studyProgram.Name = request.Name
	studyProgram.Code = request.Code
	studyProgram.Dekan = request.Dekan
	studyProgram.Address = request.Address

	if err := c.FacultyRepository.Update(tx, studyProgram); err != nil {
		c.Log.WithError(err).Error("failed to update faculty")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.FacultyToResponse(studyProgram), nil
}

func (c *FacultyUseCase) Delete(ctx context.Context, request *model.DeleteFacultyRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	studyProgram := new(entity.Faculty)
	if err := c.FacultyRepository.FindById(tx, studyProgram, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find faculty")
		return err
	}

	if err := c.FacultyRepository.Delete(tx, studyProgram); err != nil {
		c.Log.WithError(err).Error("failed to update faculty")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return err
	}

	return nil
}
