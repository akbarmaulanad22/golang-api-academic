package http

import (
	"encoding/json"
	"net/http"
	"tugasakhir/internal/delivery/http/middleware"
	"tugasakhir/internal/model"
	"tugasakhir/internal/usecase"

	"github.com/sirupsen/logrus"
)

type GradeController struct {
	Log     *logrus.Logger
	UseCase *usecase.GradeUseCase
}

func NewGradeController(useCase *usecase.GradeUseCase, logger *logrus.Logger) *GradeController {
	return &GradeController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *GradeController) List(w http.ResponseWriter, r *http.Request) {

	// get user by context in middleware
	auth := middleware.GetUser(r)

	// Panggil UseCase
	response, err := c.UseCase.GetCourseGrades(r.Context(), auth.ID)
	if err != nil {
		c.Log.Printf("Failed to get grades: %v", err)
		http.Error(w, "Failed to get grades", http.StatusInternalServerError)
		return
	}

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.GradeResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
