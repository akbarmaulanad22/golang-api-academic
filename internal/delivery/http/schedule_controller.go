package http

import (
	"encoding/json"
	"net/http"
	"tugasakhir/internal/delivery/http/middleware"
	"tugasakhir/internal/helper"
	"tugasakhir/internal/model"
	"tugasakhir/internal/usecase"

	"github.com/sirupsen/logrus"
)

type ScheduleController struct {
	Log     *logrus.Logger
	UseCase *usecase.ScheduleUseCase
}

func NewScheduleController(useCase *usecase.ScheduleUseCase, logger *logrus.Logger) *ScheduleController {
	return &ScheduleController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *ScheduleController) ListByStudentUserID(w http.ResponseWriter, r *http.Request) {

	auth := middleware.GetUser(r)

	request := &model.ListScheduleRequest{
		UserID: auth.ID,
	}

	// Panggil UseCase
	response, err := c.UseCase.ListScheduleByStudentUserID(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to get schedule: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.ScheduleResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *ScheduleController) ListByLecturerUserID(w http.ResponseWriter, r *http.Request) {

	auth := middleware.GetUser(r)

	request := &model.ListScheduleRequest{
		UserID: auth.ID,
	}

	// Panggil UseCase
	response, err := c.UseCase.ListScheduleByLecturerUserID(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to get schedule: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.ScheduleResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
