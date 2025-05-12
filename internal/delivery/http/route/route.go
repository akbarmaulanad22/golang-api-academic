package route

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteConfig struct {
	// router
	Router *mux.Router
	// middleware

	// all field controller

}

func (route *RouteConfig) Setup() {
	route.SetupGuestRoute()
	route.SetupAuthRoute()
}

func (route *RouteConfig) SetupGuestRoute() {
	route.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test")
	})
	// routes that do not require authentication
}

func (route *RouteConfig) SetupAuthRoute() {
	// routes that require authentication
}
