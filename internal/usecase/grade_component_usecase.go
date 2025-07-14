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

type GradeComponentUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	GradeComponentRepository *repository.GradeComponentRepository
}

func NewGradeComponentUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	gradeComponentRepository *repository.GradeComponentRepository,

) *GradeComponentUseCase {

	return &GradeComponentUseCase{
		DB:                       db,
		Log:                      log,
		Validate:                 validate,
		GradeComponentRepository: gradeComponentRepository,
	}

}

func (c *GradeComponentUseCase) Create(ctx context.Context, request *model.CreateGradeComponentRequest) (*model.GradeComponentAdminResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	gradeComponent := &entity.GradeComponent{
		Name:   request.Name,
		Weight: request.Weight,
	}

	if err := c.GradeComponentRepository.Create(tx, gradeComponent); err != nil {
		c.Log.WithError(err).Error("error creating grade component")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating grade component")
		return nil, err
	}

	return converter.GradeComponentToResponse(gradeComponent), nil

}

func (c *GradeComponentUseCase) List(ctx context.Context) ([]model.GradeComponentAdminResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	gradeComponents, err := c.GradeComponentRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error creating grade component")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating grade component")
		return nil, err
	}

	responses := make([]model.GradeComponentAdminResponse, len(gradeComponents))
	for i, contact := range gradeComponents {
		responses[i] = *converter.GradeComponentToResponse(&contact)
	}

	return responses, nil
}

func (c *GradeComponentUseCase) Update(ctx context.Context, request *model.UpdateGradeComponentRequest) (*model.GradeComponentAdminResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	gradeComponent := new(entity.GradeComponent)
	if err := c.GradeComponentRepository.FindById(tx, gradeComponent, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find grade component")
		return nil, err
	}

	gradeComponent.Name = request.Name
	gradeComponent.Weight = request.Weight

	if err := c.GradeComponentRepository.Update(tx, gradeComponent); err != nil {
		c.Log.WithError(err).Error("failed to update grade component")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.GradeComponentToResponse(gradeComponent), nil
}

func (c *GradeComponentUseCase) Delete(ctx context.Context, request *model.DeleteGradeComponentRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	gradeComponent := new(entity.GradeComponent)
	if err := c.GradeComponentRepository.FindById(tx, gradeComponent, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find grade component")
		return err
	}

	if err := c.GradeComponentRepository.Delete(tx, gradeComponent); err != nil {
		c.Log.WithError(err).Error("failed to update grade component")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return err
	}

	return nil
}
