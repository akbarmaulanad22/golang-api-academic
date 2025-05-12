package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tugasakhir/internal/config"
	"tugasakhir/internal/helper"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	viper := config.NewViper()
	log := config.NewLogger(viper)
	db := config.NewDatabase(viper, log)
	validator := config.NewValidator(viper)
	server := config.NewServer(viper, router)

	config.NewMux(&config.MuxConfig{
		Router:   router,
		DB:       db,
		Log:      log,
		Validate: validator,
	})

	// Server dijalankan dalam goroutine terpisah
	go func() {
		fmt.Printf("=========================================\n")
		fmt.Printf("Server berjalan pada port...\n")
		fmt.Printf("=========================================\n")

		err := server.ListenAndServe()
		helper.FatalIfErrorWithMessage(err, fmt.Sprintf("Error running server: %s", err))
	}()

	// Channel untuk menangkap signal shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	helper.FatalIfErrorWithMessage(err, fmt.Sprintf("Server forced to shutdown: %s\n", err))

	log.Println("Server Shutdown")
}
