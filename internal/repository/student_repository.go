package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StudentRepository struct {
	Repository[entity.Student]
	Log *logrus.Logger
}

func NewStudentRepository(log *logrus.Logger) *StudentRepository {

	return &StudentRepository{Log: log}

}

func (r *StudentRepository) FindAllStudentByCouseCode(db *gorm.DB, courseCode string) ([]entity.Student, error) {

	var students []entity.Student
	if err := db.
		Joins("JOIN enrollments ON students.npm = enrollments.student_npm").
		Where("enrollments.course_code = ?", courseCode).
		Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}

func (r *StudentRepository) GetStudentUserIDByNPM(db *gorm.DB, npm uint) (uint, error) {

	var userID uint
	err := db.Raw(`
        SELECT user_id FROM students st
        WHERE st.npm = ?
    `, npm).Scan(&userID).Error
	return userID, err
}

func (r *StudentRepository) FindAll(db *gorm.DB) ([]entity.Student, error) {

	var lecturers []entity.Student
	if err := db.Preload("User").Find(&lecturers).Error; err != nil {
		return nil, err
	}

	return lecturers, nil
}

func (r *StudentRepository) FindByNpm(db *gorm.DB, user *entity.Student, npm uint) error {
	return db.Preload("User").Where("npm = ?", npm).First(user).Error
}
