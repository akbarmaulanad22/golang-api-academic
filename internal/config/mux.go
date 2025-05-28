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
	// scheduleRepository := repository.NewScheduleRepository(config.Log)
	// attendanceRepository := repository.NewAttendanceRepository(config.Log)

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	// attendanceUseCase := usecase.NewAttendanceUseCase(config.DB, config.Log, config.Validate, scheduleRepository, attendanceRepository)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	// attendanceController := http.newAttendance

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		Router:         config.Router,
		UserController: userController,
		AuthMiddleware: authMiddleware,
	}
	routeConfig.Setup()

}
