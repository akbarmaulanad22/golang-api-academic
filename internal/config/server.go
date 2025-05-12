package config

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func NewServer(viper *viper.Viper, router *mux.Router) *http.Server {

	port := viper.GetString("SERVER_PORT")

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return server

}
