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

type StudentUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	StudentRepository *repository.StudentRepository
	UserRepository    *repository.UserRepository
}

func NewStudentUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	studentRepository *repository.StudentRepository,
	userRepository *repository.UserRepository,

) *StudentUseCase {

	return &StudentUseCase{
		DB:                db,
		Log:               log,
		Validate:          validate,
		StudentRepository: studentRepository,
		UserRepository:    userRepository,
	}

}

func (c *StudentUseCase) ListByCourseCode(ctx context.Context, request *model.ListStudentRequest) ([]model.StudentResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	students, err := c.StudentRepository.FindAllStudentByCouseCode(tx, request.CourseCode)
	if err != nil {
		c.Log.WithError(err).Error("failed to find students")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	responses := make([]model.StudentResponse, len(students))
	for i, student := range students {
		responses[i] = *converter.StudentToResponse(&student)
	}

	return responses, nil
}

func (c *StudentUseCase) Create(ctx context.Context, request *model.CreateStudentRequest) (*model.StudentAdminResponse, error) {

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

	student := &entity.Student{
		Npm:              request.Npm,
		Class:            request.Class,
		RegistrationWave: request.RegistrationWave,
		RegistrationDate: request.RegistrationDate,
		StudyProgramId:   request.StudyProgramId,
		Biodata: entity.StudentBio{
			Name:      request.Name,
			DateBirth: request.DateBirth,
			Address:   request.Address,
			Gender:    request.Gender,
		},
		User: &entity.User{
			Password: string(password),
			Username: request.Username,
			RoleID:   3, // student
		},
	}

	// create user
	if err := c.StudentRepository.Create(tx, student); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating student")
		return nil, err
	}

	return converter.StudentAdminToResponse(student), nil

}

func (c *StudentUseCase) List(ctx context.Context) ([]model.StudentAdminResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	students, err := c.StudentRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error creating student")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating student")
		return nil, err
	}

	responses := make([]model.StudentAdminResponse, len(students))
	for i, student := range students {
		responses[i] = *converter.StudentAdminToResponse(&student)
	}

	return responses, nil
}

func (c *StudentUseCase) Update(ctx context.Context, request *model.UpdateStudentRequest) (*model.StudentAdminResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	student := new(entity.Student)
	if err := c.StudentRepository.FindByNpm(tx, student, request.Npm); err != nil {
		c.Log.WithError(err).Error("failed to find student")
		return nil, err
	}

	if request.Password != "" {
		// encrypt password
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
			return nil, err
		}

		student.User.Password = string(password)
	}

	student.Class = request.Class
	student.RegistrationWave = request.RegistrationWave
	student.RegistrationDate = request.RegistrationDate
	student.StudyProgramId = request.StudyProgramId
	student.Biodata.Name = request.Name
	student.Biodata.DateBirth = request.DateBirth
	student.Biodata.Address = request.Address
	student.Biodata.Gender = request.Gender

	if err := c.StudentRepository.Update(tx, student); err != nil {
		c.Log.WithError(err).Error("failed to update student")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.StudentAdminToResponse(student), nil
}

func (c *StudentUseCase) Delete(ctx context.Context, request *model.DeleteStudentRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	student := new(entity.Student)
	if err := c.StudentRepository.FindByNpm(tx, student, request.Npm); err != nil {
		c.Log.WithError(err).Error("failed to find student")
		return err
	}

	if err := c.StudentRepository.Delete(tx, student); err != nil {
		c.Log.WithError(err).Error("failed to update student")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return err
	}

	return nil
}
