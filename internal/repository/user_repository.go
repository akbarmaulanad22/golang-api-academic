package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {

	return &UserRepository{Log: log}

}

func (r *UserRepository) CountByUsername(db *gorm.DB, username string) (int64, error) {

	var total int64

	err := db.Model(&entity.User{}).Where("username = ?", username).Count(&total).Error
	return total, err

}

func (r *UserRepository) FindByUsername(db *gorm.DB, user *entity.User, id any) error {
	return db.Where("username = ?", id).Take(user).Error
}

func (r *UserRepository) FindByToken(db *gorm.DB, user *entity.User, token string) error {
	return db.Where("token = ?", token).First(user).Error
}

func (r *UserRepository) FindAllJoinLecturer(db *gorm.DB) ([]entity.User, error) {

	var users []entity.User
	if err := db.Preload("Lecturer").Joins("JOIN lecturers ON lecturers.user_id = users.id").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
