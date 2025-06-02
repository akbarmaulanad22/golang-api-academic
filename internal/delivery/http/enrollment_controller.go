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

type EnrollmentController struct {
	Log     *logrus.Logger
	UseCase *usecase.EnrollmentUseCase
}

func NewEnrollmentController(useCase *usecase.EnrollmentUseCase, logger *logrus.Logger) *EnrollmentController {
	return &EnrollmentController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *EnrollmentController) List(w http.ResponseWriter, r *http.Request) {

	auth := middleware.GetUser(r)

	// Panggil UseCase
	response, err := c.UseCase.GetEnrollmentByStudentUserID(r.Context(), auth.ID)
	if err != nil {
		c.Log.Printf("Failed to get enrollment: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.EnrollmentResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
