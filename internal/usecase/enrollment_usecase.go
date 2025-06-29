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

func (c *EnrollmentUseCase) ListByStudentUserID(ctx context.Context, request *model.ListEnrollmentRequest) ([]model.EnrollmentResponse, error) {
	// start transaction
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	enrollments, err := c.EnrollmentRepository.FindAllEnrollmentByStudentUserID(tx, request.UserID)

	if err != nil {
		return nil, err
	}

	// commit db transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	responses := make([]model.EnrollmentResponse, len(enrollments))
	for i, enrollment := range enrollments {
		responses[i] = *converter.EnrollmentToResponse(&enrollment)
	}

	return responses, nil
}

func (c *EnrollmentUseCase) Create(ctx context.Context, request *model.CreateEnrollmentRequest) (*model.EnrollmentAdminResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	enrollment := &entity.Enrollment{
		Status:           request.Status,
		AcademicYear:     request.AcademicYear,
		RegistrationDate: request.RegistrationDate,
		StudentNpm:       request.StudentNpm,
		CourseCode:       request.CourseCode,
	}

	if err := c.EnrollmentRepository.Create(tx, enrollment); err != nil {
		c.Log.WithError(err).Error("error creating enrollment")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating enrollment")
		return nil, err
	}

	return converter.EnrollmentToAdminResponse(enrollment), nil

}

func (c *EnrollmentUseCase) List(ctx context.Context) ([]model.EnrollmentAdminResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	enrollments, err := c.EnrollmentRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error creating enrollment")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating enrollment")
		return nil, err
	}

	responses := make([]model.EnrollmentAdminResponse, len(enrollments))
	for i, contact := range enrollments {
		responses[i] = *converter.EnrollmentToAdminResponse(&contact)
	}

	return responses, nil
}

func (c *EnrollmentUseCase) Update(ctx context.Context, request *model.UpdateEnrollmentRequest) (*model.EnrollmentAdminResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	enrollment := new(entity.Enrollment)
	if err := c.EnrollmentRepository.FindById(tx, enrollment, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find enrollment")
		return nil, err
	}

	enrollment.ID = request.ID
	enrollment.Status = request.Status
	enrollment.AcademicYear = request.AcademicYear
	enrollment.RegistrationDate = request.RegistrationDate
	enrollment.StudentNpm = request.StudentNpm
	enrollment.CourseCode = request.CourseCode

	if err := c.EnrollmentRepository.Update(tx, enrollment); err != nil {
		c.Log.WithError(err).Error("failed to update enrollment")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.EnrollmentToAdminResponse(enrollment), nil
}

func (c *EnrollmentUseCase) Delete(ctx context.Context, request *model.DeleteEnrollmentRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	enrollment := new(entity.Enrollment)
	if err := c.EnrollmentRepository.FindById(tx, enrollment, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find enrollment")
		return err
	}

	if err := c.EnrollmentRepository.Delete(tx, enrollment); err != nil {
		c.Log.WithError(err).Error("failed to update enrollment")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return err
	}

	return nil
}
