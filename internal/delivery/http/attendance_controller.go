package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tugasakhir/internal/delivery/http/middleware"
	"tugasakhir/internal/helper"
	"tugasakhir/internal/model"
	"tugasakhir/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type AttendanceController struct {
	Log     *logrus.Logger
	UseCase *usecase.AttendanceUseCase
}

func NewAttendanceController(useCase *usecase.AttendanceUseCase, logger *logrus.Logger) *AttendanceController {
	return &AttendanceController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *AttendanceController) AttendStudent(w http.ResponseWriter, r *http.Request) {

	auth := middleware.GetUser(r)

	// Buat request untuk use case
	request := &model.AttendanceCreateResponse{
		UserId: auth.ID,
	}

	// Panggil UseCase.Logout
	response, err := c.UseCase.AttendStudent(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to attend user: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.AttendanceResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (c *AttendanceController) AttendLecturer(w http.ResponseWriter, r *http.Request) {

	auth := middleware.GetUser(r)

	// Buat request untuk use case
	request := &model.AttendanceCreateResponse{
		UserId: auth.ID,
	}

	// Panggil UseCase.Logout
	response, err := c.UseCase.AttendLecturer(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to attend user: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.AttendanceResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (c *AttendanceController) ListByUserID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	npm := vars["npm"]

	npmInt, err := strconv.Atoi(npm)
	if err != nil {
		c.Log.Printf("Failed to parse npm: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	request := &model.ListAttendanceRequest{
		Npm: uint(npmInt),
	}

	// Panggil UseCase.Logout
	response, err := c.UseCase.ListByUserID(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to attend user: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.AttendanceResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
