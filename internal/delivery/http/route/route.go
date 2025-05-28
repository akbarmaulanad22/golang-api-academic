package route

import (
	controller "tugasakhir/internal/delivery/http"

	"github.com/gorilla/mux"
)

type RouteConfig struct {
	// router
	Router *mux.Router

	// middleware
	AuthMiddleware mux.MiddlewareFunc

	// all field controller
	UserController *controller.UserController
}

func (route *RouteConfig) Setup() {
	route.SetupGuestRoute()
	route.SetupAuthRoute()
}

func (route *RouteConfig) SetupGuestRoute() {
	// routes that do not require authentication
	route.Router.HandleFunc("/register", route.UserController.Register).Methods("POST")
	route.Router.HandleFunc("/login", route.UserController.Login).Methods("POST")
}

func (route *RouteConfig) SetupAuthRoute() {

	// Buat subrouter khusus untuk route yang butuh auth
	authRouter := route.Router.PathPrefix("/").Subrouter()
	authRouter.Use(route.AuthMiddleware)

	authRouter.Use(route.AuthMiddleware)
	authRouter.HandleFunc("/logout", route.UserController.Logout).Methods("POST")

}
