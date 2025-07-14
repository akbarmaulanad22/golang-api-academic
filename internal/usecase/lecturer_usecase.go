package usecase

import (
	"context"
	"fmt"
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
	"tugasakhir/internal/model/converter"
	"tugasakhir/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LecturerUseCase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	LecturerRepository *repository.LecturerRepository
	UserRepository     *repository.UserRepository
}

func NewLecturerUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	lecturerRepository *repository.LecturerRepository,
	userRepository *repository.UserRepository,

) *LecturerUseCase {

	return &LecturerUseCase{
		DB:                 db,
		Log:                log,
		Validate:           validate,
		LecturerRepository: lecturerRepository,
		UserRepository:     userRepository,
	}

}

func (c *LecturerUseCase) Create(ctx context.Context, request *model.CreateLecturerRequest) (*model.LecturerResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	// get username
	total, err := c.UserRepository.CountByUsername(tx, request.Username)
	if err != nil {
		c.Log.Warnf("Failed count user from database : %+v", err)
		return nil, err
	}

	// check if username already exists
	if total > 0 {
		c.Log.Warnf("Please choose another username : %+v", err)
		return nil, fmt.Errorf("please choose another username %d", 0)
	}

	// encrypt password
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, err
	}

	lecturer := &entity.Lecturer{
		Nidn:       request.NIDN,
		Name:       request.Name,
		Gender:     request.Gender,
		Degree:     request.Degree,
		IsFullTime: request.IsFullTime,
		User: entity.User{
			Password: string(password),
			Username: request.Username,
			RoleID:   2, // lecturer
		},
	}

	// create user
	if err := c.LecturerRepository.Create(tx, lecturer); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating lecturer")
		return nil, err
	}

	return converter.LecturerToResponse(lecturer), nil

}

func (c *LecturerUseCase) List(ctx context.Context) ([]model.LecturerResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	lecturers, err := c.LecturerRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error creating lecturer")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating lecturer")
		return nil, err
	}

	responses := make([]model.LecturerResponse, len(lecturers))
	for i, lecturer := range lecturers {
		responses[i] = *converter.LecturerToResponse(&lecturer)
	}

	return responses, nil
}

func (c *LecturerUseCase) Update(ctx context.Context, request *model.UpdateLecturerRequest) (*model.LecturerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	lecturer := new(entity.Lecturer)
	if err := c.LecturerRepository.FindByNIDN(tx, lecturer, request.NIDN); err != nil {
		c.Log.WithError(err).Error("failed to find lecturer")
		return nil, err
	}

	if request.Password != "" {
		// encrypt password
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
			return nil, err
		}
		lecturer.User.Password = string(password)
	}

	lecturer.Nidn = request.NIDN
	lecturer.Name = request.Name
	lecturer.Gender = request.Gender
	lecturer.Degree = request.Degree
	lecturer.IsFullTime = request.IsFullTime

	if err := c.LecturerRepository.Update(tx, lecturer); err != nil {
		c.Log.WithError(err).Error("failed to update lecturer")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.LecturerToResponse(lecturer), nil
}

func (c *LecturerUseCase) Delete(ctx context.Context, request *model.DeleteLecturerRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	lecturer := new(entity.Lecturer)
	if err := c.LecturerRepository.FindByNIDN(tx, lecturer, request.NIDN); err != nil {
		c.Log.WithError(err).Error("failed to find lecturer")
		return err
	}

	if err := c.LecturerRepository.Delete(tx, lecturer); err != nil {
		c.Log.WithError(err).Error("failed to update lecturer")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return err
	}

	return nil
}
