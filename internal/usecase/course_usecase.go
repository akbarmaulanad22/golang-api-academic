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

	responses := make([]model.CourseResponse, len(courses))
	for i, address := range courses {
		responses[i] = *converter.CourseToResponse(&address)
	}

	return responses, nil
}

func (c *CourseUseCase) Create(ctx context.Context, request *model.CreateCourseRequest) (*model.CourseResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	course := &entity.Course{
		Code:     request.Code,
		Name:     request.Name,
		Sks:      request.Sks,
		Semester: request.Semester,
		// TotalMeetings: request.TotalMeetings,
		LecturerNIDN: uint(request.LecturerNIDN),
	}

	if err := c.CourseRepository.Create(tx, course); err != nil {
		c.Log.WithError(err).Error("error creating course")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating course")
		return nil, err
	}

	return converter.CourseToResponse(course), nil

}

func (c *CourseUseCase) List(ctx context.Context) ([]model.CourseResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	courses, err := c.CourseRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error creating course")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating course")
		return nil, err
	}

	responses := make([]model.CourseResponse, len(courses))
	for i, contact := range courses {
		responses[i] = *converter.CourseToResponse(&contact)
	}

	return responses, nil
}

func (c *CourseUseCase) Update(ctx context.Context, request *model.UpdateCourseRequest) (*model.CourseResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	course := new(entity.Course)
	if err := c.CourseRepository.FindByCode(tx, course, request.Code); err != nil {
		c.Log.WithError(err).Error("failed to find course")
		return nil, err
	}

	course.Name = request.Name
	course.Sks = request.Sks
	course.Semester = request.Semester
	// course.TotalMeetings = request.TotalMeetings
	course.LecturerNIDN = request.LecturerNIDN

	if err := c.CourseRepository.Update(tx, course); err != nil {
		c.Log.WithError(err).Error("failed to update course")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.CourseToResponse(course), nil
}

func (c *CourseUseCase) Delete(ctx context.Context, request *model.DeleteCourseRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	course := new(entity.Course)
	if err := c.CourseRepository.FindByCode(tx, course, request.Code); err != nil {
		c.Log.WithError(err).Error("failed to find course")
		return err
	}

	if err := c.CourseRepository.Delete(tx, course); err != nil {
		c.Log.WithError(err).Error("failed to update course")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return err
	}

	return nil
}
