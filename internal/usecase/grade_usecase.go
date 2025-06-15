package usecase

import (
	"context"
	"fmt"
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
	"tugasakhir/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// type GradeUseCase struct {
// 	DB              *gorm.DB
// 	Log             *logrus.Logger
// 	Validate        *validator.Validate
// 	GradeRepository *repository.GradeRepository
// }

// func NewGradeUseCase(
// 	db *gorm.DB,
// 	log *logrus.Logger,
// 	validate *validator.Validate,
// 	enrollmentRepository *repository.GradeRepository,

// ) *GradeUseCase {

// 	return &GradeUseCase{
// 		DB:              db,
// 		Log:             log,
// 		Validate:        validate,
// 		GradeRepository: enrollmentRepository,
// 	}

// }

type GradeUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	GradeRepo      *repository.GradeRepository
	ScheduleRepo   *repository.ScheduleRepository
	AttendanceRepo *repository.AttendanceRepository
	CourseRepo     *repository.CourseRepository
}

func NewGradeUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	gradeRepo *repository.GradeRepository,
	scheduleRepo *repository.ScheduleRepository,
	attendanceRepo *repository.AttendanceRepository,
	courseRepo *repository.CourseRepository,
) *GradeUseCase {
	return &GradeUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		GradeRepo:      gradeRepo,
		ScheduleRepo:   scheduleRepo,
		AttendanceRepo: attendanceRepo,
		CourseRepo:     courseRepo,
	}
}

func (c *GradeUseCase) GetCourseGrades(ctx context.Context, userID uint) ([]model.GradeResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var student entity.Student
	if err := tx.Where("user_id = ?", userID).First(&student).Error; err != nil {
		return nil, fmt.Errorf("mahasiswa tidak ditemukan")
	}

	var enrollments []entity.Enrollment
	if err := tx.Preload("Course").Preload("Grade.GradeComponent").Where("student_npm = ?", student.Npm).Find(&enrollments).Error; err != nil {
		return nil, fmt.Errorf("gagal ambil enrollment")
	}

	var reports []model.GradeResponse
	for _, e := range enrollments {
		hadir, _ := c.AttendanceRepo.CountAttendance(tx, userID, e.CourseCode)
		totalMeetings, _ := c.ScheduleRepo.GetTotalMeetingsByCourseCode(tx, e.CourseCode)

		persentase := float64(hadir) / float64(totalMeetings) * 100
		scoreAbsensi := persentase * 20 / 100 // weight absensi = 20%

		// Simpan nilai absensi jika belum ada
		if !c.GradeRepo.AttendanceGradeAlreadyExists(tx, e.ID) {
			err := c.GradeRepo.SaveAttendanceGrade(tx, e.ID, scoreAbsensi)
			if err != nil {
				c.Log.Warnf("Gagal simpan nilai absensi: %v", err)
			}
		}

		// Ambil semua komponen nilai termasuk absensi
		var grades []entity.Grade
		tx.Where("enrollment_id = ?", e.ID).Preload("GradeComponent").Find(&grades)

		totalScore := 0.0
		var components []model.GradeComponentScoreResponse

		for _, g := range grades {
			component := model.GradeComponentScoreResponse{
				Name:  g.GradeComponent.Name,
				Score: g.Score * float64(g.GradeComponent.Weight) / 100,
			}
			components = append(components, component)
			totalScore += component.Score
		}

		report := model.GradeResponse{
			CourseName: e.Course.Name,
			Components: components,
			TotalScore: totalScore,
		}

		reports = append(reports, report)
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Gagal commit transaksi: %+v", err)
		return nil, err
	}

	return reports, nil
}

// func (c *GradeUseCase) GetCourseGrades(ctx context.Context, userID uint) ([]model.GradeResponse, error) {
// 	// start transaction
// 	tx := c.DB.WithContext(ctx).Begin()
// 	defer tx.Rollback()

// 	var student entity.Student
// 	c.DB.Where("user_id = ?", userID).First(&student)

// 	enrollments := []entity.Enrollment{}
// 	c.DB.Where("student_npm = ?", student.Npm).Find(&enrollments)

// 	courses, err := c.CourseRepo.GetCoursesByNPM(tx, student.Npm)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal ambil course")
// 	}

// 	var reports []model.GradeResponse

// 	for _, e := range enrollments {
// 		var courseName string
// 		c.DB.Model(&model.Course{}).Where("code = ?", e.CourseCode).Pluck("name", &courseName)

// 		totalMeetings, _ := c.ScheduleRepo.GetTotalMeetingsByCourseCode(e.CourseCode)
// 		hadir, _ := c.AttendanceRepo.CountAttendance(userID, e.CourseCode)

// 		persentase := float64(hadir) / float64(totalMeetings) * 100
// 		scoreAbsensi := persentase * 20 / 100 // weight absensi = 20%

// 		if !c.GradeRepo.AttendanceGradeAlreadyExists(e.ID) {
// 			c.GradeRepo.SaveAttendanceGrade(e.ID, scoreAbsensi)
// 		}

// 		grades, _ := c.GradeRepo.GetAllGradesByEnrollmentID(e.ID)
// 		var components []model.GradeComponentScore
// 		totalScore := 0.0

// 		for _, g := range grades {
// 			component := model.GradeComponentScore{
// 				Name:  g.GradeComponent.Name,
// 				Score: float64(g.Score) * float64(g.GradeComponent.Weight) / 100,
// 			}
// 			components = append(components, component)
// 			totalScore += component.Score
// 		}

// 		report := model.GradeResponse{
// 			CourseName: courseName,
// 			TotalScore: totalScore,
// 			Components: components,
// 		}

// 		reports = append(reports, report)
// 	}

// 	// commit db transaction
// 	if err := tx.Commit().Error; err != nil {
// 		c.Log.Warnf("Failed commit transaction : %+v", err)
// 		return nil, err
// 	}

// 	return reports, nil
// }
