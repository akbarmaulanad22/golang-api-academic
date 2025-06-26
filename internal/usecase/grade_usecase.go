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
	"gorm.io/gorm"
)

type GradeUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	GradeRepo      *repository.GradeRepository
	ScheduleRepo   *repository.ScheduleRepository
	AttendanceRepo *repository.AttendanceRepository
	CourseRepo     *repository.CourseRepository
	EnrollmentRepo *repository.EnrollmentRepository
}

func NewGradeUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	gradeRepo *repository.GradeRepository,
	scheduleRepo *repository.ScheduleRepository,
	attendanceRepo *repository.AttendanceRepository,
	courseRepo *repository.CourseRepository,
	enrollmentRepo *repository.EnrollmentRepository,
) *GradeUseCase {
	return &GradeUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		GradeRepo:      gradeRepo,
		ScheduleRepo:   scheduleRepo,
		AttendanceRepo: attendanceRepo,
		CourseRepo:     courseRepo,
		EnrollmentRepo: enrollmentRepo,
	}
}

func (c *GradeUseCase) ListByStudentUserID(ctx context.Context, request *model.ListGradeRequest) ([]model.GradeResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var student entity.Student
	if err := tx.Where("user_id = ?", request.UserID).First(&student).Error; err != nil {
		return nil, fmt.Errorf("mahasiswa tidak ditemukan")
	}

	var enrollments []entity.Enrollment
	if err := tx.Preload("Course").Preload("Grade.GradeComponent").Where("student_npm = ?", student.Npm).Find(&enrollments).Error; err != nil {
		return nil, fmt.Errorf("gagal ambil enrollment")
	}

	var reports []model.GradeResponse
	for _, e := range enrollments {
		hadir, _ := c.AttendanceRepo.CountAttendance(tx, request.UserID, e.CourseCode)
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
		var components []model.GradeComponentResponse

		for _, g := range grades {
			component := model.GradeComponentResponse{
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
		c.Log.Warnf("failed commit transaction: %+v", err)
		return nil, err
	}

	return reports, nil
}

func (c *GradeUseCase) ListByNpmAndCourseCode(ctx context.Context, request *model.ListInLecturerGradeRequest) ([]model.GradeInLecturerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	enrollment := new(entity.Enrollment)
	if err := c.EnrollmentRepo.FindEnrollmentByNpmAndCourseCode(tx, enrollment, request.Npm, request.CourseCode); err != nil {
		c.Log.WithError(err).Error("failed to find enrollment")
		return nil, err
	}

	grades, err := c.GradeRepo.FindAllByEnrollmentID(tx, enrollment.ID)
	if err != nil {
		c.Log.WithError(err).Error("failed to find grade")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("failed commit transaction: %+v", err)
		return nil, err
	}

	responses := make([]model.GradeInLecturerResponse, len(grades))
	for i, grade := range grades {
		responses[i] = *converter.GradeInLecturerToResponse(&grade)
	}

	return responses, nil
}
