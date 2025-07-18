package config

import (
	"tugasakhir/internal/delivery/http"
	"tugasakhir/internal/delivery/http/middleware"
	"tugasakhir/internal/delivery/http/route"
	"tugasakhir/internal/repository"
	"tugasakhir/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type MuxConfig struct {
	Router   *mux.Router
	DB       *gorm.DB
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func NewMux(config *MuxConfig) {

	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	scheduleRepository := repository.NewScheduleRepository(config.Log)
	attendanceRepository := repository.NewAttendanceRepository(config.Log)
	enrollmentRepository := repository.NewEnrollmentRepository(config.Log)
	courseRepository := repository.NewCourseRepository(config.Log)
	gradeRepository := repository.NewGradeRepository(config.Log)
	studentRepository := repository.NewStudentRepository(config.Log)
	studyProgramRepository := repository.NewStudyProgramRepository(config.Log)
	facultyRepository := repository.NewFacultyRepository(config.Log)
	lecturerRepository := repository.NewLecturerRepository(config.Log)
	classroomRepository := repository.NewClassroomRepository(config.Log)
	gradeComponentRepository := repository.NewGradeComponentRepository(config.Log)

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	attendanceUseCase := usecase.NewAttendanceUseCase(config.DB, config.Log, config.Validate, attendanceRepository, scheduleRepository, studentRepository)
	scheduleUseCase := usecase.NewScheduleUseCase(config.DB, config.Log, config.Validate, scheduleRepository, attendanceRepository)
	enrollmentUseCase := usecase.NewEnrollmentUseCase(config.DB, config.Log, config.Validate, enrollmentRepository)
	gradeUseCase := usecase.NewGradeUseCase(config.DB, config.Log, config.Validate, gradeRepository, scheduleRepository, attendanceRepository, courseRepository, enrollmentRepository)
	courseUseCase := usecase.NewCourseUseCase(config.DB, config.Log, config.Validate, courseRepository)
	studentUseCase := usecase.NewStudentUseCase(config.DB, config.Log, config.Validate, studentRepository, userRepository)
	studyProgramUseCase := usecase.NewStudyProgramUseCase(config.DB, config.Log, config.Validate, studyProgramRepository)
	facultyUseCase := usecase.NewFacultyUseCase(config.DB, config.Log, config.Validate, facultyRepository)
	lecturerUseCase := usecase.NewLecturerUseCase(config.DB, config.Log, config.Validate, lecturerRepository, userRepository)
	classroomUseCase := usecase.NewClassroomUseCase(config.DB, config.Log, config.Validate, classroomRepository)
	gradeComponentUseCase := usecase.NewGradeComponentUseCase(config.DB, config.Log, config.Validate, gradeComponentRepository)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	attendanceController := http.NewAttendanceController(attendanceUseCase, config.Log)
	scheduleController := http.NewScheduleController(scheduleUseCase, config.Log)
	enrollmentController := http.NewEnrollmentController(enrollmentUseCase, config.Log)
	gradeController := http.NewGradeController(gradeUseCase, config.Log)
	courseController := http.NewCourseController(courseUseCase, config.Log)
	studentController := http.NewStudentController(studentUseCase, config.Log)
	studyProgramController := http.NewStudyProgramController(studyProgramUseCase, config.Log)
	facultyController := http.NewFacultyController(facultyUseCase, config.Log)
	lecturerController := http.NewLecturerController(lecturerUseCase, config.Log)
	classroomController := http.NewClassroomController(classroomUseCase, config.Log)
	gradeComponentController := http.NewGradeComponentController(gradeComponentUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		Router:                   config.Router,
		AuthMiddleware:           authMiddleware,
		UserController:           userController,
		AttendanceController:     attendanceController,
		ScheduleController:       scheduleController,
		EnrollmentController:     enrollmentController,
		GradeController:          gradeController,
		CourseController:         courseController,
		StudentController:        studentController,
		StudyProgramController:   studyProgramController,
		FacultyController:        facultyController,
		LecturerController:       lecturerController,
		ClassroomController:      classroomController,
		GradeComponentController: gradeComponentController,
	}
	routeConfig.Setup()

}
