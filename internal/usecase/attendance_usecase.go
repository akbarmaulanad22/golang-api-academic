package usecase

import (
	"context"
	"time"
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
	"tugasakhir/internal/model/converter"
	"tugasakhir/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AttendanceUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	AttendanceRepository *repository.AttendanceRepository
	ScheduleRepository   *repository.ScheduleRepository
	StudentRepository    *repository.StudentRepository
}

func NewAttendanceUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	attendanceRepository *repository.AttendanceRepository,
	scheduleRepository *repository.ScheduleRepository,
	studentRepository *repository.StudentRepository,

) *AttendanceUseCase {

	return &AttendanceUseCase{
		DB:                   db,
		Log:                  log,
		Validate:             validate,
		AttendanceRepository: attendanceRepository,
		ScheduleRepository:   scheduleRepository,
		StudentRepository:    studentRepository,
	}

}

func (c *AttendanceUseCase) AttendStudent(ctx context.Context, request *model.AttendanceCreateResponse) (*model.AttendanceResponse, error) {

	// validate
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body: %+v", err)
		return nil, err
	}

	// start transaction

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	scheduleID, err := c.ScheduleRepository.GetStudentActiveScheduleIDByUserID(tx, request.UserId)
	if err != nil {
		c.Log.Warnf("Failed get active schedule : %+v", err)
		return nil, err
	}

	if c.AttendanceRepository.IsLecturerPresent(tx, scheduleID) {
		c.Log.Warnf("failed to check lecturer status : %+v", err)
		return nil, err
	}

	if c.AttendanceRepository.HasAlreadyAttended(tx, request.UserId, scheduleID) {
		c.Log.Warnf("failed to check student status : %+v", err)
		return nil, err
	}

	attendance := &entity.Attendance{
		Status:     "Hadir",
		Time:       time.Now(),
		ScheduleId: scheduleID,
		UserId:     request.UserId,
	}

	// create attendance
	if err := c.AttendanceRepository.Create(tx, attendance); err != nil {
		c.Log.Warnf("Failed create attendance to database : %+v", err)
		return nil, err
	}

	// commit db transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	// return converter.AttendanceToResponse(attendance), nil
	return converter.AttendanceToResponse(attendance), nil
}

func (c *AttendanceUseCase) AttendLecturer(ctx context.Context, request *model.AttendanceCreateResponse) (*model.AttendanceResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// 1. Cari jadwal aktif
	scheduleID, err := c.ScheduleRepository.GetLecturerActiveScheduleByUserID(tx, request.UserId)
	if err != nil {
		c.Log.Warnf("Failed get active schedule : %+v", err)
		return nil, err
	}

	// 2. Cek apakah dosen sudah pernah absen hari ini
	if c.AttendanceRepository.HasAlreadyAttended(tx, request.UserId, scheduleID) {
		c.Log.Warnf("failed to check lecturer status : %+v", err)
		return nil, err
	}

	// 3. Simpan absensi
	attendance := &entity.Attendance{
		UserId:     request.UserId,
		ScheduleId: scheduleID,
		Status:     "Hadir",
		Time:       time.Now(),
	}

	if err := c.AttendanceRepository.Create(tx, attendance); err != nil {
		c.Log.Warnf("Failed create attendance to database : %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Printf("failed commit transaction: %v", err)
		return nil, err
	}

	return converter.AttendanceToResponse(attendance), nil
}

func (c *AttendanceUseCase) ListByUserID(ctx context.Context, request *model.ListAttendanceRequest) ([]model.AttendanceResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	userID, err := c.StudentRepository.GetStudentUserIDByNPM(tx, request.Npm)
	if err != nil {
		c.Log.WithError(err).Error("failed to find student by npm")
		return nil, err
	}

	attendances, err := c.AttendanceRepository.FindAllByUserID(tx, userID)
	if err != nil {
		c.Log.WithError(err).Error("failed to find attendances")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	responses := make([]model.AttendanceResponse, len(attendances))
	for i, attendance := range attendances {
		responses[i] = *converter.AttendanceToResponse(&attendance)
	}

	return responses, nil
}
