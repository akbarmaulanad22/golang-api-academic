package config

import (
	"tugasakhir/internal/delivery/http/route"

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
	// userRepository := repository.NewUserRepository(config.Log)

	// setup use cases
	// userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository, userProducer)

	// setup controller
	// userController := http.NewUserController(userUseCase, config.Log)

	// setup middleware
	// authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		Router: config.Router,
		// AuthMiddleware:    authMiddleware,
		// AddressController: addressController,
	}
	routeConfig.Setup()

}
