package usecase

import (
	"context"
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
	"tugasakhir/internal/model/converter"
	"tugasakhir/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, repository *repository.UserRepository) *UserUseCase {

	return &UserUseCase{
		DB:         db,
		Log:        log,
		Validate:   validate,
		Repository: repository,
	}

}

func (c *UserUseCase) Create(ctx context.Context, request *model.UserRegisterRequest) (*model.UserResponse, error) {
	// validate
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body: %+v", err)
		return nil, err
	}

	// start transaction

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// get username
	total, err := c.Repository.CountByUsername(tx, request.Username)
	if err != nil {
		c.Log.Warnf("Failed count user from database : %+v", err)
		return nil, err
	}

	// check if username already exists
	if total > 0 {
		c.Log.Warnf("Please choose another username : %+v", err)
		return nil, err
	}

	// encrypt password
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, err
	}

	user := &entity.User{
		Password: string(password),
		Username: request.Username,
		RoleID:   request.RoleID,
	}

	// create user
	if err := c.Repository.Create(tx, user); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, err
	}

	// commit db transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	// end transaction

	return converter.UserToResponse(user), nil
}

func (c *UserUseCase) Login(ctx context.Context, request *model.UserLoginRequest) (*model.UserResponse, error) {
	// validate
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body: %+v", err)
		return nil, err
	}

	// start transaction

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// get user
	user := new(entity.User)
	err = c.Repository.FindByUsername(tx, user, request.Username)
	if err != nil {
		c.Log.Warnf("Cannot find username : %s, error: %+v", request.Username, err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.Log.Warnf("Failed to compare user password with bcrype hash : %+v", err)
		return nil, err
	}

	user.Token = uuid.New().String()
	if err := c.Repository.Update(tx, user); err != nil {
		c.Log.Warnf("Failed save user : %+v", err)
		return nil, err
	}

	// commit db transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	return converter.UserToTokenResponse(user), nil

}

func (c *UserUseCase) Logout(ctx context.Context, request *model.UserLogoutRequest) (bool, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return false, err
	}

	user := new(entity.User)
	if err := c.Repository.FindByUsername(tx, user, request.Username); err != nil {
		c.Log.Warnf("User not found : %+v", err)
		return false, err
	}

	user.Token = ""

	if err := c.Repository.Update(tx, user); err != nil {
		c.Log.Warnf("Failed save user : %+v", err)
		return false, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return false, err
	}

	return true, nil
}

func (c *UserUseCase) Delete(ctx context.Context, request *model.UserDeleteRequest) error {
	// validate
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body: %+v", err)
		return err
	}

	// start transaction

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// get user
	user := new(entity.User)
	err = c.Repository.FindByUsername(tx, user, request.Username)
	if err != nil {
		c.Log.Warnf("User not found : %+v", err)
		return err
	}

	err = c.Repository.Delete(tx, user)
	if err != nil {
		c.Log.Warnf("Failed delete user : %+v", err)
		return err
	}

	// commit db transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return err
	}

	return nil
}

func (c *UserUseCase) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, err
	}

	user := new(entity.User)
	if err := c.Repository.FindByToken(tx, user, request.Token); err != nil {
		c.Log.Warnf("Failed find user by token : %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	return &model.Auth{Username: user.Username, ID: user.ID}, nil
}

func (c *UserUseCase) Current(ctx context.Context, request *model.GetUserRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, err
	}

	user := new(entity.User)
	if err := c.Repository.FindByUsername(tx, user, request.Username); err != nil {
		c.Log.Warnf("Failed find user by id : %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	return converter.UserToResponse(user), nil
}
