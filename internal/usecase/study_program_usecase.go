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

type StudyProgramUseCase struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	Validate               *validator.Validate
	StudyProgramRepository *repository.StudyProgramRepository
}

func NewStudyProgramUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	studyProgramRepository *repository.StudyProgramRepository,

) *StudyProgramUseCase {

	return &StudyProgramUseCase{
		DB:                     db,
		Log:                    log,
		Validate:               validate,
		StudyProgramRepository: studyProgramRepository,
	}

}

func (c *StudyProgramUseCase) Create(ctx context.Context, request *model.CreateStudyProgramRequest) (*model.StudyProgramResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	studyProgram := &entity.StudyProgram{
		Name:             request.Name,
		Level:            request.Level,
		Accreditation:    request.Accreditation,
		DurationSemester: request.DurationSemester,
		FacultyId:        request.FacultyId,
	}

	if err := c.StudyProgramRepository.Create(tx, studyProgram); err != nil {
		c.Log.WithError(err).Error("error creating study program")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating study program")
		return nil, err
	}

	return converter.StudyProgramToResponse(studyProgram), nil

}

func (c *StudyProgramUseCase) List(ctx context.Context) ([]model.StudyProgramResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	studyPrograms, err := c.StudyProgramRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error creating study program")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating study program")
		return nil, err
	}

	responses := make([]model.StudyProgramResponse, len(studyPrograms))
	for i, contact := range studyPrograms {
		responses[i] = *converter.StudyProgramToResponse(&contact)
	}

	return responses, nil
}

func (c *StudyProgramUseCase) Update(ctx context.Context, request *model.UpdateStudyProgramRequest) (*model.StudyProgramResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	studyProgram := new(entity.StudyProgram)
	if err := c.StudyProgramRepository.FindById(tx, studyProgram, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find study program")
		return nil, err
	}

	studyProgram.Name = request.Name
	studyProgram.Level = request.Level
	studyProgram.Accreditation = request.Accreditation
	studyProgram.DurationSemester = request.DurationSemester
	studyProgram.FacultyId = request.FacultyId

	if err := c.StudyProgramRepository.Update(tx, studyProgram); err != nil {
		c.Log.WithError(err).Error("failed to update study program")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.StudyProgramToResponse(studyProgram), nil
}

func (c *StudyProgramUseCase) Delete(ctx context.Context, request *model.DeleteStudyProgramRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	studyProgram := new(entity.StudyProgram)
	if err := c.StudyProgramRepository.FindById(tx, studyProgram, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find study program")
		return err
	}

	if err := c.StudyProgramRepository.Delete(tx, studyProgram); err != nil {
		c.Log.WithError(err).Error("failed to update study program")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return err
	}

	return nil
}
